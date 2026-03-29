# ss Command(Socket Statistics)

`netstat` is the older version

```shell
ss
```
### Data:
- **Netid**: Type of socket, common types are tcp, udp, u_str(Unix Stream), u_seq(Unix Sequence)
- **State**: State of socket which are ESTAB, UNCONN, LISTEN
- **Recv-Q**: Number of received packets in queue
- **Send-Q**: Number of sent packets in queue


## Netstat vs ss

The ss command is considered a replacement command for the obsolete netstat. The speed and better filtering options of CLI utilities from the iproute2 software package are preferable to the net-tools software package.

- `-n`: Numeric
- `-a`: All ports
- `-r`: Routes

```shell
# See which ports are in listening mode
ss -na | grep LISTEN | grep tcp
# Best combination
ss -tulpn
```