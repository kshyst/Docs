# Process Management

## Manipulating processes

### bg, fg and jobs

- Stopping process with `Ctrl + Z`
- Finishing process with `Ctrl + C`

Sending the suspended job to foreground

Without specifing the number when sending to fg, it will send the one with + in front of it in `jobs`
```shell
fg %1 # the number is the number shown in jobs
bg %1 # the number is the number shown in jobs
```

Sending process to background with `&`
```shell
xeyes &
```

Showing jobs running in this shell 

```shell
jobs
# Will show process ids
jobs -l
```

If we close the shell, all the fg and bg processes will be closed too because they are forked from it.

### nohup

It does not answer to `HUP` or `1` signal.

To not end the process after parent is killed:

```shell
nohup xeyes &

# Famous nohup usage:
nohup script.sh > mynohup.out 2>&1 &
```

### kill

format is `kill -SIGNALID_OR_NAME process_id`

- `1` or `HUP`: Informing the process that its controlling terminal is terminated
- `15` or `TERM`: Normal termination request
- `9` or `KILL`: Forcefully kills the process

### killall

```shell
killall xeyes
```

### pkill

Kills using pattern

```shell
# Will kill all xeyes s
pkill xey
```

## Monitoring Processes

### ps

```shell
# All processes in this shell
ps
# All processes in the system
ps -e
# Gives more info
ps -ef
```

### pgrep

```shell
pgrep xclock
```


### top, htop

Shows how many process was waiting
```
Load Average: Last_1_Minute Last_5_Minute Last_15_Minute
```

How to check if it's bad: If we have 8 core and the first number is more than 8 like 9, then system is under pressure

### free

Shows free memory

```shell
# Human Readable
free -h
```

### watch

Watch the program output in realtime displaying its output and errors. By default, command is run every 2 seconds and watch
will run until interrupted

```shell
watch ls
```

- `-n`: interval seconds
- `-b`: beep in case of non-zero exit
- `-d`: shows difference between runs

### uptime

shows uptime

## Multiplexer

### screen

idc

tmux is better

### tmux

commands work with `CTRL + B`

- `%`: Splits the screen vertically
- `"`: Splits the screen horizontally
- `ARROW KEYS`: Move between screens
- `d`: Detach from tmux
- `c`: Creates a new window
- `{window number}`: Switches to wanted window
- `t`: Turns screen into clock

#### commands

```shell
# Starts tmux
tmux
# Lists
tmux ls
# Attach
tmux attach -t 0
```