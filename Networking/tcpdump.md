# tcpdump

For choosing network interface use `-i` flag.

For getting the list of interfaces use `ip a s`

For capturing packets use `-w` flag and provide file name with `pcapg` extension.

For reading the capture file use `-r` flag.

For specifying them number of packets you want to capture use `-c` flag.

By default, tcpdump resolves IPs. To turn this feature off use `-n` flag. Also you can use `-nn` to stop both port number and dns lookups.

To produce verbose output you can use `-v` and for more `-vv` and ...
