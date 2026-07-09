# Warp and SIMT Execution Model

In GPU programming (specifically Nvidia's CUDA architecture), a **Warp** is the fundamental execution unit. Understanding warps, warp scheduling, divergence, and warp-level primitives is essential to writing high-performance GPU kernels.

> [!NOTE]
> For a visual diagram of how warps relate to threads, blocks, grids, and the physical GPU Streaming Multiprocessors (SMs), see the [CUDA Memory and Execution Hierarchy](file:///home/kshyst/Desktop/Docs/Docs/docs/MultiCore_Computing/CUDA%20Memory%20and%20Execution%20Hierarchy.md) page.

---

## 1. What is a Warp?

A warp is a group of **32 threads** that are executed physically in parallel on a Streaming Multiprocessor (SM). It is the smallest unit of execution scheduled by the GPU hardware.

```
       ┌───────────────────────────────────────────────────┐
       │                 SM Warp Scheduler                 │
       └─────────────────────────┬─────────────────────────┘
                                 │  (Fetches & issues 1 instruction)
                                 ▼
┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐
│T0 │T1 │T2 │T3 │T4 │T5 │T6 │T7 │...│...│...│...│...│...│...│T31│  <-- 32 Threads in a Warp
└───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┘
```

### The SIMT Model
GPUs use the **SIMT (Single Instruction, Multiple Threads)** execution model:
* A warp scheduler issues the **same instruction** to all 32 threads in the warp simultaneously.
* Each thread has its own instruction address counter and register state, allowing them to carry out the instruction on different data paths.
* SIMT is distinct from SIMD (Single Instruction, Multiple Data) because in SIMT, each thread is treated as an independent execution path by the programmer, even though the hardware group-schedules them.

---

## 2. Warp Partitioning (How Threads Map to Warps)

When you launch a grid of thread blocks, the GPU hardware automatically partitions each block into warps. 

### Warp Allocation Rules
1. Warps are formed from threads with consecutive thread IDs.
2. For a 1D thread block, Thread $0$ to $31$ form Warp 0, Thread $32$ to $63$ form Warp 1, and so on.
3. For 2D and 3D thread blocks, threads are first flattened into a 1D index, and then grouped.

#### The Flattening Formula
For a 3D block with dimensions $(D_x, D_y, D_z)$, the flat thread ID $T_{\text{flat}}$ for a thread at index $(x, y, z)$ is:

$$T_{\text{flat}} = x + (y \times D_x) + (z \times D_x \times D_y)$$

Once flattened, Warp $k$ contains threads where:

$$\lfloor T_{\text{flat}} / 32 \rfloor = k$$

> [!IMPORTANT]
> If a thread block's size is not a multiple of 32, the last warp will be **partially populated** (e.g., if block size is 80, there will be 3 warps: Warp 0 and 1 have 32 threads, Warp 2 has 16 active threads). The inactive threads are masked out by the hardware, which wastes execution resources. Always try to make block sizes multiples of 32.

---

## 3. Warp Scheduling and Latency Hiding

In CPUs, latency is managed via large caches and complex branch prediction. GPUs take a different approach called **Latency Hiding** using massive multithreading and fast context switching.

```
Time ──►
Warp 0:  [ Compute ] [ Stalled on Memory Read ] ────────────────────────► [ Resume ]
Warp 1:              [ Compute ] [ Stalled on Cache Miss ] ─────────────►
Warp 2:                          [ Compute ] [ Stalled on Pipeline ] ───►
Warp 3:                                      [ Compute ] ───────────────►
```

### Zero-Overhead Context Switching
* An SM contains multiple warp schedulers. At any clock cycle, the scheduler selects a **ready warp** (a warp whose instructions and operands are ready) and issues an instruction to its execution units.
* If a warp stalls (e.g., waiting 200 cycles for a global memory read), the warp scheduler immediately switches to another ready warp.
* This context switch has **zero overhead** because all register files, local memory states, and execution contexts for all active warps are stored directly on-chip. There is no saving or restoring of registers to main memory.

---

## 4. Warp Divergence

Since all threads in a warp must execute the same instruction in a clock cycle, a problem arises when threads in a warp encounter a conditional branch (like `if-else`) and take different paths.

```
       if (threadIdx.x < 16) {
           // Path A
       } else {
           // Path B
       }
```

### The Serialization Penalty
1. When Warp 0 executes, threads 0–15 satisfy the condition and enter **Path A**. Threads 16–31 do not satisfy it and are **masked out** (disabled) by the hardware.
2. The SM executes **Path A** instructions for threads 0–15.
3. Once completed, the SM disables threads 0–15, enables threads 16–31, and executes **Path B** instructions.
4. Finally, the execution paths merge back, and all threads become active again.

This serialization is called **Warp Divergence**. 

```
Cycle 1: [T0...T15: Active (Path A)] [T16...T31: MASKED]
Cycle 2: [T0...T15: MASKED]          [T16...T31: Active (Path B)]
```

> [!WARNING]
> Warp divergence degrades instruction throughput. If a warp diverges into two branches, execution time for that branch block roughly doubles. If it diverges into 32 branches (e.g., inside a switch statement), performance can drop to $\sim 3\%$ of peak efficiency.

### Mitigating Divergence
* **Branch Alignment**: Restructure conditions so that whole warps take the same branch (e.g., group branches by multiples of 32).
* **Branchless Math (Predication)**: Use ternary operators or mathematical formulas instead of conditional statements where possible. The compiler often uses predication for short branches, executing both paths but writing results conditionally without branching.

---

## 5. Warp-Level Primitives (Intrinsics)

Historically, threads within a block communicated using Shared Memory. However, CUDA introduces **Warp-Level Primitives** that allow threads within the same warp to share data, vote, or synchronize directly using register transfer, bypassing shared memory entirely. This is extremely fast.

### A. Data Shuffle Intrinsics
Shuffle instructions allow threads to read variables directly from registers of other threads in the same warp.

* `__shfl_sync(unsigned mask, T var, int srcLane, int width=32)`: Reads `var` from the thread at index `srcLane` (lane ID).
* `__shfl_up_sync(unsigned mask, T var, unsigned int delta, int width=32)`: Shifts data up the warp (lane ID minus `delta`).
* `__shfl_down_sync(unsigned mask, T var, unsigned int delta, int width=32)`: Shifts data down the warp (lane ID plus `delta`).
* `__shfl_xor_sync(unsigned mask, T var, int laneMask, int width=32)`: Swaps data between threads using bitwise XOR on lane IDs (extremely useful for tree reductions).

> **What is the `mask` parameter?**
> The `mask` is a 32-bit unsigned integer representing the warp threads participating in the primitive (typically `0xffffffff` for all threads). It is required in modern CUDA (CUDA 9+) to prevent undefined behavior under independent thread scheduling.

### B. Warp Vote Intrinsics
Allows threads to query conditions across the entire warp.

* `__any_sync(unsigned mask, int predicate)`: Returns non-zero if the predicate is true for *any* participating thread.
* `__all_sync(unsigned mask, int predicate)`: Returns non-zero if the predicate is true for *all* participating thread.
* `__ballot_sync(unsigned mask, int predicate)`: Returns a 32-bit mask where bit $N$ is set if the predicate is true for thread $N$.

---

## 6. Code Comparison: Warp-Level Reduction vs. Shared Memory Reduction

Parallel reduction (summing up an array of numbers) is a standard pattern. Below, we compare a classic shared-memory block reduction with a highly optimized warp-shuffle reduction.

```cuda
#include <device_launch_parameters.h>

// ============================================================================
// 1. STANDARD BLOCK REDUCTION (Using Shared Memory)
// ============================================================================
__global__ void sharedMemoryReduction(float *g_out, const float *g_in, int n) {
    __shared__ float sdata[256];

    unsigned int tid = threadIdx.x;
    unsigned int idx = blockIdx.x * blockDim.x + threadIdx.x;

    // Load data from global memory
    sdata[tid] = (idx < n) ? g_in[idx] : 0.0f;
    __syncthreads();

    // Do reduction in shared memory
    for (unsigned int s = blockDim.x / 2; s > 0; s >>= 1) {
        if (tid < s) {
            sdata[tid] += sdata[tid + s];
        }
        __syncthreads(); // Required at every step to prevent data hazards
    }

    // Write block result to global memory
    if (tid == 0) {
        g_out[blockIdx.x] = sdata[0];
    }
}

// ============================================================================
// 2. WARP-LEVEL SHUFFLE REDUCTION (Bypassing Shared Memory)
// ============================================================================
__inline__ __device__ float warpReduceSum(float val) {
    // Perform tree reduction using XOR shuffle
    // 0xffffffff represents all 32 threads in the warp
    val += __shfl_xor_sync(0xffffffff, val, 16);
    val += __shfl_xor_sync(0xffffffff, val, 8);
    val += __shfl_xor_sync(0xffffffff, val, 4);
    val += __shfl_xor_sync(0xffffffff, val, 2);
    val += __shfl_xor_sync(0xffffffff, val, 1);
    return val;
}

__global__ void warpShuffleReduction(float *g_out, const float *g_in, int n) {
    // 8 warps in a block of 256 threads
    __shared__ float warp_sums[32]; 

    unsigned int tid = threadIdx.x;
    unsigned int idx = blockIdx.x * blockDim.x + threadIdx.x;
    
    // Load and perform initial warp-level reduction
    float val = (idx < n) ? g_in[idx] : 0.0f;
    val = warpReduceSum(val);

    int lane = threadIdx.x % 32;
    int warp_id = threadIdx.x / 32;

    // The first thread of each warp writes its sum to shared memory
    if (lane == 0) {
        warp_sums[warp_id] = val;
    }
    __syncthreads();

    // Now, reduce the warp sums using the first warp
    if (warp_id == 0) {
        // Load the warp sum (if thread belongs to an active warp, else 0)
        val = (lane < blockDim.x / 32) ? warp_sums[lane] : 0.0f;
        // Final warp reduction
        val = warpReduceSum(val);
    }

    // Write final block result to global memory
    if (tid == 0) {
        g_out[blockIdx.x] = val;
    }
}
```

### Why the Warp Shuffle Implementation is Faster:
1. **Fewer Barrier Synchronizations**: The shared memory version requires a `__syncthreads()` call at *every* step of the loop. The warp shuffle version only requires a single block-level sync.
2. **Registers Instead of Cache**: Registers are read/written in a single instruction cycle, bypassing the shared memory pipeline entirely, which reduces memory bandwidth pressure and latency.
3. **Zero Bank Conflicts**: Since registers are thread-local and transfer is direct, there are no bank conflicts in `warpReduceSum()`.
