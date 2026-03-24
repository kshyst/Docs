# Text Process and Filters

### cat 

Used for concatenate 

Other types:
- `bzcat`
- `xzcat`
- `zcat`
- `gzcat`

### less

For large files. It paginates and starts with the beginning of the file

> Old version from DOS is `more`. `less` is more than `more`.

- `q`: Exit
- `/foo`: Searches for foo
- `n`: Next in search
- `?foo`: Searches backwards for foo
- `G`: Go to the end
- `nG`: Go to line n

### od

Octal Dump

Useful to find problems in text files for example unintended tabs and spaces.

```shell
od salam.txt

# Shows whitespaces
od -a salam.txt
```

### split

Splitting files

```shell
# Split file into chunks of 10 byte and puts in xnn files
split -b 10 salam.txt
# To concatenate all of them
cat x*

# Split file 2 line 2 line
split -l 2 salam.txt

# Splits into 10 parts and puts it in salams{n}{n} files for example salams01
split -d -n 10 salam.txt salams
```

### head and tail

head = سر

tail = دم

```shell
# First 10 lines
head salam.txt
head -5 salam.txt

# Last 10 lines
tail salam.txt
# Follows the tail live
tail -f salam.txt
```

### cut

```shell
# Second field, seperator delimiter is comma
cut -f2 -d, logs.txt 
```
input:
30 Jul, 30 USD

output:
30 USD

> Default delimiter is TAB


### nl

number  line

Prints with line number

```shell
nl salam.txt
```

### sort and uniq

```shell
# Sorts lines lexicographically
sort salam.txt

# doesnt show repeated lines
uniq salam.txt
# With count 
uniq -c salam.txt
```

### paste

Pastes lines from 2 or more files side by side

```shell
paste salam.txt khodafez.txt
```

### tr

Translates. Its a pure filter and doesn't accept input file.

```shell
cat programming.txt | tr 'JS' 'js' 
```

### sed

Stream editor.

Search "sed one linner" for easy use

```shell
# changes salam to hello
sed 's/salam/hello/' salam.txt
# changes globally meaning if there were multiple of salam in one line it will
# change all of them to hello, not just the first one.
sed 's/salam/hello/g' salam.txt

# Use regex
sed -r 's/(A|B)/C/g'

# Suppress the output
sed -rn 's/(A|B)/C/g'
```

### wc

Word Count

```shell
wc salam.txt
# line count
wc -l salam.txt
```