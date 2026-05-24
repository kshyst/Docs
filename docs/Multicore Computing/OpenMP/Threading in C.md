# Threading in C

## Volatile Variables

In C, `volatile` is a keyword (a type qualifier) that tells the compiler that a variable’s value can change unexpectedly from outside the current code execution flow (such as by hardware, interrupts, or other threads).

Because the value can change at any time, it forces the compiler to read the variable from memory every single time it is used, and prevents the compiler from applying certain optimizations.

``` 
volatile int temp = 0;
for (int j = 0; j < work; j++) {
    temp++;
}
```
## OMP Environment Variables

- `OMP_NUM_THREADS`: Sets maximum number of threads
- `OMP_SCHEDULE`: Determines how iterations are scheduled when schedule clause is set to `runtime`.
    - empty
    - `chunk`
- `OMP_DYNAMIC`: Dynamic adjustment for threads for parallel regions
    - true
    - false
- `OMP_NESTED`: Nested parallelism.
    - true
    - false

## Getting Thread Info

For getting current thread's numbers
```
omp_get_thread_num();
```
For getting count of all threads
``` 
omp_get_num_threads();
```

## Getting System Info

For getting current time.
```
omp_get_wtime()
```
Time between ticks in seconds
``` 
omp_get_wtick()
```

## Shared Memory

### Shared Variable

In OpenMP, a shared variable is a variable that is accessible by all threads within a parallel region.

There is only one copy of a shared variable in memory. If one thread changes its value, that change is immediately visible to all other threads.

By default in C and C++, most variables declared before a #pragma omp parallel block are treated as shared. However, you can also explicitly define them using the shared() clause.

