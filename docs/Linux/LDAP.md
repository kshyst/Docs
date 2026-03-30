# LDAP (Lightweight Directory Access Protocol)

It’s a protocol used to store and query directory information over a network. 

It acts as a central database for:

- user accounts
- passwords
- groups
- computers
- email addresses
- organizational info


Example: 

In a company with 500 employees, instead of creating users on every Linux server separately, all user accounts are stored in one LDAP server. Every machine asks that server when someone logs in.

- User logs into a Linux server.
- The system checks `/etc/nsswitch.conf`.
- If LDAP is configured, it queries the LDAP server.
- The LDAP server returns the user information.

```text
dc=company,dc=com
 ├── ou=people
 │    ├── uid=alice
 │    ├── uid=bob
 │
 └── ou=groups
      ├── cn=admins
      └── cn=developers
```

- dc = domain component
- ou = organizational unit
- uid = user ID
- cn = common name
