# Wildcards

- `*` any string
- `?` any single character
- `[ABC]` matches A, B and C
- `[!x]` means not x

```shell
# any starts with s
ls [s]* 
# salam.txt

# Copies everything with 3 character then a dot then anything (even nothing)
# to temp
cp ???.* /temp
```