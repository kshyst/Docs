# Streams and Pipes

## Standard IO

On linux systems most shells use streams for input and output. These streams can be from and toward various things
like keyboard, devices, blocks ...

Famous streams:

- `stdin`       : 0
- `stdout`    : 1
- `stderr` : 2

## Redirections

- `>` : Redirects STDOUT to file and overwrites if it exists
- `>>` : Redirects STDOUT to file and appends if it exists
- `2>` : Redirects STDERR to file and overwrites if it exists
- `2>>` : Redirects STDERR to file and appends if it exists
- `&>` : Redirects STDOUT and STDERR to file and overwrites if it exists
- `&>>` : Redirects STDOUT and STDERR to file and appends if it exists
- `<` : Redirects STDIN from a file
- `<>` : Redirects STDIN from a file and sends the STDOUT to it


- `&0 &1 &2` : Target locations where STDs go to 

This means: Print stderr to the current location of stdout (terminal) and then
change stdout to file1
```shell
ls 2>&1 > file1
```

### here-document

Give inputs live

```shell
cat << END
file1.txt
helloooo
bye bye
END
```

## Pipes

With `|` you can pass all STDs between commands

### xargs

Reads space, tab, newline and end-of-file delimited strings from standard input 
and executes the provided utility with the strings as arguments.

> Default command is `echo`

- `-I`: Gives input in the middle of the command
- `-L`: Will works with newlines
- `-n 1`: Will run after being provided with one argument

```shell
# creates a folder with every filename in the ls where .txt is removed
ls | sed -e 's/.txt//g' | xargs mkdir

echo wow | xargs -I DATA echo i said DATA  .
```

### tee

Save the output to a file and also see it on the screen

```shell
echo "hi how are you" | tee hi.txt hi2.txt
```

