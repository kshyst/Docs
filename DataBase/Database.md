# Database

## What is Database? 

Database is a collection of data. It is used to store digital data in a structured format. Many modern applications use databases to store data.
They are managed using database management systems (DBMS).

## Types of Database

### Hierarchical databases

In this type , data is organized in a tree-liked structure.

This structure shows parent-child relationship between data where each parent can have multiple children but each child has only one parent.

![Hierarchical databases](https://media.geeksforgeeks.org/wp-content/uploads/20200414091301/Hierarchical-DB.png)

The Disadvantages are : less flexibility and harder to scale. adding new data often takes time to traverse the tree.

### Network databases

It's like Hierarchical databases but with more flexibility. It allows each child to have multiple parents. It's like a huge graph. 

It is effective for many-to-many relationships.


![Network databases](https://media.geeksforgeeks.org/wp-content/uploads/20200414091455/Network-DB.png)

The Disadvantages are : It's highly dependent on a pre-defined schema and it's hard to maintain. making changes will be time consuming. 


### Object-Oriented Databases

It's based on princibles of object-oriented programming. It stores data in the form of objects. These objects include attrriutes and methods. This make them easily referenced and manipulated.

![Object-Oriented Databases](https://media.geeksforgeeks.org/wp-content/uploads/20200414093006/Object-Oriented-DB.png)

This approach reduces workload on database by allowing objects to be reused. and linked directly, streamlining data access and manipulation. Each object behaves as an instance of the database model.

### Relational Databases

In this type, data is stored in tables. Each table has rows and columns. Each row is a record and each column is a field.

![Relational Databases](https://media.geeksforgeeks.org/wp-content/uploads/20200508135718/Capture3231.png)

Every row in a table is uniquely identified by a primary key. This key is used to establish a relationship between tables.

Similiarly, foreign keys are used to link tables. This allows data to be distributed and organized in a way that makes sense.

### NoSQL Databases

A NoSQL database (short for “non-SQL” or “non-relational”) provides a mechanism for storing and retrieving data that does not rely on traditional table-based relational models. Instead, it uses flexible data models like key-value pairs, documents, column families, or graphs, making it ideal for handling unstructured, semi-structured, and structured data.

