# Domain Name System(DNS)

DNS operates at application layer. DNS traffic uses UDP port `53` by default and TCP port `53` as default fallback.

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