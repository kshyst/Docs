# Entity Relationship Model

[Useful Website for ER Diagram](https://www.softwareideas.net/chen-er-diagram-erd)

It's a model used to describe the relationships between entities in a database. The ER data model specifies enterprise schema that represents the overall logical structure of a database graphically.

## Why use ER Model?

- Easy to be converted into tables. 
- Does't require any technical knowledge about the underlying DBMS.
- visualizing data logically.

## Symbols in ER Model

![ER Model Symbols](https://media.geeksforgeeks.org/wp-content/uploads/20230428115454/Introduction-to-ER-Model-2-768.webp)

- Rectangles: Rectangles represent Entities in the ER Model.
- Ellipses: Ellipses represent Attributes in the ER Model.
- Diamond: Diamonds represent Relationships among Entities.
- Lines: Lines represent attributes to entities and entity sets with other relationship types.
- Double Ellipse: Double Ellipses represent Multi-Valued Attributes.
- Double Rectangle: Double Rectangle represents a Weak Entity.

## Components of ER Model

- Entity: It's a real-world object that is distinguishable from other objects. It has attributes that describe it.
- Relationship: It's an association among two or more entities.
- Attribute: It's a property of an entity. It's a field that describes the entity.

![ER Model Components](https://media.geeksforgeeks.org/wp-content/uploads/20230428090323/Introduction-to-ER-Model-1-768.webp)

## Types of Relationshipsh 

- One-to-One: One entity is associated with only one entity.
- One-to-Many: One entity is associated with many entities.
- Many-to-One: Many entities are associated with one entity.
- Many-to-Many: Many entities are associated with many entities.

## Types of Entities

- Strong Entity: A Strong Entity is a type of entity that has a key Attribute. Strong Entity does not depend on other Entity in the Schema. It has a primary key, that helps in identifying it uniquely, and it is represented by a rectangle. These are called Strong Entity Types.
- Weak Entity: An Entity type has a key attribute that uniquely identifies each entity in the entity set. But some entity type exists for which key attributes can’t be defined. These are called Weak Entity types . For Example, A company may store the information of dependents (Parents, Children, Spouse) of an Employee. But the dependents can’t exist without the employee. So Dependent will be a Weak Entity Type and Employee will be Identifying Entity type for Dependent, which means it is Strong Entity Type .

## Foreign Key and Primary Key

- Primary Key: A primary key is a field in a table that uniquely identifies each record in the table. It must contain a unique value for each row of data.
- Foreign Key: A foreign key is a column or group of columns in a relational database table that provides a link between data in two tables. It acts as a cross-reference between tables because it references the primary key of another table.

![Primary and Foreign Key](https://d8it4huxumps7.cloudfront.net/uploads/images/65044910e1123_primary_key_versus_foreign_key_02.jpg?d=2000x2000)

## Recursive Entity and Relationship

- Recursive Entity: A recursive entity is an entity that is related to itself. It is used to represent a relationship between instances of the same entity set.
- Recursive Relationship: A recursive relationship is a relationship between instances of the same entity set. It is used to represent relationships between instances of the same entity set.