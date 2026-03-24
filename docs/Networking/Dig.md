# What is dig command in linux?

The **dig** command in Linux, short for `Domain Information Groper`, is a versatile tool used for querying Domain Name System (DNS) servers. It's a command-line utility that retrieves information about domain names, such as their IP addresses, mail servers, and other DNS records. dig is commonly used by network administrators and system administrators for troubleshooting DNS-related issues and verifying domain configurations.

```shell
dig aparat.com
```

```text
; <<>> DiG 9.18.30-0ubuntu0.24.04.2-Ubuntu <<>> aparat.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 18324
;; flags: qr rd ra; QUERY: 1, ANSWER: 4, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;aparat.com.                    IN      A

;; ANSWER SECTION:
aparat.com.             6593    IN      A       185.147.178.14
aparat.com.             6593    IN      A       185.147.178.12
aparat.com.             6593    IN      A       185.147.178.11
aparat.com.             6593    IN      A       185.147.178.13

;; Query time: 0 msec
;; SERVER: 127.0.0.53#53(127.0.0.53) (UDP)
;; WHEN: Fri Jun 20 23:12:22 +0330 2025
;; MSG SIZE  rcvd: 103
```