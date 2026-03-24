# Shell

### env

See env vars

```shell
env
```

### set

To change parameters on run-time. 

The switches are main shits about this.

- `-b`: Causes the status of the background terminated immediately rather than seeing it in next shell prompt.
- `-e`: Returns in case a command or pipeline or else returns non-zero.
- `-t`: Exit after reading and executing one command
- `-n`: Reads the commands but doesn't execute them. Used for syntax error checking and debugging. Ignored in interactive shells.
- `-C`: Prevents redirection streams like `>`, `>&` and `<>` from overwriting existing files.


```shell
# This shows all variables on this shell
set | less


### unset

Unsets env variables

```shell
unset MY_VAR
```

### source and .

Runs the executable as part of the current environment. Used for entering virtual environment.

Without using this, The executable will be run in a child process forked from current shell process.

```shell
./hi.sh
source hi.sh
```




### shbang

Tells how it should be run

```shell
#!/bin/bash
```


### alias

```shell
# Shows all the aliases
alias

alias ll='ls -aLF'
```

### functions

```shell
showls() {
    ls $1 | less
}
```

## Different Shell Envs

1. When you login into ssh or any linux machine you enter an **interactive shell login session** called login shell.
2. Running a terminal app in GUI. This is and **interactive shell**.
3. Running a command from a shell, a **sup-shell** starts and runs the command then returns back to main shell. Its a **non interactive shell**.

### Login Shell

1. First it runs `/etc/profile`
2. A line in `/etc/profile` runs whatever is in `/etc/profile.d`
3. `/home/USERNAME/.bash_profile`
4. `/home/USERNAME/.bash_login`
5. `/home/USERNAME/.profile`

> Only one of the 3,4 and 5 will run. If one exists, system won't go further.

6. At the end it runs `~/.bashrc` or `~/.zshrc`. This is where we enter personal configs.

> On non-interactive shells there is a special env var call `BASH_ENV` which will run anything in it.

> bashrc will override PS1 if is set.


### `.bash_logout`
This will run when you logout from a login shell. In many distros it only clears the screen so the new users wouldn't see what you were doing.

## `/etc/skel`
This directory contains files which is used as starting template for new users.



