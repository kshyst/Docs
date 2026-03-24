# Automation

### crontab

`A B C D E commands`

- A : Minute : 0-59
- B : hour : 0-23
- C : Day of month: 1-31
- D : Month: 1-12
- E : Day of week: 0 - 6

If you set asterisk `*` for any of these it means any

- `* * * * *` means every minute.
- `*/15` for A means every 5 minute
- You can give different for each for example `1,2,4,6` for E

> When cron runs, the output will be emailed to the owner of the cron

```shell
# List crons
crontab -l
# Write crons
crontab -e
```

cron files are saved under `var/spool/cron/tabs` or `var/spool/crontabs`


Never edit these files. Use -e

These crons are user specific

### System Wide Crons

Stored in `/etc/crontab`

Writing crons in this file requires giving username too

Inside `/etc/cron.daily` an etc, you can create runnables that will run daily and etc.

### cron access

- `/etc/cron.allow`
- `/etc/cron.deny`

users in allow can use cron and users in deny cant and others cant

This files doesn't exists at default

### at

```shell
at now + 1 min

atq # To see at jobs

atrm 2 # Deletes job number 2

at teatime # 4pm

at tomorrow # At this exact time tomorrow

at 17:30
```

Then write commands. Then `Ctrl + d`

We have allow and deny for this too

- `/etc/at.allow`
- `/etc/at.deny`

## Systemd 

systemd timers are good

There are `.timer` units which are specifically used like crontabs

They use `OnCalendar`

`.timer` units will run the service with same name as itself.