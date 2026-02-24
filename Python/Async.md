# Async

## Python Async vs Threads

### Threads

The OS runs multiple threads, switching between them automatically (preemptive multitasking). Each thread has its own stack and runs independently.

```python
import threading
import requests

def fetch(url):
    response = requests.get(url)
    print(response.status_code)

threads = [threading.Thread(target=fetch, args=(url,)) for url in urls]
for t in threads: t.start()
for t in threads: t.join()
```

### Async

```python
import asyncio
import aiohttp

async def fetch(url):
    async with aiohttp.ClientSession() as session:
        response = await session.get(url)
        print(response.status)

asyncio.run(asyncio.gather(*[fetch(url) for url in urls]))
```

| Feature | Threads | Async |
|---|---|---|
| Switching control | OS decides (preemptive) | You decide (cooperative) |
| Execution model | Multiple threads | Single thread, event loop |
| Memory per task | ~1MB per thread | ~few KB per coroutine |
| Max concurrent tasks | ~hundreds | ~tens of thousands |
| Shared state safety |  Race conditions possible |  Safer by default |
| Python GIL impact |  Released during I/O |  Not an issue |
| CPU-bound work |  Still GIL-limited |  Blocks event loop |
| Complexity | Low (familiar model) | Medium (async/await everywhere) |
| Debugging | Harder (non-deterministic) | Easier (deterministic flow) |


## Task Groups

`asyncio.TaskGroup` is a modern feature introduced in Python 3.11 to manage asynchronous tasks. It provides a safer, cleaner way to run multiple tasks concurrently compared to the older methods like `asyncio.gather()`.

It implements the concept of Structured Concurrency.

- **Scoping**: All tasks are spawned inside a async with block.
- **Waiting**: The block will not exit until all tasks inside it have finished.
- **Safety**: If one task crashes (raises an exception), the TaskGroup immediately cancels all other running tasks in the group and raises an exception.
