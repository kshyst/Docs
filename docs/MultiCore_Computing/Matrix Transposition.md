# Case Study: Matrix Transposition

Matrix transposition is a classic parallel programming benchmark. While mathematically simple, implementing an efficient matrix transpose on a GPU exposes the core trade-offs between **Global Memory Coalescing** and **Shared Memory Bank Conflicts**.

This case study provides a deep-dive analysis of why naive implementations perform poorly, how shared memory serves as a coalesced staging area, and how memory padding eliminates bank conflicts.

---

## 1. The Global Memory Bottleneck: Coalescing

In GPU architectures (like Nvidia CUDA), global memory (DRAM) is accessed in chunks of 32, 64, or 128 bytes. A **coalesced memory access** occurs when all 32 threads in a warp access consecutive memory locations in the same instruction cycle. When this happens, the memory controller merges these individual requests into a single transaction, achieving maximum bandwidth.

If threads in a warp access scattered (non-consecutive) memory locations, the hardware must perform multiple separate memory transactions (uncoalesced access), resulting in a severe performance drop.

```
Coalesced Access (1 Transaction):
Thread:   0   1   2   3  ...  31
Address: 00  04  08  12  ... 124  --> [ [DRAM Cache Line] ]

Uncoalesced Access (32 Transactions):
Thread:   0   1   2   3  ...  31
Address: 00  64 128 192  ...      --> [DRAM Line 1] [DRAM Line 2] [DRAM Line 3] ...
```

### The Transpose Dilemma
By definition, transposing a matrix $A$ of size $W \times H$ to $A^T$ requires mapping:

$$A^T[x][y] = A[y][x]$$

If we attempt a naive copy from input to output in global memory:
* **Reading is coalesced**: Thread $i$ reads $A[y][x + i]$ (consecutive elements in a row).
* **Writing is uncoalesced**: Thread $i$ writes to $A^T[x + i][y]$ (elements down a column, separated by $W$ elements in memory).

Since DRAM latency is high (~100–400 cycles), the uncoalesced write phase cripples performance.

---

## 2. Shared Memory as a Coalescing Bridge

To solve this, we can use **Shared Memory** (which has extremely low latency and does not suffer from global-style coalescing penalties) as a temporary staging tile. 

The strategy is:
1. **Coalesced Read**: Read a block of data from global memory by rows.
2. **Coalesced Write to Shared**: Write this block into a 2D shared memory tile by rows.
3. **Synchronization**: Call `__syncthreads()` to ensure the entire tile is loaded.
4. **Coalesced Read from Shared (Transposed)**: Read from the shared memory tile *by columns*.
5. **Coalesced Write to Global**: Write the data to global memory by rows (which writes consecutive transposed values).

```
Global Input (Row-Major)
 ┌───┬───┬───┬───┐
 │ 0 │ 1 │ 2 │ 3 │  <-- Coalesced Read
 └───┴───┴───┴───┘
         │
         ▼
 Shared Memory Tile
 ┌───┬───┬───┬───┐
 │ 0 │ 1 │ 2 │ 3 │  <-- Row-Major Write
 ├───┼───┼───┼───┤
 │   │   │   │   │
 └───┴───┴───┴───┘
         │
         ▼ (Transposed Read from Shared)
 ┌───┐
 │ 0 │
 ├───┤
 │   │
 └───┘
         │
         ▼
Global Output (Row-Major Transposed)
 ┌───┬───┬───┬───┐
 │ 0 │   │   │   │  <-- Coalesced Write
 └───┴───┴───┴───┘
```

While this technique achieves **100% coalesced global memory access**, the transposed read from shared memory (reading column-by-column) creates a **shared memory bank conflict**!

---

## 3. Deep Dive: The Naive Shared Memory Bank Conflict

Let's analyze a tile of size $32 \times 32$ declared as:
```cuda
__shared__ float tile[32][32];
```

A warp of 32 threads executes a transposed read from the tile. Across the warp, `threadIdx.y` is a constant value $C$ (the column being read) and `threadIdx.x` (represented by $i$ for $i \in [0, 31]$) represents the row. The instruction is:

```cuda
float val = tile[threadIdx.x][threadIdx.y]; // tile[i][C]
```

### The Mathematical Explanation
For a 2D array of size $M \times N$, the flat index of element `array[r][c]` is:

$$\text{Index} = r \times N + c$$

For `tile[i][C]` where $N = 32$:

$$\text{Index} = i \times 32 + C$$

Shared memory maps 32-bit (4-byte) elements to 32 memory banks using a modulo 32 operation:

$$\text{Bank Index} = \text{Index} \bmod 32$$

$$\text{Bank Index} = (i \times 32 + C) \bmod 32$$

Since $(i \times 32) \bmod 32 = 0$ for any integer $i$:

$$\text{Bank Index} = C \bmod 32$$

Because $C$ is a constant value for the entire warp, **every thread in the warp accesses Bank $C$**. 

### Trace Table: Naive $32 \times 32$ Tile (32-Way Conflict)
Assuming we are reading column $C = 5$:

| Warp Thread | Requested Element | Flat 1D Index | Byte Offset (Index $\times 4$) | Bank Formula $(\text{Index} \bmod 32)$ | Target Memory Bank |
| :---: | :---: | :---: | :---: | :---: | :---: |
| **Thread 0** | `tile[0][5]` | 5 | 20 | $5 \bmod 32$ | **Bank 5** |
| **Thread 1** | `tile[1][5]` | 37 | 148 | $37 \bmod 32$ | **Bank 5** |
| **Thread 2** | `tile[2][5]` | 69 | 276 | $69 \bmod 32$ | **Bank 5** |
| **Thread 3** | `tile[3][5]` | 101 | 404 | $101 \bmod 32$ | **Bank 5** |
| ... | ... | ... | ... | ... | ... |
| **Thread 31** | `tile[31][5]` | 997 | 3988 | $997 \bmod 32$ | **Bank 5** |

All 32 threads request different addresses within **Bank 5**. The hardware has to serialize these requests, taking **32 clock cycles** to complete a single instruction.

---

## 4. Deep Dive: The Padded Solution

We can resolve this conflict completely by adding **padding** to the shared memory columns. We change the tile allocation to:
```cuda
__shared__ float tile[32][32 + 1]; // tile[32][33]
```

### The Mathematical Explanation
With $N = 33$, the flat index of `tile[i][C]` becomes:

$$\text{Index} = i \times 33 + C$$

Let's calculate the bank index:

$$\text{Bank Index} = (i \times 33 + C) \bmod 32$$

Using modulo arithmetic, $33 \equiv 1 \pmod{32}$:

$$\text{Bank Index} = (i \times 1 + C) \bmod 32 = (i + C) \bmod 32$$

Since $i$ (the thread index) goes from $0$ to $31$:
* Thread 0 accesses Bank $(0 + C) \bmod 32$
* Thread 1 accesses Bank $(1 + C) \bmod 32$
* Thread 2 accesses Bank $(2 + C) \bmod 32$
* ...
* Thread 31 accesses Bank $(31 + C) \bmod 32$

Every thread maps to a **unique bank** from $0$ to $31$. There are zero conflicts, and the read executes in **1 clock cycle**.

### Trace Table: Padded $32 \times 33$ Tile (Conflict-Free)
Assuming we are reading column $C = 5$:

| Warp Thread | Requested Element | Flat 1D Index | Byte Offset (Index $\times 4$) | Bank Formula $(\text{Index} \bmod 32)$ | Target Memory Bank |
| :---: | :---: | :---: | :---: | :---: | :---: |
| **Thread 0** | `tile[0][5]` | 5 | 20 | $5 \bmod 32$ | **Bank 5** |
| **Thread 1** | `tile[1][5]` | 38 | 152 | $38 \bmod 32$ | **Bank 6** |
| **Thread 2** | `tile[2][5]` | 71 | 284 | $71 \bmod 32$ | **Bank 7** |
| **Thread 3** | `tile[3][5]` | 104 | 416 | $104 \bmod 32$ | **Bank 8** |
| ... | ... | ... | ... | ... | ... |
| **Thread 31** | `tile[31][5]` | 1028 | 4112 | $1028 \bmod 32$ | **Bank 4** |

Each thread is assigned to a distinct memory bank. The access pattern is perfectly diagonal and parallelized.

---

## 5. Dynamic Shared Memory Case

When using **Dynamic Shared Memory**, we cannot allocate a 2D array like `tile[32][33]` directly because the compiler expects a 1D array declaration:
```cuda
extern __shared__ float tile[];
```

To implement padding in dynamic shared memory, we must perform the 2D-to-1D index linearization manually. The pitch (stride) of each row must be set to `TILE_DIM + 1` (33) instead of `TILE_DIM` (32).

### Index Mapping
* **Write to Shared** (Row-major):
  ```cuda
  tile[threadIdx.y * (TILE_DIM + 1) + threadIdx.x] = idata[y * width + x];
  ```
* **Read from Shared** (Transposed, column-major):
  ```cuda
  odata[y_t * height + x_t] = tile[threadIdx.x * (TILE_DIM + 1) + threadIdx.y];
  ```

---

## 6. Complete CUDA Kernel Comparisons

Below is the comparative reference code showing all three variants:

```cuda
#include <device_launch_parameters.h>

#define TILE_DIM 32

// ============================================================================
// 1. NAIVE KERNEL (With 32-way Bank Conflict)
// ============================================================================
__global__ void transposeNaive(float *odata, const float *idata, int width, int height) {
    __shared__ float tile[TILE_DIM][TILE_DIM];

    int x = blockIdx.x * TILE_DIM + threadIdx.x;
    int y = blockIdx.y * TILE_DIM + threadIdx.y;

    if (x < width && y < height) {
        tile[threadIdx.y][threadIdx.x] = idata[y * width + x];
    }
    
    __syncthreads();

    int x_t = blockIdx.y * TILE_DIM + threadIdx.x;
    int y_t = blockIdx.x * TILE_DIM + threadIdx.y;

    if (x_t < height && y_t < width) {
        // Strided shared memory access: 32-way bank conflict
        odata[y_t * height + x_t] = tile[threadIdx.x][threadIdx.y];
    }
}

// ============================================================================
// 2. OPTIMIZED STATIC KERNEL (Conflict-Free via Padding)
// ============================================================================
__global__ void transposePaddedStatic(float *odata, const float *idata, int width, int height) {
    // Column dimension padded to 33
    __shared__ float tile[TILE_DIM][TILE_DIM + 1];

    int x = blockIdx.x * TILE_DIM + threadIdx.x;
    int y = blockIdx.y * TILE_DIM + threadIdx.y;

    if (x < width && y < height) {
        tile[threadIdx.y][threadIdx.x] = idata[y * width + x];
    }
    
    __syncthreads();

    int x_t = blockIdx.y * TILE_DIM + threadIdx.x;
    int y_t = blockIdx.x * TILE_DIM + threadIdx.y;

    if (x_t < height && y_t < width) {
        // Shifted shared memory access: 100% conflict-free
        odata[y_t * height + x_t] = tile[threadIdx.x][threadIdx.y];
    }
}

// ============================================================================
// 3. OPTIMIZED DYNAMIC KERNEL (Conflict-Free via Padding)
// ============================================================================
__global__ void transposePaddedDynamic(float *odata, const float *idata, int width, int height) {
    // Dynamic 1D allocation
    extern __shared__ float tile[];

    int x = blockIdx.x * TILE_DIM + threadIdx.x;
    int y = blockIdx.y * TILE_DIM + threadIdx.y;

    if (x < width && y < height) {
        // Manual 2D indexing with padded row stride (TILE_DIM + 1)
        int shared_idx = threadIdx.y * (TILE_DIM + 1) + threadIdx.x;
        tile[shared_idx] = idata[y * width + x];
    }
    
    __syncthreads();

    int x_t = blockIdx.y * TILE_DIM + threadIdx.x;
    int y_t = blockIdx.x * TILE_DIM + threadIdx.y;

    if (x_t < height && y_t < width) {
        // Manual transposed 2D indexing with padded row stride
        int shared_idx_t = threadIdx.x * (TILE_DIM + 1) + threadIdx.y;
        odata[y_t * height + x_t] = tile[shared_idx_t];
    }
}
```
