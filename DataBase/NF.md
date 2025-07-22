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
