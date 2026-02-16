# Task Groups

`asyncio.TaskGroup` is a modern feature introduced in Python 3.11 to manage asynchronous tasks. It provides a safer, cleaner way to run multiple tasks concurrently compared to the older methods like `asyncio.gather()`.

It implements the concept of Structured Concurrency.

- **Scoping**: All tasks are spawned inside a async with block.
- **Waiting**: The block will not exit until all tasks inside it have finished.
- **Safety**: If one task crashes (raises an exception), the TaskGroup immediately cancels all other running tasks in the group and raises an exception.