# Network Configuration

## Network Interface

The NIC (or Network Interface Card) is the physical network hardware in your computer. This can be the chip+antenna in your mobile phone or an Ethernet Card connected to a network cable on your PC.

See `ifconf` in `Network Commands/ip.md`.

Configuration file location : `/etc/network/interfaces`

## Network Manager

This service can "watch" the status of network and various configuration and configure the network cards (specially the wifi ones) accordingly. This is what makes our laptop connected whenever we open it in an area with a known WiFi or ask about the password if we want to connect to a new network or assign IP addresses as soon as we connect the cable to our Ethernet card. This IP assignment might happen via the "permanent IP configuration" on your device or a protocol called DHCP. When using DHCP (Dynamic Host Configuration Protocol), your computer asks a DHCP server (say your home's wifi router) about the IP, Netmask, Default gateway, DNS and other stuff and sets them.

### nmtui

Various frontend GUI (graphical user interface) or TUI (textual user interface. try `nmtui` ) or CLI (command line interfaces) programs exists to control or configure the NetworkManager daemon.

### nmcli

Use `nmcli` command to control network manager.

| Command | Usage |
|--------|------|
| `general` | NetworkManager’s general status and operations |
| `networking` | Overall networking control |
| `radio` | NetworkManager radio switches |
| `connection` | Control network connections |
| `device` | Devices controlled by NetworkManager |
| `agent` | Secret or polkit agent |
| `monitor` | Monitor network changes |

```shell
nmcli device
```

## Hostname

Hostname of the machine is stored in `/etc/hostname`.

```shell
# Shows hostname
hostname
# Shows info about host
hostnamectl
# Changes hostname Or you can change it as transient, which is a temporary change using the --transient switch.
hostnamectl set-hostname mycoolmachine
# Set pretty hostname
hostnamectl --pretty set-hostname "LAN Shared Storage"
```

`/etc/hosts` Contains a list of IPs and their corresponding names, including your own computers.

```shell
head -20 /etc/hosts
```

Here we can like set

```shell
127.0.0.1 digikala.com
```

it would open localhost everytime we type in digikala.com

This command will show host entries including the custom ones we wrote.
```shell
getent hosts
```

## DNS Configuration

Under `/etc/resolve.conf`.

```shell
nameserver 192.168.1.1
nameserver 4.2.2.4
domain jadi.net
search jadi.net company.com
```

Here I'm telling my computer to contact the DNS on my home network (192.168.1.1) or a DNS located at 4.2.2.4 if it needed to translate an address to an IP.

The `domain` configuration sets a local domain name so the machines in this domain will be able to use a short name (tv, instead of jadi.net) and the `search` config does kind the same and tells the resolver to search for jadi.net and company.com if it was trying to resolve 

## nsswitch

Name Server Switch

Under `/etc/nsswitch.conf`

used to configure which services are to be used to determine information such as hostnames, password files, and group files.

Instead of always reading from only local files, Linux can search multiple sources (files, DNS, LDAP, etc.)

```shell
# Begin /etc/nsswitch.conf

passwd: files
group: files
shadow: files

publickey: files

hosts: files dns myhostname
networks: files

protocols: files
services: files
ethers: files
rpc: files

netgroup: files

# End /etc/nsswitch.conf
```

So if someone wants to check a password, the system will try the password _file_ on the system. Or if they want to check an ip address of a hostname, my config says `hosts: files dns myhostname` so the computer first tries the files (`/etc/hosts`) and then goes for DNS. If I reverse these and change the line to

- `passwd: files systemd`

When looking up users, check `/etc/passwd` first, then systemd.

- `hosts: files dns`

When resolving hostnames, check `/etc/hosts` first, then DNS.

## systemd-resolved

systemd provides a DNS called `systemd-resolved`. It listens for DNS requests on `127.0.0.53` and answers back after consulting the `/etc/systemd/resolv.conf` or `/etc/resolv.conf`.