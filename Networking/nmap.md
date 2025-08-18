# Nmap(Network Mapper)

- IP range using -: If you want to scan all the IP addresses from 192.168.0.1 to 192.168.0.10, you can write 192.168.0.1-10
- IP subnet using /: If you want to scan a subnet, you can express it as 192.168.0.1/24, and this would be equivalent to 192.168.0.0-255
- Hostname: You can also specify your target by hostname, for example, example.thm

## Target ports

Nmap scans 1000 most common ports by default . Other options are:
- `-F` is for Fast mode, which scans the 100 most common ports (instead of the default 1000).
- `-p[range]` allows you to specify a range of ports to scan. For example, -p10-1024 scans from port 10 to port 1024, while -p-25 will scan all the ports between 1 and 25. Note that `-p-` scans all the ports and is equivalent to -p1-65535 and is the best option if you want to be as thorough as possible.

## Timing 

Scanning too fast might trigger the IDS so nmap offers templates for scan timing.

### Timer
Use `-T[0-5]` to set the timer.

paranoid (0), sneaky (1), polite (2), normal (3), aggressive (4), and insane (5)

### Parallelism
Another useful option is the number of parallel probes. `--min-parallelism <numprobes>` and `--max-parallelism <numprobes>`. By default nmap configures it automatically. If network performs poorly the number decrease, and if flawlessly it will increase.

### Rate
A similar helpful option is the `--min-rate <number>` and `--max-rate <number>`. As the names indicate, they can control the minimum and maximum rates at which nmap sends packets. The rate is provided as the number of packets per second. It is worth mentioning that the specified rate applies to the whole scan and not to a single host.

### Max waiting time
This option specifies the maximum time you are willing to wait, and it is suitable for slow hosts or hosts with slow network connections.

`--host-timeout <time>`

## Getting More From Scans

Use `-v` for verbosity and get more info. add more v for more verbosity and also you can use like `-v3`.

Use `-d` for debugging level output. Add more d for more info.

## Output Scans

- `-oN <filename>` - Normal output
- `-oX <filename>` - XML output
- `-oG <filename>` - grep-able output (useful for grep and awk)
- `-oA <basename>` - Output in all major formats

## Options

- `-sn`: For discovery of any type of TCP or UDP
- `-sL`: Only lists the host that are gonna be scanned
- `-sT`: Connect scan. Establishes 3 way handshake on every target TCP port.
- `-sS`: SYN scan. Unlike the Connect scan which establishes 3 way handshake, it only sends SYN packet. Advantages are that it logs less packets and is also stealty.
- `-sV`: Version detection.
- `-sU`: UDP port scan.
- `-PA[port-list]`: TCP ACK discovery
- `-PS[port-list]`: TCP SYN discovery
- `-PU[port-list]`: UDP discovery
- `-Pn`: Treats all hosts as online and do the port scanning on all of them. Even when they appear offline in discovery phase.
- `-O`: OS detection
- `-A`: OS detection, version scanning, and traceroute, among other things.


This screenshot shows -sT which sends ACK too
![-sT](img/nmap1.png)

This screenshot shows -sS which only sends SYN
![-sS](img/nmap2.png)