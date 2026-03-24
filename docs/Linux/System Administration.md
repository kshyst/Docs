# System Administration

### Changing Password

```shell
passwd
```

> root user can change other users password without providing their current password

```shell
sudo passwd kshyst
```

## Users and Groups


### Add User
- `-d`: Sets home directory
- `-m`: Creates home directory
- `-s`: Specifies shell
- `-G`: Adds to additional groups
- `-c`: Comment

> Use `/etc/skel` for creating templates 

```shell
sudo useradd <name>
sudo adduser
```

### Modify User

- `-L`: Lock this account
- `-U`: Unlock this account
- `-aG`: Adds to more groups
- `-G`: Deletes previous groups for users and sets new ones.

```shell
sudo usermod
sudo usermod -aG sudo newuser
```

### Deleting Users

```shell
sudo userdel newuser
```

Using `-r` to delete everything related to it including home directory and mail spool.

### Group

- `-g`: Sets group id

```shell
groupadd -g 2000 meowers
groupmod -g 2000 meowers
groupdel meowers
```

> In most systems, the human made groups start from id 1000 and goes on

## Important Files

### `/etc/passwd`

Contains all user information on the system

User password used to be in here but moved to `/etc/shadow`. Thats why we see `x` for everyone on the second slot.

### `/etc/shadow`

There are 4 numbers after users hashed passwords:

1. When was the last time password was changed
2. User won't be able to change password n days after a change
3. After n many days user have to change their password
4. n days before password change deadline, user will be informed

> The day count is how many days since 1970/01/01

For setting password age use `chage`:

- `-E`: Set date in YYYY-MM-DD format

```shell
chage kshyst
```

### `/etc/group`

> Groups can have passwords too thats why there is x in second slot

### `/etc/gshodow`

Contains group passwords

---

To get data from database about users use:

```shell
getent passwd kshyst
getent shadow kshyst
```

This chooses database and user
