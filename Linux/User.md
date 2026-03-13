# User

Create User:

```shell
sudo adduser ali
```

Also we can use `useradd` but it doesnt make user dir.

Give sudo access:

```shell
sudo usermod -aG sudo ali
```

Switch user:

```shell
su - ali
```