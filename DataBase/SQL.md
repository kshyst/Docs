# SQL

## What is SQL?

SQL is a standard query language used to access and manipulate data in databases. SQL stands for Structured Query Language. We can create, update, delete, and retrieve data in databases like MySQL, Oracle, PostgreSQL, etc. Overall, SQL is a query language that communicates with databases.

## What is a query in SQL?

A query in SQL (Structured Query Language) is a request made to a database to retrieve or manipulate data. It consists of a set of instructions to perform tasks such as fetching, updating, deleting, or inserting data into the database.

## Constraints in SQL

Constraints act as rules or conditions imposed on the data, dictating what values are permissible and what actions can be taken. They play a crucial role in maintaining the quality and coherence of the database by preventing errors. So, explore this section to get a hand on SQL Data Constraints.

- NOT NULL Constraints
- Primary Key Constraints
- Foreign Key Constraints
- Composite Key
- Unique Constraints
- Alternate Key
- CHECK Constraints
- DEFAULT Constraints

### Example of Constraints

```sql
CREATE TABLE table_Name
(
column1 data_type(size) NOT NULL,
column2 data_type(size) NOT NULL,
);

```

PRIMARY KEY in SQL is a column (or group of columns) that uniquely identifies the records in that table. A primary key must contain unique values and can not have any NULL value.

There can only be one primary key in a table, but that primary key can consist of one or more columns. When there are two or more columns in the primary key it is called a composite key.

A primary key automatically has a UNIQUE constraint defined on it, and it ensures that there are no duplicate or NULL values in that column.

```sql
CREATE TABLE table_name (
 column1 datatype constraint,
 column2 datatype constraint,
 CONSTRAINT pk_constraint_name PRIMARY KEY (column1, column2, …)
);
```

## JOINS And Normalization

### INNER JOIN

The INNER JOIN is a process that matches rows from the first table and the second table which have the same key (as defined by the ON constraint) to create a result row with the combined columns from both tables. After the tables are joined, the other clauses we learned previously are then applied.

```sql
SELECT column, another_table_column, …
FROM mytable
INNER JOIN another_table 
    ON mytable.id = another_table.id
WHERE condition(s)
ORDER BY column, … ASC/DESC
LIMIT num_limit OFFSET num_offset;
```

### OUTER JOIN

If the two tables have asymmetric data, which can easily happen when data is entered in different stages, then we would have to use a LEFT JOIN, RIGHT JOIN or FULL JOIN instead to ensure that the data you need is not left out of the results.

```sql
SELECT column, another_column, …
FROM mytable
INNER/LEFT/RIGHT/FULL JOIN another_table 
    ON mytable.id = another_table.matching_id
WHERE condition(s)
ORDER BY column, … ASC/DESC
LIMIT num_limit OFFSET num_offset;
```

## Expressions

Use expressions to write more complex logic on column values in a query. These expressions can use mathematical and string functions along with basic arithmetic to transform values when the query is executed.

```sql
SELECT particle_speed / 2.0 AS half_particle_speed
FROM physics_data
WHERE ABS(particle_position) * 10.0 > 500;
```


