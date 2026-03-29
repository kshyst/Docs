# IPv4

32 bit number made with 4 octets(numbers between dots)

## Network and Broadcast Addresses

At the risk of oversimplifying things, the 0 and 255 are reserved for the network and broadcast addresses, respectively. In other words, 192.168.1.0 is the network address, while 192.168.1.255 is the broadcast address. Sending to the broadcast address targets all the hosts on the network.

The network address is a Logical AND between the IP address and the networkmask. The broadcast address is the network address with a Logical OR when all the host bits are changed to 1.

# Private IP Addresses

RFC 1918 defines the following three ranges of private IP addresses:

- 10.0.0.0 - 10.255.255.255 (10/8)
- 172.16.0.0 - 172.31.255.255 (172.16/12)
- 192.168.0.0 - 192.168.255.255 (192.168/16)

For a private IP address to access the Internet, the router must have a public IP address and must support Network Address Translation (NAT). 

# IP Ranges

| Class | First Octet | Range |
|------|-------------|-----------------------------|
| A | 1–126 | 1.0.0.0 to 126.255.255.255 |
| B | 128–191 | 128.0.0.0 to 191.255.255.255 |
| C | 192–223 | 192.0.0.0 to 223.255.255.255 |


# Additional Data

Port number max is 65535 which is 2^16 -1