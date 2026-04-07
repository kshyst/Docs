# Users and Permissions

## Users

Check users in system
```shell
# You have a line for each logged in users (every single shell window is a separated login).
w
# More detail
who
# Logged out
last
```

> there is a way to check the failed logins too: `last -f /var/log/btmp`


Switch user

```shell
su kshyst
# Switching to root using the sudoer user password instead of root password
sudo su -
```

See groups and id and yourself

```shell
groups

id # gid for all groups

whoami
```

See who is in each group:
```shell
cat /etc/group
```

Assign groups to a user:
```shell
# This adds kshyst to sudoers
sudo usermod -aG sudo kshyst
```

Or edit `/etc/sudoers`

How root gets the right to run all the commands: `txt root ALL=(ALL:ALL) ALL`

The `ALL:ALL` means these users can run as any user and any group. The last ALL tells the sudo that these users/groups can run ALL commands.

To change this file use `visudo` instead of `vi`

Add this line to `visudo` to make user a passwordless sudoer:

```text
kshyst ALL=(ALL) NOPASSWD: ALL
```

Changing group for something:
```shell
sudo chgrp postgres newfile
```

Users can be in many groups. Each user has a default group that is used when creating files and dirs by that user. To change default group:
```shell
newgrp MyGroup
```

Create User:

```shell
sudo adduser ali
```

Also we can use `useradd` but it doesnt make user dir.

Give sudo access:

```shell
sudo usermod -aG sudo ali
```

### user limits

The resources on a Linux machine can be managed for users by the `ulimit` command. It is part of the PAM system.

```shell
# See all limit options
ulimit -a
# This will limit the CPU TIME of any process to 1 second. 
ulimit -t 1
```

## Permissions

If `d` at the beggining means its directory. If `-` means file.

Permissions are : User - Group - Other in order

Always 10 slots. First is type, other 3 for each is Read, Write, Execute

| Octal | Symbolic | Meaning |
|------|----------|--------|
| 0 | --- | No permissions |
| 1 | --x | Execute |
| 2 | -w- | Write |
| 3 | -wx | Write + Execute |
| 4 | r-- | Read |
| 5 | r-x | Read + Execute |
| 6 | rw- | Read + Write |
| 7 | rwx | Read + Write + Execute |


For example rw is r (4) + w(2) = 6

### chmod

Main command for changing permissions is `chmod`.

```shell
chmod 777 filename.sh

chmod u+rwx filename.sh

chmod o-rwx filename.sh

chmod u+rwx,o-rw,g+x filename.sh

```

- `-R`: Recursively gives all that permission


### chown

Change owner

```shell
chown kshyst:kshyst newfile.sh
```


## Access Modes

When we want to change password we use `passwd`. It writes the passwords in `/etc/shadow`.
When we run `cat/etc/shadow` it gives permission denied. When we  `ls -l /etc/shadow` we see that only root has write access and the `shadow` group has read access.
This is weird because we can run passwd without sudo.
When we `ls -l /usr/bin/passwd` we see that it has this permissions: `-rwsr-xr-x`. Which has `s` instead of `x` for the user. It is call `SUID` which means set user id. 

It is done like this.
```shell
sudo chmod u+s /usr/bin/passwd
```

When something has SUID, it will be run by the access of its owner, not the user running the command. The passwd is owned by root so when we run it, it will use the acces level of root.

| Octal | Symbolic | Name |
|------|----------|------|
| 4000 | u+s | Set UID (SUID) |
| 2000 | g+s | Set GID (SGID) |
| 1000 | +t | Sticky Bit |

> Sticky means that no one can delete that file except the owner of that file.
> The `t` will sit where `x` does.


### umask

First run it:
```shell
umask
```

It will output something like `002`

If we subract this from 666:

666 - 002 = 664

This is the default permissions that new files will get.

We can change umask like this:

```shell
umask 066
```

Which will result in only rw by user files.
