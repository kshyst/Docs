# Netcat (nc)

The `nc` (or netcat) utility is used for just about anything under the sun involving TCP, UDP, or UNIX-domain sockets. It can open TCP connections, send UDP packets, listen on arbitrary TCP and UDP ports, do port scanning, and deal with both IPv4 and IPv6. Unlike telnet, nc can be used easily in the scripts and separates error messages onto standard error instead of sending them to standard output, as telnet does with some. 

```shell
# Open a listener
nc -l 1337
# Opening connection to that port and send data
nc localhost 1337
```