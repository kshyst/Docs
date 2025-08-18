# File Transfer Protocol

Can transfer files faster than HTTP.

FTP Commands:

- `USER`: To input username
- `PASS`: To input password
- `RETR`: To download a file from FTP server to client
- `STOR`: To upload a file into FTP server
- `LIST`: Runs ls on server

By default it uses port 21 using TCP.

Simple FTP to get a txt file:

```shell
ftp 1.1.1.1
type ascii
get kiarash.txt
```