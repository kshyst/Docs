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

## Why Use *args and **kwargs?

> - Flexibility: *args and **kwargs allow you to pass a variable number of arguments to a function. This can be useful when you don't know in advance how many arguments will be passed to the function.
> - Decorators: *args and **kwargs are commonly used in decorators to pass arguments to the decorated function without knowing the exact number of arguments in advance.

## Comprehensions

> Comprehensions are constructs that allow sequences to be built from other sequences.

### List Comprehensions

> List comprehensions provide a concise way to create lists. Common applications are to make new lists where each element is the result of some operation applied to each member of another sequence or iterable, or to create a subsequence of those elements that satisfy a certain condition.

#### Blueprint

```python
variable = [out_exp for out_exp in input_list if out_exp == 2]
```

```python
squares = []
for x in range(10):
    squares.append(x**2)

squares = list(map(lambda x: x**2, range(10)))

squares = [x**2 for x in range(10)]
```

### Dictionary Comprehensions

> Dictionary comprehensions are similar, but they construct dictionaries instead of lists. A dictionary comprehension looks like this:

#### Blueprint
```python
{v: k for k, v in some_dict.items()}
```

```python
squares = {x: x**2 for x in range(10)}
```

> another example:

```python
mcase = {'a': 10, 'b': 34, 'A': 7, 'Z': 3}

mcase_frequency = {
    k.lower(): mcase.get(k.lower(), 0) + mcase.get(k.upper(), 0)
    for k in mcase.keys()
}

# mcase_frequency == {'a': 17, 'z': 3, 'b': 34}
```

### Set Comprehensions

> They are also similar to list comprehensions. The only difference is that they use braces {}. Here is an example:

```python
squared = {x**2 for x in [1, 1, 2]}
print(squared)
# Output: {1, 4}
```

### Generator Comprehensions

> Generator comprehensions are similar to list comprehensions, but they return a generator object instead of a list. This means that they generate values on-the-fly instead of storing them in memory. Generator comprehensions are created using parentheses () instead of square brackets [].

```python
gen = (x**2 for x in range(100))
print(gen)

for x in gen:
    print(x)
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

## Some Looping Techniques

> - use items() method to loop over key-value pairs :

```python
knights = {'gallahad': 'the pure', 'robin': 'the brave'}
for k, v in knights.items():
    print(k, v)
```

> - use enumerate() to get index and value of list : 

```python
for i, v in enumerate(['tic', 'tac', 'toe']):
    print(i, v)
```

> - use zip() to loop over two or more sequences at the same time :

```python
questions = ['name', 'quest', 'favorite color']
answers = ['lancelot', 'the holy grail', 'blue']
for q, a in zip(questions, answers):
    print('What is your {0}?  It is {1}.'.format(q, a))
```

> - use reversed() to loop over a sequence in reverse :

```python
for i in reversed(range(1, 10, 2)):
    print(i)
```

> - use sorted() to loop over a sequence in sorted order and set() to eliminate duplicates :

```python
basket = ['apple', 'orange', 'apple', 'pear', 'orange', 'banana']
for f in sorted(set(basket)):
    print(f)
```

## More on Conditions
> - short circuit operations are "and" and "or" , they evaluate from left to right and return the value of the last expression evaluated
> - "not" has the highest priority among them
> - walrus operator := assigns values to variables as part of a larger expression

```python
# walrus example
n = 0
while (n := n + 1) < 10:
    print(n)

# is equivalent to
    
n = 0
while n < 10:
    n += 1
    print(n)

```

## Modules

> - Can use global variables without interrupting users variables
> - "from fibo import * " imports all names except those beginning with an underscore (_) , "from fibo import fib" imports only fib

### about __name__ == "__main__"
>That means that by adding this code at the end of your module:

```python
     if __name__ == "__main__":
         import sys
         fib(int(sys.argv[1]))
```
> you can make the file usable as a script as well as an importable module, because the code that parses the command line only runs if the module is executed as the “main” file:

```python
python fibo.py 50
0 1 1 2 3 5 8 13 21 34
```

### Importing with *
> - import * does not import items whose name starts with an underscore
> - f a package’s __init__.py code defines a list named __all__, it is taken to be the list of module names that should be imported when from package import * is encountered.

```python
__all__ = ["echo", "surround", "reverse"]
```

> Be aware that submodules might become shadowed by locally defined names. For example, if you added a reverse function to the sound/effects/__init__.py file, the from sound.effects import * would only import the two submodules echo and surround, but not the reverse submodule, because it is shadowed by the locally defined reverse function:

```python
__all__ = [
    "echo",      # refers to the 'echo.py' file
    "surround",  # refers to the 'surround.py' file
    "reverse",   # !!! refers to the 'reverse' function now !!!
]

def reverse(msg: str):  # <-- this name shadows the 'reverse.py' submodule
    return msg[::-1]    #     in the case of a 'from sound.effects import *'
```

> If \_\_all__ is not defined, the statement from sound.effects import * does not import all submodules from the package sound.effects into the current namespace; it only ensures that the package sound.effects has been imported (possibly running any initialization code in __init__.py) and then imports whatever names are defined in the package. This includes any names defined (and submodules explicitly loaded) by __init__.py. It also includes any submodules of the package that were explicitly loaded by previous import statements.

> Example :

```python
import sound.effects.echo
import sound.effects.surround
from sound.effects import *
```

### Module search path
> - searches the current directory
> - looks in PYTHONPATH
> - looks in installation-dependent default

> After initialization, Python programs can modify sys.path. The directory containing the script being run is placed at the beginning of the search path, ahead of the standard library path. This means that scripts in that directory will be loaded instead of modules of the same name in the library directory.


## Built-in Exceptions List

> - BaseException : the base class for all built-in exceptions
> - SystemExit : raised by sys.exit() function
> - KeyboardInterrupt : raised when the user hits the interrupt key (Ctrl+C or Delete)
> - Exception : the base class for all built-in exceptions
> - StopIteration : raised by next() function to indicate that there is no further item to be returned by iterator
> - GeneratorExit : raised when a generator is closed
> - SystemExit : raised by sys.exit() function


## About Classes

> - Execution of a derived class definition proceeds the same as for a base class. When the class object is constructed, the base class is remembered. This is used for resolving attribute references: if a requested attribute is not found in the class, the search proceeds to look in the base class. This rule is applied recursively if the base class itself is derived from some other class
> - Derived classes may override methods of their base classes. Because methods have no special privileges when calling other methods of the same object, a method of a base class that calls another method defined in the same base class may end up calling a method of a derived class that overrides it. (For C++ programmers: all methods in Python are effectively virtual.)

### Multiple Inheritance

> - Python supports a form of multiple inheritance as well. A class definition with multiple base classes looks like this:

```python
class DerivedClassName(Base1, Base2, Base3):
    <statement-1>
    .
    .
    .
    <statement-N>
```

> - If a requested attribute is not found in the class, the search 
> - proceeds to look in the base classes. This rule is applied recursively if the base class itself is derived from some other class.

## Iterators

> implementing iterator in a class :

```python
class Reverse:
    """Iterator for looping over a sequence backwards."""
    def __init__(self, data):
        self.data = data
        self.index = len(data)

    def __iter__(self):
        return self

    def __next__(self):
        if self.index == 0:
            raise StopIteration
        self.index = self.index - 1
        return self.data[self.index]
```

### iter() and next() functions

> - iter() returns an iterator object that produces values from the iterable.

### Difference between iterator , iterable and iteration

> - Iterable: An object capable of returning its members one at a time. Examples of iterables include all sequence types (such as list, str, and tuple) and some non-sequence types like dict, file objects, and objects of any classes you define with an __iter__() or __getitem__() method.
> - Iterator: An object representing a stream of data. Repeated calls to the iterator’s __next__() method return successive items in the stream. When no more data is available, a StopIteration exception is raised.
> - Iteration: The process of taking an item from an iterable. When using a for loop, the process of iteration is hidden from the user.
> - Generators are iterators, but you can only iterate over them once. It’s because they do not store all the values in memory, they generate the values on the fly. You use them by iterating over them, either with a ‘for’ loop or by passing them to any function or construct that iterates. Most of the time generators are implemented as functions. However, they do not return a value, they yield it


## Generators

> Generators are a simple and powerful tool for creating iterators. They are written like regular functions but use the yield statement whenever they want to return data. Each time next() is called on it, the generator resumes where it left off (it remembers all the data values and which statement was last executed).

```python

def reverse(data):
    for index in range(len(data)-1, -1, -1):
        yield data[index]

for char in reverse('golf'):
    print(char)
```

> - What makes generators so compact is that the __iter__() and __next__() methods are created automatically.
> - Another key feature is that the local variables and execution state are automatically saved between calls. This made the function easier to write and much more clear than an approach using instance variables like self.index and self.data.
> - In addition to automatic method creation and saving program state, when generators terminate, they automatically raise StopIteration.
> - Generators are best for calculating large sets of results (particularly calculations involving loops themselves) where you don’t want to allocate the memory for all results at the same time

## Class Methods and Static Methods
#### class methods
> - class methods are methods that are not bound to an object, but to a class. They can be called on both the class and its objects.
> - class methods are defined using the @classmethod decorator.
> - class methods take a cls parameter that points to the class and not the object instance when the method is called.
> - class methods can modify the class state that applies across all instances of the class.
> - class methods are used in factory methods, where we want to return an instance of the class.
> - class methods are used as alternative constructors.
> - class methods are used when we need to access or modify class state.

#### static methods
> - static methods are methods that are not bound to an object or class. They are defined using the @staticmethod decorator.
> - static methods do not take cls or self as the first parameter.
> - static methods are used when a method does not access or modify the class state or object state.
> - static methods are used when we need a utility method that does not access or modify the class state or object state.

```python

class Person:
    count = 0

    def __init__(self, first_name, last_name):
        self.first_name = first_name
        self.last_name = last_name
        Person.count += 1

    @classmethod
    def count_instances(cls):
        return f"Created {cls.count} instances of Person class"

    @staticmethod
    def hello():
        return "Hello, static method called"

p1 = Person("John", "Doe")
p2 = Person("Jane", "Smith")

print(Person.count_instances())
print(Person.hello())
```

## Decorators

> In Python, decorators are a powerful and flexible way to modify or extend the behavior of functions or methods, without changing their actual code. A decorator is essentially a function that takes another function as an argument and returns a new function with enhanced functionality.

#### Explanation of Parameters

> 1. decorator_name(func):
> 
> decorator_name: This is the name of the decorator function.
> func: This parameter represents the function being decorated. When you use a decorator, the decorated function is passed to this parameter.
> 2. wrapper(*args, **kwargs):
> 
> wrapper: This is a nested function inside the decorator. It wraps the original function, adding additional functionality.
> *args: This collects any positional arguments passed to the decorated function into a tuple.
> **kwargs: This collects any keyword arguments passed to the decorated function into a dictionary.
> The wrapper function allows the decorator to handle functions with any number and types of arguments.
> 3. @decorator_name:
> 
> This syntax applies the decorator to the function_to_decorate function. It is equivalent to writing function_to_decorate = decorator_name(function_to_decorate).

#### Example

```python
def my_decorator(func):
    def wrapper(*args, **kwargs):
        print("Something is happening before the function is called.")
        result = func(*args, **kwargs)
        print("Something is happening after the function is called.")
        return result
    return wrapper

@my_decorator
def say_hello(name):
    print(f"Hello, {name}!")

say_hello("Alice")
```

#### Output

```text
Something is happening before the function is called.
Hello, Alice!
Something is happening after the function is called.
```

> Decorating a function changes the function's identity to the wrapper function.
> 
> To fix this problem python provides a tool called __funtools.wraps__ to fix this issue.


```python
from functools import wraps

def my_decorator(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        print("Something is happening before the function is called.")
        result = func(*args, **kwargs)
        print("Something is happening after the function is called.")
        return result
    return wrapper

@my_decorator
def say_hello(name):
    print(f"Hello, {name}!")
    
print(say_hello.__name__)

```

### Different Types of Decorators

> - Function Decorators: These are the most common type of decorators in Python. They are applied directly to functions using the @decorator_name syntax.
> - Class Decorators: These decorators are applied to classes. They are defined using the @decorator_name syntax before the class definition.
> - Method Decorators: These decorators are applied to methods within a class. They are defined using the @decorator_name syntax before the method definition.
> - Property Decorators: These decorators are used to define properties in classes. They are defined using the @property decorator before the method definition.
> - Static Method Decorators: These decorators are used to define static methods in classes. They are defined using the @staticmethod decorator before the method definition.
> - Class Method Decorators: These decorators are used to define class methods in classes. They are defined using the @classmethod decorator before the method definition.
> - Decorators with Arguments: Decorators can also take arguments. In this case, you need to define an additional function that returns the actual decorator. This allows you to pass arguments to the decorator function.



## Format function

> use format function to pass variables the hard way into a string

```python
    print("Hello, {0}! You are {1} years old.".format(name, age))
```

## f-strings

> f-strings are a more concise and readable way to format strings in Python. They are prefixed with an 'f' or 'F' and are written as normal strings with curly braces {} containing expressions that will be replaced with their values.

```python
    print(f"Hello, {name}! You are {age} years old.")
```

## Monkey Patching

> Monkey patching is a technique used to dynamically modify or extend the behavior of a class or module at runtime. It involves changing the code of a module or class without changing the original source code. This can be useful for adding new features, fixing bugs, or testing code.

```python
class MyClass:
    def my_method(self):
        return "Original method"

def monkey_patch(self):
    return "Patched method"

obj = MyClass()

# Patching the method
MyClass.my_method = monkey_patch

print(obj.my_method())  # Output: Patched method
```


## Map Filters and Reduce

> - map() applies a function to all the items in an input_list : map(function_to_apply, list_of_inputs)
> - filter() creates a list of elements for which a function returns true.

### map()

```python
items = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, items))
```

> is equivalent to :

```python
items = [1, 2, 3, 4, 5]
squared = []
for i in items:
    squared.append(i**2)
```

> another example :

```python
def multiply(x):
    return (x*x)
def add(x):
    return (x+x)

funcs = [multiply, add]
for i in range(5):
    value = list(map(lambda x: x(i), funcs))
    print(value)
```

> Output :

```text
[0, 0]
[1, 2]
[4, 4]
[9, 6]
[16, 8]
```

### filter()

```python
number_list = range(-5, 5)
less_than_zero = list(filter(lambda x: x < 0, number_list))
print(less_than_zero)
```

> is equivalent to :

```python
number_list = range(-5, 5)
less_than_zero = []
for i in number_list:
    if i < 0:
        less_than_zero.append(i)
print(less_than_zero)
```

### reduce()

> Reduce is a really useful function for performing some computation on a list and returning the result. It applies a rolling computation to sequential pairs of values in a list. For example, if you wanted to compute the product of a list of integers.

```python
from functools import reduce
product = reduce((lambda x, y: x * y), [1, 2, 3, 4])
```

> is equivalent to :

```python

product = 1
for i in [1, 2, 3, 4]:
    product *= i
```

### Difference between map() and reduce()

> - map() applies a function to all the items in an input_list. A new list is returned which contains items returned by that function for each item.
> - reduce() applies a rolling computation to sequential pairs of values in a list. For example, if you wanted to compute the product of a list of integers, you could use reduce() to perform the computation.


## Python Debugger

> Python provides a built-in debugger module called pdb. It is a powerful tool for debugging and analyzing code. You can use it to set breakpoints, step through code, and inspect variables.

> To start the debugger, you can run your script with the -m pdb option:

```shell

python -m pdb my_script.py

```

> Once the debugger is running, you can use the following commands to control the debugger:
> - h or help: Display a list of available commands.
> - l or list: Show the current line of code being executed
> - n or next: Execute the current line of code and move to the next line
> - c or continue: Continue running the code until the next breakpoint is encountered
> - q or quit: Exit the debugger
> - p or print: Print the value of a variable
> - s or step: Step into a function call
> - r or return: Continue running the code until the current function returns
> - b or break: Set a breakpoint at a specific line number
> - w or where: Show the current call stack

## File Handling

> to make sure that the file gets closed whether an exception occurs or not, pack it into a with statement:

```python
with open('photo.jpg', 'r+') as f:
    jpgdata = f.read()
```

> - If you want to read the file, pass in r
> - If you want to read and write the file, pass in r+
> - If you want to overwrite the file, pass in w
> - If you want to append to the file, pass in a


