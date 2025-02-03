## Else for loops
>In a for or while loop the break statement may be paired with an else clause. If the loop finishes without executing the break, the else clause executes. This can be useful for error handling or to run code only if a loop completes successfully.

```python
for n in range(2, 10):
    for x in range(2, n):
        if n % x == 0:
            print(n, 'equals', x, '*', n//x)
            break
    else:
        # loop fell through without finding a factor
        print(n, 'is a prime number')
```

Output:
```
2 is a prime number
3 is a prime number
4 equals 2 * 2
5 is a prime number
...
```

