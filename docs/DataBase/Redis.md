# Redis

Redis is an open-source, advanced key-value store and an apt solution for building high-performance, scalable web applications. Redis has three main peculiarities that set it apart:

- Redis holds its database entirely in memory, using the disk only for persistence.
- Redis has a relatively rich set of data types when compared to many key-value data stores.
- Redis can replicate data to any number of slaves.

## Starting Redis

```bash
redis-server
redis-cli
```

## Commands Guide

### Retrieving All Values
The command to get "everything" from a key depends on its **Data Type**:

| Type | Command | Description |
| :--- | :--- | :--- |
| **String** | `GET key` | Returns the single value stored at the key. |
| **Hash** | `HGETALL key` | Returns all fields and values in the hash. |
| **List** | `LRANGE key 0 -1` | Returns all elements in the list (from first to last). |
| **Set** | `SMEMBERS key` | Returns all unique members of the set. |
| **Sorted Set** | `ZRANGE key 0 -1` | Returns all members in the sorted set, ordered by score. |

### General Key Operations
- `KEYS *`: List all keys in the database (⚠️ **Do not use in production**).
- `SCAN 0`: Iteratively list keys (Safer alternative to `KEYS`).
- `TYPE key`: Check the data type of a key.
- `EXISTS key`: Check if a key exists.

## Useful Links
- [Complete Guide to Redis Commands](https://www.geeksforgeeks.org/complete-guide-to-redis-commands/)
- [Redis Official Documentation](https://redis.io/commands/)

