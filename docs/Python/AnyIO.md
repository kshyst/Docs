# AnyIO and TaskGroups

AnyIO is a asynchronous compatibility library that allows you to write code which can run on top of either **asyncio** or **trio**. It is heavily inspired by Trio's philosophy of **Structured Concurrency**.

## Why AnyIO?

- **Back-end Agnostic**: Write once, run on asyncio or trio.
- **Structured Concurrency**: Prevents "leaking" tasks and ensures clean exits.
- **Improved Task Management**: Superior error handling compared to `asyncio.gather`.
- **Nursery Pattern**: Uses TaskGroups (similar to Trio's nurseries) to manage task lifetimes.

## TaskGroups: The Core Primitive

A `TaskGroup` is an asynchronous context manager that serves as a container for concurrent tasks.

### Basic Usage

You use `create_task_group()` to start a group and `tg.start_soon()` to spawn tasks.

```python
import anyio

async def worker(name, delay):
    print(f"Task {name} starting...")
    await anyio.sleep(delay)
    print(f"Task {name} finished after {delay}s")

async def main():
    async with anyio.create_task_group() as tg:
        tg.start_soon(worker, "A", 1)
        tg.start_soon(worker, "B", 2)
    
    # Execution only reaches here after ALL tasks in the group finish
    print("All tasks in the group have finished.")

anyio.run(main)
```

### Structured Concurrency & Fail-Fast

If any task in a `TaskGroup` raises an exception, AnyIO immediately:
1. Cancels all other tasks in that group.
2. Waits for those tasks to finish (cleanup).
3. Raises an `ExceptionGroup` (or a single exception if only one failed).

```python
async def failing_task():
    await anyio.sleep(1)
    raise ValueError("Boom!")

async def long_task():
    try:
        await anyio.sleep(10)
    except anyio.get_cancelled_exc_class():
        print("Long task was cancelled because a sibling failed!")
        raise

async def main():
    try:
        async with anyio.create_task_group() as tg:
            tg.start_soon(failing_task)
            tg.start_soon(long_task)
    except* ValueError as eg:
        print(f"Caught expected error: {eg.exceptions}")

anyio.run(main)
```

### Waiting for Initialization (`tg.start`)

Sometimes you need a task to perform some setup (like starting a server) before the main code continues. `tg.start()` allows the task to signal when it has reached a "ready" state.

```python
async def my_service(task_status=anyio.TASK_STATUS_IGNORED):
    # Setup
    await anyio.sleep(1)
    # Signal that we are ready
    task_status.started("Service Payload")
    
    # Main loop
    await anyio.sleep(5)

async def main():
    async with anyio.create_task_group() as tg:
        # Blocks until task_status.started() is called
        value = await tg.start(my_service)
        print(f"Service started with: {value}")
        # Now we can safely interact with the service
```

## Comparison: AnyIO vs. standard asyncio

| Feature | `asyncio.gather` | AnyIO `TaskGroup` |
| :--- | :--- | :--- |
| **Structure** | Unstructured (tasks can "leak") | Structured (tasks bound to scope) |
| **Failure** | Sibling tasks keep running | **Fail-fast**: Sibling tasks cancelled |
| **Exceptions** | Usually only first exception | Collects all in `ExceptionGroup` |
| **Cancellations** | Manual management | Automatic and clean |

!!! note "Python 3.11+ asyncio.TaskGroup"
    Python 3.11 introduced `asyncio.TaskGroup` which is very similar to AnyIO's version. However, AnyIO remains relevant for cross-backend compatibility and for users on older Python versions (it backports `ExceptionGroup` behavior).

## Best Practices

1. **Prefer `TaskGroup` over `gather`**: Even if using pure asyncio, the structured approach is safer.
2. **Use `start_soon` for independent tasks**: When you don't need a specific sync point.
3. **Use `start` for dependencies**: When the next step depends on the task being initialized.
4. **Handle `ExceptionGroup`**: Use the `except*` syntax (Python 3.11+) or `ExceptionGroup` backport to handle multiple failures gracefully.
