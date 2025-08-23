# Hash for Password Storage

## Same Password Problem

What if two users have the same password? As a hash function will always turn the same input into the same output, you will store the same password hash for each user. That means if someone cracks that hash, they gain access to more than one account. It also means someone can create a `Rainbow Table` to break the hashes.

## Rainbow Table

A Rainbow Table is a lookup table of hashes to plaintexts, so you can quickly find out what password a user had just from the hash. A rainbow table trades the time to crack a hash for hard disk space, but it takes time to create.

Hash	                       Password
- 02c75fb22c75b23dc963c7eb91a062cc	: zxcvbnm
- b0baee9d279d34fa1dfd71aadb908c3f	: 11111
- c44a471bd78cc6c2fea32b9fe028d30a	: asdfghjkl
- d0199f51d2728db6011945145a1b607a	: basketball
- dcddb75469b4b4875094e14561e573d8	: 000000
- e10adc3949ba59abbe56e057f20f883e	: 123456
- e19d5cd5af0378da05f63f891c7467af	: abcd1234
- e99a18c428cb38d5f260853678922e03	: abc123
- fcea920f7412b5da7be0cf42b8c93759	: 1234567
- 4c5923b6a6fac7b7355f53bfe2b8f8c1	: inS3CyourP4$$

### Protect against rainbow tables

To protect against rainbow tables, we add a salt to the passwords. The salt is a randomly generated value stored in the database and should be unique to each user. In theory, you could use the same salt for all users, but duplicate passwords would still have the same hash and a rainbow table could still be created for passwords with that salt.

The salt is added to either the start or the end of the password before it’s hashed, and this means that every user will have a different password hash even if they have the same password. Hash functions like Bcrypt and Scrypt handle this automatically. Salts don’t need to be kept private.

- We select a secure hashing function, such as Argon2, Scrypt, Bcrypt, or PBKDF2.
- We add a unique salt to the password, such as Y4UV*^(=go_!
- Concatenate the password with the unique salt. For example, if the password is AL4RMc10k, the result string would be AL4RMc10kY4UV*^(=go_!
- Calculate the hash value of the combined password and salt. In this example, using the chosen algorithm, you need to calculate the hash value of AL4RMc10kY4UV*^(=go_!.
- Store the hash value and the unique salt used (Y4UV*^(=go_!).- 

## Linux Passwords

In linux password hashes are stored in `/etc/shadow` and is only readable by root.

The `shadow` file contains the password information. Each line contains `nine fields`, separated by colons (:). The first two fields are the login name and the encrypted password. More information about the other fields can be found by executing `man 5 shadow` on a Linux system.

The encrypted password field contains the hashed passphrase with four components: prefix (algorithm id), options (parameters), salt, and hash. It is saved in the format `$prefix$options$salt$hash`. The prefix makes it easy to recognise Unix and Linux-style passwords; it specifies the hashing algorithm used to generate the hash.

Use `man 5 crypt` to get info about prefixes.

Example for linux shadow:

```shell
    sudo cat /etc/shadow | grep strategos
    strategos:$y$j9T$76UzfgEM5PnymhQ7TlJey1$/OOSg64dhfF.TigVPdzqiFang6uZA4QA1pzzegKdVm4:19965:0:99999:7:::
```

- `y` indicates the hash algorithm used, yescrypt
- `j9T` is a parameter passed to the algorithm
- `76UzfgEM5PnymhQ7TlJey1` is the salt used
- `/OOSg64dhfF.TigVPdzqiFang6uZA4QA1pzzegKdVm4` is the hash value

## MS Windows Passwords

MS Windows passwords are hashed using `NTLM`, a variant of `MD4`. They’re visually identical to MD4 and MD5 hashes, so it’s very important to use context to determine the hash type.

On MS Windows, password hashes are stored in the `SAM` (Security Accounts Manager). MS Windows tries to prevent normal users from dumping them, but tools like mimikatz exist to circumvent MS Windows security. Notably, the hashes found there are split into NT hashes and LM hashes.