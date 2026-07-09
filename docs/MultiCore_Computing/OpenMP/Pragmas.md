# Pragmas

```c 
#pragma omp parallel for schedule(static)
```

`schedule(static)` This defines exactly how the iterations are distributed among the threads. static means the total number of iterations is divided into roughly equal chunks and assigned to the threads before the loop even starts.

If your loop runs 100 times and you have 4 threads:

With schedule(static), OpenMP will immediately assign:

- Iterations 0-24 to Thread 0
- Iterations 25-49 to Thread 1
- Iterations 50-74 to Thread 2
- Iterations 75-99 to Thread 3

---

```c 
#pragma omp parallel for schedule(dynamic)
```

This directive tells OpenMP to distribute the iterations of a for loop among available threads dynamically at runtime.

Instead of pre-assigning blocks of iterations to threads before the loop starts, OpenMP hands out small “chunks” of iterations to threads as they become available.

A thread gets a chunk of iterations (the default chunk size is 1). Once the thread finishes that chunk, it goes back to the scheduler and requests another chunk. This continues until all iterations are finished.

You can specify how many iterations are handed out at a time, like this: `#pragma omp for schedule(dynamic, 5)`. 

---

```c 
#pragma omp sections
```

While `omp for` distributes iterations of a loop among threads, `omp sections` distributes distinct, independent blocks of code among threads.
 
```c
#include <stdio.h>
#include <omp.h>

int main() {
    // Create a team of threads
    #pragma omp parallel
    {
        // Distribute the following sections among the threads
        #pragma omp sections
        {
            #pragma omp section
            {
                printf("Task A executed by thread %d\n", omp_get_thread_num());
                // Do some specific work for Task A
            }

            #pragma omp section
            {
                printf("Task B executed by thread %d\n", omp_get_thread_num());
                // Do some completely different work for Task B
            }

            #pragma omp section
            {
                printf("Task C executed by thread %d\n", omp_get_thread_num());
                // Do some other work for Task C
            }
        } // Implicit barrier here: threads wait until all sections are done
    }

    return 0;
}
```

Outputs: 

```shell
Task C executed by thread 13
Task A executed by thread 7
Task B executed by thread 8

Task B executed by thread 3
Task C executed by thread 1
Task A executed by thread 9

Task B executed by thread 5
Task C executed by thread 13
Task A executed by thread 6
```

By default, threads will wait at the end of the omp sections block until all individual sections have finished executing. You can bypass this by adding nowait (`#pragma omp sections nowait`). This creates non-blocking threads

---

```
#pragma omp barrier
```

In OpenMP, `#pragma omp barrier` is an explicit synchronization point.

When threads in a parallel region reach this directive, they will stop and wait. No thread is allowed to proceed past the barrier until all threads in the team have reached that exact same point.

You use a barrier when you have multiple phases of parallel work, and Phase 2 absolutely relies on Phase 1 being 100% complete by all threads. It prevents race conditions by ensuring all threads are on the same page before moving forward.

```c 
#include <stdio.h>
#include <omp.h>
#include <unistd.h> // For sleep()

int main() {
    int shared_data[4] = {0, 0, 0, 0};

    #pragma omp parallel num_threads(4)
    {
        int tid = omp_get_thread_num();

        // PHASE 1: Each thread writes to its specific index
        printf("Thread %d is starting Phase 1.\n", tid);
        shared_data[tid] = tid * 10; 
        
        if (tid == 0) {
            sleep(2); // Simulate thread 0 taking a long time
        }

        // Wait for all threads to finish Phase 1 before moving on
        #pragma omp barrier

        // PHASE 2: Threads can safely read from the array knowing it is fully populated
        printf("Thread %d is starting Phase 2. Data from thread 0 is %d.\n", 
               tid, shared_data[0]);
    }

    return 0;
}
```

You often don’t need to write `#pragma omp barrier` because OpenMP automatically inserts implicit barriers at the end of work-sharing constructs like `#pragma omp for`, `#pragma omp sections`, and `#pragma omp single`.

### Implicit Barriers

**nowait clause**: If you want to remove an implicit barrier at the end of a for loop or sections block, you use the nowait clause (e.g., `#pragma omp for nowait`), which lets threads proceed immediately without waiting for others.

> **Deadlocks:** Be careful! If a barrier is placed inside a conditional statement (like an if block) that only some threads will enter, the threads that enter will wait forever for the threads that didn’t, causing your program to freeze (deadlock). All threads in the team must be able to hit the barrier.

---

```
#pragma omp critical
```

A critical section is a block of code that can be executed by only one thread at a time. If a thread reaches a critical section and another thread is already inside it, the arriving thread will wait until the first thread exits the section.

You use it to prevent race conditions when multiple threads need to safely read and write to the same shared variable. Without it, threads might overwrite each other’s changes, leading to incorrect results.

```c 
#include <stdio.h>
#include <omp.h>

int main() {
    int total_sum = 0;

    #pragma omp parallel num_threads(4)
    {
        int local_sum = 0;
        int tid = omp_get_thread_num();
        
        // Do some independent work (simulated by a simple loop)
        for (int i = 0; i < 1000; i++) {
            local_sum += 1;
        }

        // Only one thread at a time can add its local_sum to total_sum
        #pragma omp critical
        {
            printf("Thread %d is adding %d to total_sum.\n", tid, local_sum);
            total_sum += local_sum; 
        }
    }

    printf("Final total_sum: %d\n", total_sum);
    return 0;
}
```

---

``` 
#pragma omp parallel for reduction(+:area)
```

The reduction clause instructs the compiler to create private copies of the area variable for every thread. At the end of the loop partial sums are combined on the shared `area` variable.


