# Shell Scripting

### Combining Commands

```shell
# This will run next anyway
cd /tmp; ls
# This will run only if first was success
cd /tmp && ls
# This will run only if first was failure
cd /tmp || ls
```

### Shebang

Starts with `#!` and tells the script that which interpreter must be used.

```shell
#!/bin/bash

NAME=kiarash

echo 
echo "hello $NAME"
```

### Variables

```shell
NAME=kiarash
echo $NAME
```

Use `$1` and `$2` and ... to get commandline arguments.

### Command Substitution

Storing output of a command in a variable.

```shell
FILES=$(ls -l)
FILES=`ls -l`

t=`date +'%Y-%M-%d-%H%M'`
touch backup_$t.tar
```


### Conditions

```shell

if [ condition ]
then
    command
else
    command
fi
```

> `$#` shows argument count

Conditioning:

- `=`: Equality of two strings
- `!=`: Not Equality of two strings
- `-lt -gt -le -ge -ne -eq`: Number comparison
- `-f FILENAME`: Checks if FILENAME exists
- `-s FILENAME`: If FILENAME exists and its size is more than 0
- `-x FILENAME`: FILENAME exists and is executable.

### Exit

`0` means program ran correctly. Anything else means error.

```shell
exit 0
exit 1
```

To see the return value of the latest ran program:

```shell
echo $?
```

### Read

```shell
read NAME

echo $NAME
```

### Loops

```shell
# For loops
for VAR in SOME_LIST;
do
    some stuff with $VAR
    commands
done

# While loops
while [ condition ]
do
    commands
done
```


### let

using `let` you can work with numbers:

```shell
VAR=10

let VAR=VAR-3

# If you do this you get 10-3
VAR=$VAR-3





