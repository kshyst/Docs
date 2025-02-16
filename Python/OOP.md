# Object-Oriented Programming in Python

## Inheritance

### Single Inheritance

In single inheritance, a class inherits from a single parent class.

```python

class Parent:
    def __init__(self, name):
        self.name = name

    def greet(self):
        return f'Hello, {self.name}'

class Child(Parent):

    def greet(self):
        return f'Hi, {self.name}'

parent = Parent('Alice')
child = Child('Bob')
print(parent.greet())  # Output: Hello, Alice
print(child.greet())  # Output: Hi, Bob

```

### Multiple Inheritance

In multiple inheritance, a class inherits from multiple parent classes.

```python

class Mother:
    def __init__(self, mother_name):
        self.mother_name = mother_name

    def greet(self):
        return f'Hello, I am {self.mother_name}'

class Father:

    def __init__(self, father_name):
        self.father_name = father_name

    def greet(self):
        return f'Hi, I am {self.father_name}'

class Child(Mother, Father):

    def greet(self):
        return f'Hi, I am {self.mother_name} {self.father_name}'

child = Child('Alice', 'Bob')
print(child.greet())  # Output: Hi, I am Alice Bob

```

### Multilevel Inheritance

In multilevel inheritance, a class inherits from a parent class, which in turn inherits from another parent class.

```python

class Grandparent:
    def __init__(self, grandparent_name):
        self.grandparent_name = grandparent_name

    def greet(self):
        return f'Hello, I am {self.grandparent_name}'

class Parent(Grandparent):

    def __init__(self, parent_name, grandparent_name):
        super().__init__(grandparent_name)
        self.parent_name = parent_name

    def greet(self):
        return f'Hi, I am {self.parent_name}'

class Child(Parent):

    def __init__(self, child_name, parent_name, grandparent_name):
        super().__init__(parent_name, grandparent_name)
        self.child_name = child_name

    def greet(self):
        return f'Hi, I am {self.child_name} {self.parent_name} {self.grandparent_name}'

child = Child('Alice', 'Bob', 'Charlie')
print(child.greet())  # Output: Hi, I am Alice Bob Charlie
```

### Hierarchical Inheritance

In hierarchical inheritance, multiple classes inherit from a single parent class.

```python

class Parent:
    def __init__(self, parent_name):
        self.parent_name = parent_name

    def greet(self):
        return f'Hello, I am {self.parent_name}'

class Child1(Parent):

    def greet(self):
        return f'Hi, I am {self.parent_name}'

class Child2(Parent):

    def greet(self):
        return f'Hi, I am {self.parent_name}'

child1 = Child1('Alice')

child2 = Child2('Bob')

print(child1.greet())  # Output: Hi, I am Alice

print(child2.greet())  # Output: Hi, I am Bob

```

## Hybrid Inheritance

Hybrid inheritance is a combination of multiple inheritance and hierarchical inheritance.

```python

class Grandparent:
    def __init__(self, grandparent_name):
        self.grandparent_name = grandparent_name

    def greet(self):
        return f'Hello, I am {self.grandparent_name}'

class Parent1(Grandparent):

    def __init__(self, parent1_name, grandparent_name):
        super().__init__(grandparent_name)
        self.parent1_name = parent1_name

    def greet(self):
        return f'Hi, I am {self.parent1_name}'

class Parent2(Grandparent):

    def __init__(self, parent2_name, grandparent_name):
        super().__init__(grandparent_name)
        self.parent2_name = parent2_name

    def greet(self):
        return f'Hi, I am {self.parent2_name}'

class Child(Parent1, Parent2):

    def __init__(self, child_name, parent1_name, parent2_name, grandparent_name):
        Parent1.__init__(self, parent1_name, grandparent_name)
        Parent2.__init__(self, parent2_name, grandparent_name)
        self.child_name = child_name

    def greet(self):
        return f'Hi, I am {self.child_name} {self.parent1_name} {self.parent2_name} {self.grandparent_name}'

child = Child('Alice', 'Bob', 'Charlie', 'David')
print(child.greet())  # Output: Hi, I am Alice Bob Charlie David
```

## Polymorphism

Polymorphism is the ability of an object to take on multiple forms. In Python, polymorphism allows objects of different classes to be treated as objects of a common superclass.
Polymorphism allows methods to have the same name but behave differently based on the object’s context. It can be achieved through method overriding or overloading.

### Type of Polymorphism

1. Compile-time Polymorphism (Method Overloading)
2. Run-time Polymorphism (Method Overriding)

## Encapsulation

Encapsulation is the process of wrapping data (variables) and methods (functions) into a single unit called a class. It restricts direct access to some of the object's components, which prevents the accidental modification of data.
A class is an example of encapsulation as it encapsulates all the data that is member functions, variables, etc.

### Private Variables

In Python, private variables can be created by prefixing the variable name with two underscores (`__`). This makes the variable name mangled, which means it can't be accessed directly from outside the class.

```python

class Person:
    def __init__(self, name, age):
        self.__name = name
        self.__age = age

    def display(self):
        return f'Name: {self.__name}, Age: {self.__age}'

person = Person('Alice', 30)
print(person.display())  # Output: Name: Alice, Age: 30
print(person.__name)  # AttributeError: 'Person' object has no attribute '__name'
```

### Public Variables

Public variables are accessible from outside the class. In Python, all variables are public by default.

### Protected Variables

In Python, protected variables can be created by prefixing the variable name with a single underscore (`_`). This indicates that the variable should not be accessed directly from outside the class. But can be accessed from a subclass.

```python

class Person:
    def __init__(self, name, age):
        self._name = name
        self._age = age

    def display(self):
        return f'Name: {self._name}, Age: {self._age}'
```

## Abstraction
Abstraction hides the internal implementation details while exposing only the necessary functionality. It helps focus on “what to do” rather than “how to do it.”

### Types of Abstraction

1. Partial Abstraction - Abstract class contains both abstract and concrete methods.
2. Full Abstraction - Abstract class contains only abstract methods.

```python
from abc import ABC, abstractmethod

class Dog(ABC):  # Abstract Class
    def __init__(self, name):
        self.name = name

    @abstractmethod
    def sound(self):  # Abstract Method
        pass

    def display_name(self):  # Concrete Method
        print(f"Dog's Name: {self.name}")

class Labrador(Dog):  # Partial Abstraction
    def sound(self):
        print("Labrador Woof!")

class Beagle(Dog):  # Partial Abstraction
    def sound(self):
        print("Beagle Bark!")

# Example Usage
dogs = [Labrador("Buddy"), Beagle("Charlie")]
for dog in dogs:
    dog.display_name()  # Calls concrete method
    dog.sound()  # Calls implemented abstract method
```

Why Use It: Abstraction ensures consistency in derived classes by enforcing the implementation of abstract methods.
