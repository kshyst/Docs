# Domain Name System(DNS)

DNS operates at application layer. DNS traffic uses UDP port `53` by default and TCP port `53` as default fallback.

## DNS Configs

Main config is under `/etc/resolv.conf`. But it is linked to something else:

```text
ls -l /etc/resolv.conf
# resolv.conf -> /run/systemd/resolve/resolv.conf
```

Some systems its linked to systemd and some other to NetworkManager.
This file is auto configured by the network we are connected to. for example mobile phones
route dns queries through its hotspot gateway and we dont see the actual
DNS server. Also VPNs put their custom DNS servers.

To see the upstream dns that the system actually uses:

```shell
resolvectl status
```

To permanently set DNS so that changing networks wont affect it:

```shell
sudo vim /etc/systemd/resolved.conf
```

and set something like:

```text
[Resolve]
DNS=1.1.1.1 8.8.8.8
FallbackDNS=9.9.9.9
```

By default `Networkctl` changes your DNS server on certain Network Interfaces.

To change it:

```shell
nmcli connection modify "Nigga Killa" ipv4.ignore-auto-dns yes
nmcli connection modify "Nigga Killa" ipv6.ignore-auto-dns yes
sudo systemctl restart NetworkManager
sudo systemctl restart systemd-resolved
resolvectl status # Check the 'Current DNS Server' Section of you interface
```

Now it will use systemd configs as the first priority

To make the systemd dns not use port 53 for listening dns request, inside `/etc/systemd/resolve.conf`:

```text
[Resolve]
DNSStubListener=no
```

## DNS Records

- `A Record`: The A (Address) record maps a hostname to one or more IPv4 addresses.
- `AAAA Record`: A but for IPv6
- `CNAME Record`: Canonical Name record maps a domain name to another domain name.
- `MX Record`: Mail Exchange Record specifies the mail server responsible for handling emails for domain.

For finding IP addresses use `nslookup`.

```shell
nslookup www.example.com
```

```shell
Non-authoritative answer:
Name:   www.example.com
Address: 93.184.215.14
Name:   www.example.com
Address: 2606:2800:21f:cb07:6820:80da:af6b:8b2c
```

The query above led to four packets. In the terminal below, we can see that the first and third packets send DNS queries for the A and AAAA records, respectively. The second and fourth packets show the DNS query responses.

```shell
tshark -r dns-query.pcapng -Nn
```

```shell
    1 0.000000000 192.168.66.89 → 192.168.66.1 DNS 86 Standard query 0x2e0f A www.example.com OPT
    2 0.059049584 192.168.66.1 → 192.168.66.89 DNS 102 Standard query response 0x2e0f A www.example.com A 93.184.215.14 OPT
    3 0.059721705 192.168.66.89 → 192.168.66.1 DNS 86 Standard query 0x96e1 AAAA www.example.com OPT
    4 0.101568276 192.168.66.1 → 192.168.66.89 DNS 114 Standard query response 0x96e1 AAAA www.example.com AAAA 2606:2800:21f:cb07:6820:80da:af6b:8b2c OPT
```

## Resource Records (RRs)

### QUERY (Question Section)

This section contains the question the resolver is asking.

Each entry includes:

- QNAME – domain name being requested
- QTYPE – record type (A, AAAA, MX, NS, TXT, etc.)
- QCLASS – usually IN (internet)

Example:

```text
example.com  IN  A
```

### ANSWER

This section contains the actual answer to the question if the server knows it.

```text
example.com   300   IN   A   93.184.216.34
```

This result shows in order:

- name
- TTL
- class
- record type
- value

### AUTHORITY

This section tells you which DNS servers are authoritative for the domain.

These records are typically NS records.

```text
example.com   IN   NS   ns1.example-dns.com
example.com   IN   NS   ns2.example-dns.com
```

### ADDITIONAL

This section contains extra helpful records to avoid more DNS lookups.

Commonly includes:

- A/AAAA records for nameservers
- EDNS options
- DNSSEC records

Example:

```text
ns1.example-dns.com   IN   A   192.0.2.1
ns2.example-dns.com   IN   A   192.0.2.2
```

##  EDNS (Extension Mechanisms for DNS)

When the original DNS protocol was designed in the 1980s, it had strict limitations. The most significant limitation was that DNS messages sent over UDP could not exceed
`512
bytes`. If a response was larger than
`512
bytes`, the server had to truncate it and tell the client to reconnect using TCP, which is slower and consumes more resources.

As the internet grew, DNS needed to carry more data. Features like IPv6 (which has much longer IP addresses) and DNSSEC (which attaches large cryptographic signatures to DNS records) meant that the
`512
-byte` limit was no longer sufficient.

EDNS (specifically EDNS0) was introduced to solve this. 
It is a set of extensions that allows DNS clients and servers to:

- Negotiate larger UDP packet sizes (often up to `4096 bytes`).
- Support expanded error codes (RCODEs) beyond the original limit of 16.
- Attach additional metadata to DNS queries and responses.

```shell
dig snapp.ir +dnssec
```

Answer: 
```text
; <<>> DiG 9.18.39-0ubuntu0.24.04.3-Ubuntu <<>> snapp.ir +dnssec
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 60283
;; flags: qr rd ra; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags: do; udp: 1280
;; QUESTION SECTION:
;snapp.ir.                      IN      A

;; ANSWER SECTION:
snapp.ir.               77      IN      A       185.143.234.122
snapp.ir.               77      IN      A       185.143.233.122

;; Query time: 48 msec
;; SERVER: 192.168.109.122#53(192.168.109.122) (UDP)
;; WHEN: Sat Apr 04 13:12:18 +0330 2026
;; MSG SIZE  rcvd: 69
```

- **EDNS**: version: 0: Both the client and server are communicating using EDNS version 0.
- **flags**: **do**: This stands for “DNSSEC OK”. It tells the server, “I understand EDNS, and please send me the cryptographic signatures (RRSIG records) along with the IP address.”
- **udp**: 1280: The client is telling the server, “I can accept a UDP response packet up to 1280
bytes in size.” (This is much larger than the old
512
-byte limit, allowing the large RRSIG security record to fit in the response).

## OPT Record

To implement EDNS without breaking older, legacy DNS servers, engineers needed a way to attach this new extension data to DNS messages. They did this by creating the OPT record.

The OPT record is a “pseudo-record.” Unlike standard DNS records (like A, CNAME, or MX records), an OPT record does not actually exist in a domain’s zone file on the server. Instead, it is generated on the fly and added to the “Additional Data” section of a DNS query or response packet.

When a client sends a DNS query with an OPT record, it is essentially telling the DNS server: “*I support EDNS. Here are my parameters.*”

- **UDP Payload Size**: The maximum size of a UDP packet the sender can accept (e.g., `4096 bytes`).
- **Extended RCODE**: Allows for more specific error messages.
- **Version**: The EDNS version being used (currently version 0).
- **Flags**: For example, the DO (DNSSEC OK) bit, which a client sets to tell the server it wants to receive DNSSEC cryptographic records.
- **Variable Options**: This allows for further extensions. A very common one is EDNS Client Subnet (ECS), which includes a portion of the user’s IP address in the query so that Content Delivery Networks (CDNs) can route the user to the closest geographical server.

```text
;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
```

## Resolve DNS

```shell
resolvectl query snapp.ir

dig snapp.ir

nslookup snapp.ir
```