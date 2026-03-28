# System Logging

## History

Logs are an important part of Unix philosophy and design. The Kernel, services, programs and most events (let it be a crash or a login attempt) will generate logs. These logs can be examines / monitored to gain insights about the system and its status. If you watch a Linux guy on a windows machine for 1 hour, you will hear "where are the logs?".

If every single program tries to handle its own logs we will face 2 issues (at least!); 1. programming will be difficult because you need to invent (or import) the wheel and 2. different programmers will generate different kinds, formats and location for their logs and it will be difficult to keep them under your control. To overcome this, Unix world had its own solution: central logging service.

The older utility to manage logs was called `syslog`, then we had `syslog-ng` (new-generation) and later `rsyslog` which is what most distros were using till a couple of years ago. `rsyslog` is still present and is part of the GNU/Linux universe. It can receive logs from different programs (even over network) and save them in different places (files or files in network or even run some actions on them). It's logs are saved in the text format and can be managed by the tool of your choice, let it be `grep`, `less`, `tail`, `zless` (less a compressed file without explicitly opening it), `zcat` (catting a file without explicitly opening it), `grep`, ....

But these days `systemd` is conquering most of the linux world and logs are not an exception. `SystemD`s logging service is called `journald` and we use the `journalctl` command to read its logs; unfortunately they are not text files anymore.

> Its fun to know how the logs work in more depth. Here it is: whoever wants to log something eventually sends its messages to the /dev/log or /dev/kmsg devices (often by using a helper tool). The logging tool (say rsyslog) will read from these devices and process the logs based on its configs; and now, the messages are logs!

## Boot Logs

More info in boot docs but mostly its about `dmesg`.

## Log Rotation

Log cleanuper

Main config is under `/etc/logrotate.conf`:

Specific service create their logs in `/etc/logrotate.d`

```shell
logrotate

cat /etc/logrotate.d/ufw
#/var/log/ufw.log
#{
#        rotate 4
#        weekly
#        missingok
#        notifempty
#        compress
#        delaycompress
#        sharedscripts
#        postrotate
#                [ -x /usr/lib/rsyslog/rsyslog-rotate ] && /usr/lib/rsyslog/rsyslog-rotate || true
#        endscript
#}

```

| Parameter | Meaning |
|-----------|--------|
| `weekly` | Rotate logs weekly |
| `missingok` | It is fine if there is no log for this week |
| `rotate 52` | Keep the latest 52 logs and delete the older ones |
| `compress` | Compress the logs |
| `create 0640 www-data adm` | Create the files with these permissions and ownership |
| `pre & post rotate` | Run these scripts or commands before and after the rotation |

The logrotate runs using crons and does its job on a daily basis based on a configuration on `/etc/cron.daily/logrotate`


## Famous log files

Generally logs are saved at `/var/log`. In case of any issue and if you do not know what to do, it is a common practice to run a `ls -ltrh /var/log/` to see if any program generated a new log or not.

`/var/log/auth.log` (in Debian Based)
Authentication processes will log here. Things like cron jobs, failed logins, sudo informations, ...

`/var/log/syslog`
A centralized place most of the logs received by rsyslogd if a specific logfile is not provided in `/etc/rsyslog.conf`

`/var/log/debug`
Debug information from programs.

`/var/log/kern.log`
Kernel messages.

`/var/log/messages`
Informative messages from services. This is also the default place for logs for remote clients.

`/var/run/utmp & /var/log/wtmp`
Successful logins.

`/var/log/btmp`
Failed logins. You can check this to see if anyone is trying to guess your passwords!

`/var/log/faillog`
Failed authentication attempts.

`/var/log/lastlog`
Date and time of recent user logins.

### Service logs
Service logs do create files or directories at `/var/log` and update their logs there. For example there might be a `/var/log/apache2/` (or `/var/log/httpd`) for the Apache HTTP server or a `/var/log/mysql` for the MySQL DB.

## rsyslog

This is the newer generation of syslogs and is used in many environments. Its main configuration is located at `/etc/rsyslog.conf` and it also reads anything in `/etc/rssylog.d/` directory. Its configuration consists of 3 main sections:

- `MODULES`, modules used. For example to let the rsyslog use UDP connections
- `GLOBAL DIRECTIVES`, general configurations like directories accesses
- `RULES`, A combination of facilities, priorities and actions tells the rsyslog what to do with each log.

```shell
root@debian:~# cat /etc/rsyslog.conf
# /etc/rsyslog.conf configuration file for rsyslog
#
# For more information install rsyslog-doc and see
# /usr/share/doc/rsyslog-doc/html/configuration/index.html


#################
#### MODULES ####
#################

module(load="imuxsock") # provides support for local system logging
module(load="imklog")   # provides kernel logging support
#module(load="immark")  # provides --MARK-- message capability

# provides UDP syslog reception
#module(load="imudp")
#input(type="imudp" port="514")

# provides TCP syslog reception
#module(load="imtcp")
#input(type="imtcp" port="514")


###########################
#### GLOBAL DIRECTIVES ####
###########################

#
# Set the default permissions for all log files.
#
$FileOwner root
$FileGroup adm
$FileCreateMode 0640
$DirCreateMode 0755
$Umask 0022

#
# Where to place spool and state files
#
$WorkDirectory /var/spool/rsyslog

#
# Include all config files in /etc/rsyslog.d/
#
$IncludeConfig /etc/rsyslog.d/*.conf


###############
#### RULES ####
###############

#
# Log anything besides private authentication messages to a single log file
#
*.*;auth,authpriv.none      -/var/log/syslog

#
# Log commonly used facilities to their own log file
#
auth,authpriv.*         /var/log/auth.log
cron.*              -/var/log/cron.log
kern.*              -/var/log/kern.log
mail.*              -/var/log/mail.log
user.*              -/var/log/user.log

#
# Emergencies are sent to everybody logged in.
#
*.emerg             :omusrmsg:*
```

Facilities can be like : kern, user, mail, daemon, cron, auth, ntp, security, console, syslog, ...

Priorities can be : `emerg/panic`, `alert`, `crit`, `err/error`, `warn/warning`, `notice`, `info` or `debug`.

| Action    | Sample               | Meaning                                                                 |
|-----------|-----------------------|-------------------------------------------------------------------------|
| filename  | `/usr/log/logins.log` | Will write the log to this file                                         |
| username  | `jadi`                | Will notify that person on the screen                                   |
| @ip       | `@192.168.1.100`      | Will send this log to that log server, which will handle it per its config |


Example:
A line like this will show the kernel panics to a remote log server and also will log everything on every level to a log file:

```shell
kern.panic      @192.168.1.100
*.*             /var/log/messages
```

If you log some specific priority, all the more important things will be logged too! So if you write `cron.notice /var/log/cron/cron.log`, you are logging the emerg/panic, alert, critical, error, warning and notice logs of the cron category too.

According to Syslog(3) man page:

```shell
  This determines the importance of the message.  The levels are, in order of decreasing importance:

       emerge      system is unusable

       alert      action must be taken immediately

       crit       critical conditions

       error        error conditions

       warning    warning conditions

       notice     normal, but significant, condition

       info       informational message

       debug      debug-level message


       panic      deprecated synonym for emerg
       error      deprecated synonym for err
       warn       deprecated synonym for warning
```

If you need to log ONLY one specific level, add an equal sign (=) before the priority like this `local3.=alert /var/log/user.alert.log`.

It is important to know that the binary which logs the *kern category is a standalone daemon. This daemon is called `klogd` and uses same configuration files. Why? so even after everything is crashed, `klogd` will be able to log the kernel's crashes ;).

### logger

If you need to send something toward the `rsyslog`, you can use the `logger` tool.

```shell
root@debian:# logger Testing my lovely tool
root@debian:# logger local1.emerg Nothing emergent for sure
root@debian:# tail -3 /var/log/syslog
2023-06-30T06:05:01.325358-04:00 debian CRON[914]: (root) CMD (command -v debian-sa1 > /dev/null && debian-sa1 1 1)
2023-06-30T06:13:27.671410-04:00 debian root: Testing my lovely tool
2023-06-30T06:13:45.795158-04:00 debian root: local1.emerg Nothing emergent for sure
```

### journald

Systemd logs are visible through `journalctl` which is part of `systemd-journald`

```shell
systemctl status systemd-journald
```

Configuration file is under `/etc/systemd/journald.conf`.

Journalctl usage:

| Switch | Meaning |
|--------|---------|
| `-r` | Show in reverse; newer on top |
| `-f` | Keep showing the tail (follow) |
| `-e` | Show and go to the end |
| `-n` | Show this many lines from the end (newest ones) |
| `-k` | Show kernel messages (same as `dmesg`) |
| `-b` | Show logs from a specific boot; `-1` is previous boot, `0` is current. List boots with `--list-boots` |
| `-p` | Filter by priority, e.g., `-p err` |
| `-u` | Show logs from a specific systemd unit |

You can also use the `--since` and -`-until` to show a specific time range, 
It is possible to provide time as `YYYY-MM-DD HH:MM:SS` or use things like
`yesterday`, `today`, `tomorrow`, `now` or even `--since "10 minutes ago"` which is 
equals to `--since "-10 minutes"`.

You can also add the name of the program to the command to see the logs from that specific program; say `journalctl /usr/bin/xrdp` or use some fields like `PRIORITY=`, `_PID=` and provide values.

The systemd journals can keep their logs in memory or write them to the disk or drop all the logs. This is determined by the configurations in `/etc/systemd/journald.conf`. But the default behaviour is as follow:

### systemd-cat

This tools is used when you want to manually send logs to the journaling system. It sends its input to the journal or if provided, runs the command and sends the result to the journal.

```shell
echo "This is my first test" | systemd-cat
systemd-cat -p info uptime #sending priority too
journalctl -n 3
```

### Storing logs

The systemd journals can keep their logs in memory or write them to the disk or drop all the logs. This is determined by the configurations in `/etc/systemd/journald.conf`. But the default behaviour is as follow:

- The system checks the `/var/log/journal`. If this directory, exists logs will be saved in a directory inside this. The name of the directory will be decided by looking at `/etc/machine-id`
- If the `/var/log/journal` is not there, the logs will be saved in memory at `/run/log/jouranl` and in the directory decided by `/etc/machine-id`.

If the logs are being written to disk, these variable will manage the disk usage:

| Variable Name       | Usage                                                                 |
|----------------------|------------------------------------------------------------------------|
| `SystemMaxUse`       | Maximum total disk usage, e.g., 500M. Default is 10% of the filesystem |
| `SystemKeepFree`     | Amount of free space to keep, e.g., 1G. Default is 15%                 |
| `SystemMaxFileSize`  | Maximum size of each individual log file. Default is 1/8 of SystemMaxUse |
| `SystemMaxFiles`     | Maximum number of non-active log files. Default is 100                 |

In case you need to do the clean-up manually (with the help of systemd tools for sure), you can run `journalctl` with these switches:

| Switch | Usage |
|-------|-------|
| `--vacuum-time` | Remove logs older than a specified time. Example: `--vacuum-time=3months`. You can use `s` (seconds), `m` (minutes), `h` (hours), `d`/`days`, `weeks`/`w`, `months`, and `years`/`y`. |
| `--vacuum-size` | Remove old logs until the total size reaches a specified limit, e.g., `--vacuum-size=1G`. |
| `--vacuum-files` | Keep only the specified number of archived log files. |


## Logs from recovered system

If you have a crashed / non-booting system, you can still examine its logs; if the files are there. As you know the files are in `/var/log/journal/{mchine-id}` where machine-id is in `/etc/machine-id` of the crashed machine.


You can move these files to a directory after booting the crashed machine with a live linux or mount its hard on another machine or even examine them in place. The `-D` (or `--directory`) switch of journalctl indicates the location of the journal files. So if your crashed machine id is ec22e43962c64359b9b25cfa650b025b and you've mounted its `/var/` under your `/mnt/var/` directory, you can issue this command to read its logs and see what had happened:

```shell
journalctl -D=/mnt/var/log/journal/ec22e43962c64359b9b25cfa650b025b/
```

You can also use the `--merge` switch to merge these logs into your machine or use `--file` to check only one specific journal file. Lastly if the exact location of journal files are not known, you can use `--root /mnt` and tell the `journalctl` to search there for journal files.