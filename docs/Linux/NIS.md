# NIS (Network Information Service)

It’s a system used to share configuration databases (like users and groups) across multiple Linux/Unix machines from a central server.


So instead of having separate `/etc/passwd` files on every machine, clients can retrieve user information from an NIS server.

`nsswitch.conf` example:

```text
passwd: files nis
group:  files nis
hosts:  files dns nis
```

This means:

1. Check local files
2. If not found query the NIS server

## Commands

```shell
sudo apt install yp-tools
# show which NIS server is used
ypwhich
# display NIS maps
ypcat passwd
# NIS server daemon
ypserv
# Client service that connect to a NIS server
ypbind
```

NIS is considered outdated and insecure because of weak authentication and broadcasting data on network.

Replacements are `LDAP` for linux and `Active Directory` Windows.