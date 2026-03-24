# Regex and File Search

### grep 

There is  dictionaries in `/usr/share/dict`

Switches:

- `-c`: Shows line match count
- `-v`: Reverse the search. Shows not matching lines
- `-n`: Shows line numbers
- `-i`: Case insensitive
- `-l`: Shows only file names
- `-r`: Read all files under each directory recursively

```shell
# Regex format
grep j..a  /usr/share/dict/american-english

# Searches in all english dicts
grep j..a  /usr/share/dict/*english

# With count
grep j..a  /usr/share/dict/*english -c
#/usr/share/dict/american-english:72
#/usr/share/dict/british-english:72

grep -r 192.168. /etc
```

### egrep (Extended grep)

```shell
egrep
grep -E
```

### fgrep (Fixed grep)

doesn't use regex. Just searches

```shell
fgrep
grep -F
```

