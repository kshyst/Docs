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


## Is and == difference
> - is is used to compare the identity of two objects, == is used to compare the equality of two objects.
> - is is reference equality, == is value equality

## Keyword arguments
> - Functions can also be called using keyword arguments of the form kwarg=value. For instance, the following function:
> - accepts one required argument (voltage) and two optional arguments (state, action). This function can be called in any of the following ways:

```python
def parrot(voltage, *args , **kwargs):
    print("voltage: ", voltage)
    for arg in args:
    print(arg)
    for key in kwargs:
    print(key, "is" , kwargs[key])

parrot(1000 , 'a' , 'b' , 'c' , state='happy' , action='jump')
parrot(1000 , 'a' , 'b' , 'c' , action='jump' , state='happy')
parrot(1000 , 'a' , 'b' , 'c' , state='happy')
```

## Special parameters

> By default, arguments may be passed to a Python function either by position or explicitly by keyword. For readability and performance, it makes sense to restrict the way arguments can be passed to a function. Python’s parameter passing mechanism has two special forms:

- Positional-only arguments (/):

>These arguments must be passed by position. You cannot pass them by keyword.
>The / symbol in the function definition separates positional-only parameters from others.

-Positional or Keyword arguments:

>These arguments can be passed either by position or by keyword. This is the default behavior for most parameters.
>This is what happens when you don't use any special symbols (just regular parameters).

- Keyword-only arguments (*):

>These arguments must be passed explicitly by keyword. You cannot pass them by position.
>The * symbol in the function definition marks the start of keyword-only parameters.

```python
def f(a, b, /, c, d, *, e, f):
    print(a, b, c, d, e, f)

f(10, 20, 30, d=40, e=50, f=60)
f(10, 20, c=30, d=40, e=50, f=60)
```

## Arbitrary Argument Lists
> You can use *args to pass arbitrary number of arguments to a function. Here is an example:

```python
def write_multiple_items(file, separator, *args):
    file.write(separator.join(args))
```

## Unpacking Argument Lists

```python
def parrot(voltage, state='a stiff', action='voom'):
    print("-- This parrot wouldn't", action, end=' ')
    print("if you put", voltage, "volts through it.", end=' ')
    print("E's", state, "!")

d = {"voltage": "four million", "state": "bleedin' demised", "action": "VOOM"}
parrot(**d)
```

## List Comprehensions

> List comprehensions provide a concise way to create lists. Common applications are to make new lists where each element is the result of some operation applied to each member of another sequence or iterable, or to create a subsequence of those elements that satisfy a certain condition.

```python
squares = []
for x in range(10):
    squares.append(x**2)

squares = list(map(lambda x: x**2, range(10)))

squares = [x**2 for x in range(10)]
```

## About Tuple
> This code is tuple ! :

```python
singleton = 'hello',
```

> Can be stored reversely:

```python
z , y , x = tuple1
```

## About Sets

> - unordered collection of unique elements
> - eliminate duplicate entries
> - supports mathematical operations like union, intersection, difference, and symmetric difference


## About Dictionaries

> - key must be immutable object and unique
> - strings and numbers can always be keys also
> - list(dict) returns a list of all the keys used in the dictionary, in insertion order (if you want it sorted, just use sorted(dict) instead).
> - convert list of paired values using dict() constructor :

```python
dict([('sape', 4139), ('guido', 4127), ('jack', 4098)])
```