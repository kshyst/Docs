# systemd

systemd is a system and service manager for Linux operating systems. When run as first process on boot (as PID 1), it acts as init system that brings up and maintains userspace services. Separate instances are started for
logged-in users to start their services.

systemd is usually not invoked directly by the user, but is installed as the /sbin/init symlink and started during early boot. The user manager instances are started automatically through the `user@.service` service.


The main command used to introspect and control systemd is `systemctl`.

The main command to see logs of a daemon process is `journalctl`.

### Logs of a service

```shell
journalctl -u nginx.service
```

### Units
Made around *units*. A unit can be a service, group of services or and action. Units have a name, type, and a configuration file.

There are 12 types of units:

- automount
- device
- mount
- path
- scope
- service
- slice
- snapshot
- socket
- swag
- target
- timer

We use `systemctl` to work with these units and `journalctl` to see logs

```shell
systemctl list-units
systemctl list-units --type=target
systemctl get-default # default targets
systemctl list-unit-files
systemctl cat boot.mount
```

Units can be found in:

- `/etc/systemd/system/`
- `/run/systemd/system/`
- `/usr/lib/systemd/system`


Working with services:
```shell
systemctl stop sshd
systemctl start sshd
systemctl status sshd
systemctl is-active sshd
systemctl is-failed sshd
systemctl restart sshd
systemctl reload sshd
systemctl enable sshd
systemctl disable sshd
systemctl daemon-reload sshd
```

Using journalctl:

```shell
journalctl # show all journal
journalctl --no-pager # do not use less
journalctl -n 10 # only 10 lines
journalctl -S -1d # last 1 day
journalctl -xe # last few logs
journalctl -u ntp # only ntp unit
journalctl _PID=1234
```

### SysV

Control files are located at `/etc/init.d` and are closer to the general bash scripts.

```shell
/etc/init.d/ntpd status
/etc/init.d/ntpd stop
/etc/init.d/ntpd start
/etc/init.d/ntpd restart
```

