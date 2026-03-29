# IP Commands

### ipcalc

To calculate network and broadcast address use `ipcalc`.

- ipcalc 192.168.0.1/24                         
- ipcalc 192.168.0.1/255.255.128.0              
- ipcalc 192.168.0.1 255.255.128.0 255.255.192.0
- ipcalc 192.168.0.1 0.0.63.255

Example output:

```shell
Address:   192.168.0.1          11000000.10101000.00000000. 00000001                │
Netmask:   255.255.255.0 = 24   11111111.11111111.11111111. 00000000                │
Wildcard:  0.0.0.255            00000000.00000000.00000000. 11111111                │
=>                                                                                  │
Network:   192.168.0.0/24       11000000.10101000.00000000. 00000000                │
HostMin:   192.168.0.1          11000000.10101000.00000000. 00000001                │
HostMax:   192.168.0.254        11000000.10101000.00000000. 11111110                │
Broadcast: 192.168.0.255        11000000.10101000.00000000. 11111111                │
Hosts/Net: 254                   Class C, Private Internet
```

### ifconfig

InterFace Configurations

```shell
ifconfig
# Change a nic's address
sudo ifconfig enp0s25 192.168.42.42
# Change Netmask
ifconfig eth0 netmask 255.255.0.0
# Turn off an interface
sudo ifconfig enp0s25 down
# To check upness of a nic
ifup ...
ifdown ...
```


### ip

```shell
ip addr add 172.19.1.10/24 dev eth2 # temporary adding an IP
ip addr show eth2
ip addr del 172.19.1.10/24 dev eth2 # deleting an IP address
ip link set eth2 up # brining a NIC up
ip route show
ip route add default via 192.168.1.1 # add default gateway
```

