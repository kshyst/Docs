# Constraints

Select query with constraints.

```sql
SELECT column, another_column, …
FROM mytable
WHERE condition
    AND/OR another_condition
    AND/OR …;
```

![SQL Constraints](img/Screenshot%20from%202025-02-16%2015-14-58.png)

![SQL Constraints](img/Screenshot from 2025-02-16 15-17-54.png)

# Filtering  and Sorting

- **DISTINCT**: This keyword will retrieve only unique records.
- **ORDER BY**: This keyword is used to sort the result-set in ascending or descending order.

```sql
SELECT column, another_column, …
FROM mytable
WHERE condition(s)
ORDER BY column ASC/DESC;
```

- **LIMIT**: This keyword will limit the number of records returned in the result set.

```sql
SELECT column, another_column, …
FROM mytable
WHERE condition(s)
ORDER BY column ASC/DESC
LIMIT num_limit OFFSET num_offset;
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

## Aggregation

```sql
SELECT AGG_FUNC(column_or_expression) AS aggregate_description, …
FROM mytable
WHERE constraint_expression;
```

![SQL Aggregation](img/agg.png)

For applying aggregation functions to certain columns in a table, you can use the GROUP BY clause. This clause divides the rows of a table into groups based on the values of one or more columns. Then, you can apply the aggregation functions to each group.

The HAVING clause constraints are written the same way as the WHERE clause constraints, and are applied to the grouped rows.
```sql
SELECT group_by_column, AGG_FUNC(column_expression) AS aggregate_result_alias, …
FROM mytable
WHERE condition
GROUP BY column
HAVING group_condition;
```

## Order Of Execution 

- FROM and JOIN
- WHERE
- GROUP BY
- HAVING
- SELECT
- DISTINCT
- ORDER BY
- LIMIT / OFFSET

## Inserting new data

```sql
INSERT INTO mytable
VALUES (value_or_expr, another_value_or_expr, …),
       (value_or_expr_2, another_value_or_expr_2, …),
       …;
```

```sql
INSERT INTO mytable
(column, another_column, …)
VALUES (value_or_expr, another_value_or_expr, …),
      (value_or_expr_2, another_value_or_expr_2, …),
      …;
```

## Updating Rows

```sql
UPDATE mytable
SET column = value_or_expr, 
    other_column = another_value_or_expr, 
    …
WHERE condition;
```

When Updating rows the best practice is to use SELECT to first check if the WHERE is correct the apply the query to the database.

## Deleting Rows

```sql
DELETE FROM mytable
WHERE condition;
```

## Creating Tables

```sql
CREATE TABLE IF NOT EXISTS mytable (
    column DataType TableConstraint DEFAULT default_value,
    another_column DataType TableConstraint DEFAULT default_value,
    …
);
```

![SQL Table](img/types.png)

### Table Constraints

![SQL Table Constraints](img/c.png)

```sql
CREATE TABLE movies (
    id INTEGER PRIMARY KEY,
    title TEXT,
    director TEXT,
    year INTEGER, 
    length_minutes INTEGER
);
```

## Altering Tables

### Adding Columns

```sql
ALTER TABLE mytable
ADD column DataType OptionalTableConstraint 
    DEFAULT default_value;
```

### Removing Columns

```sql
ALTER TABLE mytable
DROP column_to_be_deleted;
```

### Renaming Table

```sql
ALTER TABLE mytable
RENAME TO new_table_name;
```

### Renaming Columns

```sql
ALTER TABLE mytable
RENAME COLUMN column_to_be_changed TO new_column_name;
```

## Dropping Tables

```sql
DROP TABLE IF EXISTS mytable;
```

## Applying foreign key when creating a column

```sql
CREATE TABLE orders (
    order_id INTEGER PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(customer_id)
);
```

### In case you want to add a foreign key to an existing table

```sql
ALTER TABLE orders
ADD FOREIGN KEY (customer_id) REFERENCES customers(customer_id);
```

### Removing a foreign key

```sql
ALTER TABLE orders
DROP CONSTRAINT orders_customer_id_fkey;
```

## Having

`HAVING` is used because `WHERE` can't be used with aggregated functions.

```sql
SELECT column_name(s)
FROM table_name
WHERE condition
GROUP BY column_name(s)
HAVING condition
ORDER BY column_name(s);
```

## UNION

Used to combine results of multiple `SELECT` queries.

*It automatically removes duplicate rows*

```sql
SELECT column_name(s) FROM table1
UNION
SELECT column_name(s) FROM table2;
```

Requirements for `UNION` to work:

- Every `SELECT` statement within `UNION` must have the same number of columns
- The columns must also have similar data types
- The columns in every `SELECT` statement must also be in the **same order**