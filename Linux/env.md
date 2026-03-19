# env

## Setting env vars

First is only recognizable in that shell process but `export` is usable in 
processes forked from shell.

```shell
NAME=kiarash

export NAME=kiarash
```


## Intresting Vars:

### $PS1

Shows the template of your shell line

### $

Shows process id

```shell
echo $$
```