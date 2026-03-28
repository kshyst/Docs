# Character Encoding

## ASCII

7-bit encoding (Later got 8-bit) so it would provide 128 characters by default

## ISO-8859

## UTF-8

Unicode Transformation Format

Backward compatible with **Ascii**

It uses 8-bit code units

# iconv command

Changes character encoding

- `-f`: From
- `-t`: To

```shell
iconv -f WINDOWS-1258 -t UTF-8 file.txt
```