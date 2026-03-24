# Semaphore
A variable or abstract data type used to control access to a common resource by multiple threads and avoid critical section problems in a concurrent system such as a multitasking operating system.

Though semaphores are useful for preventing race conditions, they do not guarantee their absence. Semaphores that allow an arbitrary resource count are called counting semaphores, while semaphores that are restricted to the values 0 and 1 (or locked/unlocked, unavailable/available) are called binary semaphores and are used to implement locks.

## Mutex vs Semaphore

A mutex is a locking mechanism that sometimes uses the same basic implementation as the binary semaphore. However, they differ in how they are used. While a binary semaphore may be colloquially referred to as a mutex, a true mutex has a more specific use-case and definition, in that only the task that locked the mutex is supposed to unlock it. 