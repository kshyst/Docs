# NF (Normal Form)  

## 1NF (Eliminating Duplicate Records)

- Each row should have atomic value.
- Each row should be unique.
- Each column should have unique name.
- The order in which data is stored does not matter.

## 2NF (Eliminating Partial Dependency)

Every non prime attribute should be dependant on the whole primary key, not just part of it. (When we have composite primary keys)

## 3NF (Eliminating Transitive Dependency)

Non prime attributes should not depend on other non prime attributes.

## BCNF (Boyce-Codd Normal Form)

In every determinant function dependency, the left side should be a super key. (X -> Y, X should be a super key)


## 4NF (Eliminating Multi-Valued Dependency)

Consider a table where (StudentID, Language, Hobby) are attributes. If a student can have multiple hobbies and languages, a multi-valued dependency exists. To resolve this, split the table into separate tables for Languages and Hobbies.

## 5NF (Eliminating Join Dependency)

A table is in 5NF if it cannot be decomposed into smaller tables without losing information. It ensures that all join dependencies are implied by candidate keys.

Consider a table where (StudentID, Language, Hobby) are attributes. If a student can have multiple hobbies and languages, a multi-valued dependency exists. To resolve this, split the table into separate tables for Languages and Hobbies.

