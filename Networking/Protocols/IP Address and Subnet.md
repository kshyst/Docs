# IPv4

32 bit number made with 4 octets(numbers between dots)

## Network and Broadcast Addresses

At the risk of oversimplifying things, the 0 and 255 are reserved for the network and broadcast addresses, respectively. In other words, 192.168.1.0 is the network address, while 192.168.1.255 is the broadcast address. Sending to the broadcast address targets all the hosts on the network.

# Private IP Addresses

RFC 1918 defines the following three ranges of private IP addresses:

- 10.0.0.0 - 10.255.255.255 (10/8)
- 172.16.0.0 - 172.31.255.255 (172.16/12)
- 192.168.0.0 - 192.168.255.255 (192.168/16)

For a private IP address to access the Internet, the router must have a public IP address and must support Network Address Translation (NAT). 

# Additional Data

Port number max is 65535 which is 2^16 -1