# iptables

The iptables utility is a software firewall for Linux distributions that lets you control how network traffic is handled by the Linux kernel. With iptables, you can define rules that match traffic by properties like protocol, port, source or destination address, and network interface, and then decide whether to allow it, block it, or log it. These rules are organized into tables and chains (such as the INPUT, OUTPUT, and FORWARD chains) and are evaluated from top to bottom, which makes rule order an important part of getting the behavior you expect.


## How to use it
Keep in mind that rule order matters. All of the iptables commands in this guide use the -A option to append a new rule to the end of a chain. If you want to put a rule somewhere else in the chain, you can use the -I option, which allows you to specify the position of the new rule (or place it at the beginning of the chain by omitting a rule number).

By default `iptables` are ephemeral and removed between boots.

### Check current rule set

```shell
sudo iptables -S
sudo iptables -L
```
