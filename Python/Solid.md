# Solid Principles


The SOLID principles are five essential guidelines that enhance software design, making code more maintainable and scalable.

The SOLID principle helps in reducing tight coupling. Tight coupling means a group of classes are highly dependent on one another which you should avoid in your code.

Opposite of tight coupling is loose coupling and your code is considered as a good code when it has loosely-coupled classes.
Loosely coupled classes minimize changes in your code, helps in making code more reusable, maintainable, flexible and stable. 

## Single Responsibility Principle (SRP)

This principle states that “A class should have only one reason to change” which means every class should have a single responsibility or single job or single purpose. In other words, a class should have only one job or purpose within the software system.

```text
Imagine a baker who is responsible for baking bread. The baker’s role is to focus on the task of baking bread, ensuring that the bread is of high quality, properly baked, and meets the bakery’s standards.
However, if the baker is also responsible for managing the inventory, ordering supplies, serving customers, and cleaning the bakery, this would violate the SRP.
```

```python
class BreadBaker
class InventoryManager
class SupplyOrderer
class CustomerService
...
```

## Open/Closed Principle (OCP)

This principle states that “Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification” which means you should be able to extend a class behavior, without modifying it.


 Imagine you have a class called PaymentProcessor that processes payments for an online store. Initially, the PaymentProcessor class only supports processing payments using credit cards. However, you want to extend its functionality to also support processing payments using PayPal.

Instead of modifying the existing PaymentProcessor class to add PayPal support, you can create a new class called PayPalPaymentProcessor that extends the PaymentProcessor class. This way, the PaymentProcessor class remains closed for modification but open for extension, adhering to the Open-Closed Principle. Let’s understand this through the code implementation.

```python
class PaymentProcessor:
    def process_payment(self, payment_details):
        pass

class PayPalPaymentProcessor(PaymentProcessor):

    def process_payment(self, payment_details):
        # Process payment using PayPal
        pass
```

## Liskov Substitution Principle (LSP)

According to this principle “Derived or child classes must be substitutable for their base or parent classes“. This principle ensures that any class that is the child of a parent class should be usable in place of its parent without any unexpected behavior.

One of the classic examples of this principle is a rectangle having four sides. A rectangle’s height can be any value and width can be any value. A square is a rectangle with equal width and height. So we can say that we can extend the properties of the rectangle class into square class. 

```python

class Rectangle:
    def __init__(self, width, height):
        self._width = width
        self._height = height

    def set_width(self, width):
        self._width = width

    def set_height(self, height):
        self._height = height

    def area(self):
        return self._width * self._height

class Square(Rectangle):

    def set_width(self, width):
        self._width = width
        self._height = width

    def set_height(self, height):
        self._width = height
        self._height = height

```

- Rectangle Class: This is the base class that has properties for width and height. It has methods for calculating the area and for setting width and height.
- Square Class: This class inherits from Rectangle but overrides the setWidth and setHeight methods to ensure that changing one dimension affects the other, maintaining the property that all sides are equal.

## Interface Segregation Principle (ISP)

This principle is the first principle that applies to Interfaces instead of classes in SOLID and it is similar to the single responsibility principle. It states that “do not force any client to implement an interface which is irrelevant to them“. Here your main goal is to focus on avoiding fat interface and give preference to many small client-specific interfaces. You should prefer many client interfaces rather than one general interface and each interface should have a specific responsibility.

Suppose if you enter a restaurant and you are pure vegetarian. The waiter in that restaurant gave you the menu card which includes vegetarian items, non-vegetarian items, drinks, and sweets. 

In this case, as a customer, you should have a menu card which includes only vegetarian items, not everything which you don’t eat in your food. Here the menu should be different for different types of customers.
The common or general menu card for everyone can be divided into multiple cards instead of just one. Using this principle helps in reducing the side effects and frequency of required changes.

```C++

#include <iostream>
#include <vector>
#include <string>

// Interface for vegetarian menu
class IVegetarianMenu {
public:
    virtual std::vector<std::string> getVegetarianItems() = 0;
};

// Interface for non-vegetarian menu
class INonVegetarianMenu {
public:
    virtual std::vector<std::string> getNonVegetarianItems() = 0;
};

// Interface for drinks menu
class IDrinkMenu {
public:
    virtual std::vector<std::string> getDrinkItems() = 0;
};

// Class for vegetarian menu
class VegetarianMenu : public IVegetarianMenu {
public:
    std::vector<std::string> getVegetarianItems() override {
        return {"Vegetable Curry", "Paneer Tikka", "Salad"};
    }
};

// Class for non-vegetarian menu
class NonVegetarianMenu : public INonVegetarianMenu {
public:
    std::vector<std::string> getNonVegetarianItems() override {
        return {"Chicken Curry", "Fish Fry", "Mutton Biryani"};
    }
};

// Class for drinks menu
class DrinkMenu : public IDrinkMenu {
public:
    std::vector<std::string> getDrinkItems() override {
        return {"Water", "Soda", "Juice"};
    }
};

// Function to display menu items for a vegetarian customer
void displayVegetarianMenu(IVegetarianMenu* menu) {
    std::cout << "Vegetarian Menu:\n";
    for (const auto& item : menu->getVegetarianItems()) {
        std::cout << "- " << item << std::endl;
    }
}

// Function to display menu items for a non-vegetarian customer
void displayNonVegetarianMenu(INonVegetarianMenu* menu) {
    std::cout << "Non-Vegetarian Menu:\n";
    for (const auto& item : menu->getNonVegetarianItems()) {
        std::cout << "- " << item << std::endl;
    }
}

int main() {
    VegetarianMenu vegMenu;
    NonVegetarianMenu nonVegMenu;
    DrinkMenu drinkMenu;

    displayVegetarianMenu(&vegMenu);
    displayNonVegetarianMenu(&nonVegMenu);

    return 0;
}
```

## Dependency Inversion Principle (DIP)

The Dependency Inversion Principle (DIP) is a principle in object-oriented design that states that “High-level modules should not depend on low-level modules. Both should depend on abstractions“. Additionally, abstractions should not depend on details. Details should depend on abstractions.

In simpler terms, the DIP suggests that classes should rely on abstractions (e.g., interfaces or abstract classes) rather than concrete implementations.
This allows for more flexible and decoupled code, making it easier to change implementations without affecting other parts of the codebase.

In a software development team, developers depend on an abstract version control system (e.g., Git) to manage and track changes to the codebase. They don’t depend on specific details of how Git works internally. 
This allows developers to focus on writing code without needing to understand the intricacies of version control implementation. Below is the code of above example to understand Dependency Inversion Principle:

```cpp

#include <iostream>
#include <string>

// Interface for version control system
class IVersionControl {
public:
    virtual void commit(const std::string& message) = 0;
    virtual void push() = 0;
    virtual void pull() = 0;
};

// Git version control implementation
class GitVersionControl : public IVersionControl {
public:
    void commit(const std::string& message) override {
        std::cout << "Committing changes to Git with message: " << message << std::endl;
    }

    void push() override {
        std::cout << "Pushing changes to remote Git repository." << std::endl;
    }

    void pull() override {
        std::cout << "Pulling changes from remote Git repository." << std::endl;
    }
};

// Team class that relies on version control
class DevelopmentTeam {
private:
    IVersionControl* versionControl;

public:
    DevelopmentTeam(IVersionControl* vc) : versionControl(vc) {}

    void makeCommit(const std::string& message) {
        versionControl->commit(message);
    }

    void performPush() {
        versionControl->push();
    }

    void performPull() {
        versionControl->pull();
    }
};

int main() {
    GitVersionControl git;
    DevelopmentTeam team(&git);

    team.makeCommit("Initial commit");
    team.performPush();
    team.performPull();

    return 0;
}

```
- IVersionControl Interface: This defines the operations that any version control system should support, like commit, push, and pull. It serves as an abstraction that decouples high-level code from low-level implementations.
- GitVersionControl Class: This class implements the IVersionControl interface, providing specific functionality for managing version control using Git.
- DevelopmentTeam Class: This class relies on the IVersionControl interface, meaning it can work with any version control implementation that adheres to the interface. It does not need to know the details of how Git works internally.