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