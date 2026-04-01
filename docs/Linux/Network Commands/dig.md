# Dig

The `dig` command is a DNS lookup tool. If you are having a problem with a domain name, you can check how it is being resolved to IPs; and by whom.

```shell
dig joemama.com
# Tells dig what dns server to use
dig @8.8.8.8 google.com
# digging a linux hostname
dig predator
# Only get IPs
dig google.com +short
```

# host

```shell
host digikala.com
```