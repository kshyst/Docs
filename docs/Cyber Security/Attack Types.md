# Attack Types

## Spoofing

> In the context of information security, and especially network security, a spoofing attack is a situation in which a person or program successfully identifies as another by falsifying data, to gain an illegitimate advantage.

## CSRF Attack

> CSRF (Cross-Site Request Forgery) is an attack that impersonates a trusted user and sends a website unwanted commands.
> This can be done, for example, by including malicious parameters in a URL behind a link that purports to go somewhere else:

```html
<img src="https://www.example.com/index.php?action=delete&id=123" />
```

> For users who have modification permissions on https://www.example.com, the \<img> element executes action on https://www.example.com without their noticing, even if the element is not at https://www.example.com.

## SQL Injection

SQL Injection (SQLi) is a security vulnerability that occurs when an attacker is able to manipulate a web application’s database queries by inserting malicious SQL code into user input fields. These injected queries can manipulate the underlying database to retrieve, modify, or delete sensitive data. In some cases, attackers can even escalate privileges, gaining full control over the database or server.

### Types of SQL Injection

1. **In-band SQLi (Classic SQLi):** In-band SQLi is the most common type of SQL injection. It occurs when an attacker is able to use the same communication channel to both launch the attack and gather results. This type of SQL injection is typically easy to exploit and can be used to retrieve data from the database.
2. **Inferential SQLi (Blind SQLi):** Inferential SQLi, also known as Blind SQLi, is a type of SQL Injection vulnerability that does not return the results of the injected SQL query in the application’s responses. This makes the exploitation process more challenging but not impossible.
3. **Out-of-band SQLi:** Out-of-band SQLi is an advanced SQL injection technique that requires the attacker to use the database server’s out-of-band channels to launch the attack. This type of SQL injection is less common than in-band SQLi and inferential SQLi but can be just as dangerous.
4. **Error-based SQLi:** Error-based SQLi is an in-band SQL injection technique that relies on error messages thrown by the database server to obtain information about the structure of the database. This information can be used to further refine the attack.
5. **Union-based SQLi:** Union-based SQLi is an in-band SQL injection technique that uses the UNION SQL operator to combine the results of two or more SELECT statements into a single result. This technique can be used to retrieve data from other database tables.

### SQL Injection Prevention

1. Use Prepared Statements and Parameterized Queries : Prepared statements and parameterized queries ensure that user inputs are treated as data rather than part of the SQL query. This approach eliminates the risk of SQL injection.
2. Employ Stored Procedures : Stored procedures are predefined SQL queries stored in the database. These procedures can help prevent SQL injection because they don’t dynamically construct SQL queries.

```sql
CREATE PROCEDURE GetUserByUsername (IN username VARCHAR(50))
BEGIN
   SELECT * FROM users WHERE username = username;
END;
```

3. Whitelist Input Validation : Ensure that user inputs are validated before being used in SQL queries. Only allow certain characters and patterns, such as alphanumeric input, for fields like usernames or email addresses.

4. Use an ORM : Object-Relational Mapping (ORM) frameworks like Hibernate or Entity Framework can help prevent SQL injection by automatically handling query generation, preventing dynamic query construction.


