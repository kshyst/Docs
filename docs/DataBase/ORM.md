# ORM

## What is ORM?

Object-relational mapping (ORM) is a key concept in the field of Database Management Systems (DBMS), addressing the bridge between the object-oriented programming approach and relational databases. ORM is critical in data interaction simplification, code optimization, and smooth blending of applications and databases.

## What is ORDBMS?

An ORDBMS stands for an object-relational database management system that further enhances the functionalities of a relational database by incorporating object-oriented principles. It deals with complex data types, encapsulation, inheritance, and other concepts in an object-oriented way, which is the right support for applications that require both relational and object-oriented abilities.

### Entities

In the realm of ORM, entities are synonymous with the objects or classes in object-oriented programming that are bound to tables in the relational databases. They serve as abstractions of business objects or the concepts within the application and their definition is in the code. The ORM component carries out the transformation of these entities into database tables and thus provides smooth communication between the application and the database that lies underneath that application.

### Relationships

Relationships in ORM are the associations between entities that are defined in the code. These relationships are translated into the relational database tables by the ORM component. The relationships can be one-to-one, one-to-many, or many-to-many, and they are crucial in defining the structure of the database and the application.

### Persistence

Persistence refers to the capability to keep data after an application is ended. The use of Object Relational Mapping (ORM) causes data to persist even when the application is closed or restarted because it is stored in a relational database that makes it secure and available even when the application is off. This is the most important function of the ORM because it allows the temporal persistence of the data used by the application.

## ORM in DBMS

With Object-Relational Mapping, it becomes much easier to work with an object-oriented programming language and relational database. Fundamentally, it acts as a translator, translating data between the database and the application without any hitch. ORM enables developers to work with objects in their programming language which are mapped to corresponding database entities, such as tables, views, or stored procedures.

### Entity Mapping

The first initialization step in the Object-Relational Mapping (ORM) process is to identify entities from the object-oriented model and match them with corresponding tables in the relational database. Developers specify objects or classes for business entities in their code and the ORM framework manages the conversion of such entities into tables in databases. Each feature in the class maps to a new column in the table, and instances of the class turn into rows in the table.

### Relationships Mapping

Having entities mapped the next important stage is to provide relations between them and structuring the relational database schema accordingly. ORM frameworks provide functions that enable us to represent the relationship between entities like one-to-one, one-to-many, and many-to-many. Such relationships are transformed into foreign keys, which allow the data to be kept both consistent and valid even when stored in tables that could be related.

### Data Type Mapping

Data mapping is the practice of an object-oriented model mapping the data type to the database. ORM frameworks handle the conversion of data types, that enable the objects of the application to align with the data types . It becomes the most important step, ensuring cohesion and preventing data type mismatches that will result in errors.

### CRUD Operations

Crud (Create, Read, Update, Delete) operations constitute the basis of any database interaction. The use of ORM libraries simplifies the process of the implementation of these operations by implementing the high-level abstractions. Developers can make changes to entities in their program and then the ORM allows for translation of these operations into corresponding SQL queries. That process feeds the need to write complex SQL conditions by developers automatically, creating the interaction with the database more user-friendly and with less errors.

### Query Language

ORM framework is usually a query language designed with object-oriented features for interacting with the database. For example, Hibernate Package which is a Java-based ORM framework, makes use of Hibernate Query Language (HQL). HQL allows developers to express database queries by using object-oriented syntax enabling themselves to get data with ease, manipulating it without directly dealing with SQL.