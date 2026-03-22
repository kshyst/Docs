# Process Management

Process Types:

- `D` :    uninterruptible sleep (usually IO)
- `I` :    Idle kernel thread
- `R` :    running or runnable (on run queue)
- `S` :    interruptible sleep (waiting for an event to complete)
- `T` :    stopped by job control signal
- `t` :    stopped by debugger during the tracing
- `W` :    paging (not valid since the 2.6.xx kernel)
- `X` :    dead (should never be seen)
- `Z` :    defunct ("zombie") process, terminated but not reaped by its parent

Additional Characters that displays:

- `<` :   high-priority (not nice to other users)
- `N` :   low-priority (nice to other users)
- `L` :   has pages locked into memory (for real-time and custom IO)
- `s` :   is a session leader
- `l` :   is multi-threaded (using CLONE_THREAD, like NPTL pthreads do)
- `+` :   is in the foreground process group

#### Zombie Process

Finished and returned something, but the parent process
hasn't read the return. Process is dead but the return value still exists. 
Doesn't put pressure on system. It is just bad programming problem.


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

## Change Process Priorities

### nice

Niceness is `-20` to `19`

The you are nicer, the lower is the priority

If CPU is struggling, i will give more power to lower nicenesses

- -20 : ANGRY
- 19 : SUPER NICE

Example: You are backupping system using tar, Its better to give high niceness 

```shell
# Run ls with 14 niceness
nice -n 14 ls
# Default niceness is 10
nice ls
# Without nice is 0
ls
```

> Not root users cant give minus niceness
> 
> `nice: cannot set niceness: Permission denied`

### renice

renices

Without sudo, you can only add to the niceness

```shell
renice -n 10 3783
# Nice renice
pgrep xeyes | xargs renice -n 19 
```