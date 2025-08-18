# tcpdump

## Flags

For choosing network interface use `-i` flag.

For getting the list of interfaces use `ip a s`

For capturing packets use `-w` flag and provide file name with `pcapg` extension.

For reading the capture file use `-r` flag.

For specifying them number of packets you want to capture use `-c` flag.

By default, tcpdump resolves IPs. To turn this feature off use `-n` flag. Also you can use `-nn` to stop both port number and dns lookups.

To produce verbose output you can use `-v` and for more `-vv` and ...


## Filtering Expressions

Filtering by host:

```shell
sudo tcpdump host example.com -w http.pcap
```

Also you can define `src` or `dst`

```shell
sudo tcpdump src host example.com -w http.pcap
```

Filtering by port:

```shell
tcpdump port PORT_NUMBER
```

You can also use `src` and `dst` on this.

Filtering by Protocol: (only ip, ip6, icmp, tcp, udp, arp)

```shell
tcpdump http
```

You can also use logical operations like `and`, `or` and `not`.

## Advanced Packet Filtering

Use `man pcap-filter` to see all filters.

- `greater LENGTH`: Filters packets that have a length greater than or equal to the specified length
- `less LENGTH`: Filters packets that have a length less than or equal to the specified length

### Some Interesting Filters

- `ether[0] & 1 != 0` takes the first byte in the Ethernet header and the decimal number 1 (i.e., 0000 0001 in binary) and applies the & (the And binary operation). It will return true if the result is not equal to the number 0 (i.e., 0000 0000). The purpose of this filter is to show packets sent to a multicast address. A multicast Ethernet address is a particular address that identifies a group of devices intended to receive the same data.
- `ip[0] & 0xf != 5` takes the first byte in the IP header and compares it with the hexadecimal number F (i.e., 0000 1111 in binary). It will return true if the result is not equal to the (decimal) number 5 (i.e., 0000 0101 in binary). The purpose of this filter is to catch all IP packets with options.

## Display Packets

- `-q`: Quick output; print brief packet information
- `-e`: Print the link-level header
- `-A`: Show packet data in ASCII
- `-xx`: Show packet data in hexadecimal format, referred to as hex
- `-X`: Show packet headers and data in hex and ASCII

