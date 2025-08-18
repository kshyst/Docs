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

