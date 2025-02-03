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

## About Functions
> - arguments are passed using call by value
> - The execution of a function introduces a new symbol table used for the local variables of the function. More precisely, all variable assignments in a function store the value in the local symbol table; whereas variable references first look in the local symbol table, then in the local symbol tables of enclosing functions, then in the global symbol table, and finally in the table of built-in names. Thus, global variables and variables of enclosing functions cannot be directly assigned a value within a function (unless, for global variables, named in a global statement, or, for variables of enclosing functions, named in a nonlocal statement), although they may be referenced.
> - Functions without return value , return None