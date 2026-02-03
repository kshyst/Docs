# Run Levels

For SystemV. Defines what tasks can be accomplished in the current state of a Linux system.

## systemd

Has different targets which are groups of services.

```shell
systemctl list-units --type=target
```

Getting the default one

```shell
systemctl get-default 
```

### targets
defines a specific state the system should be in by synchronizing various services and units.

Debian primarily uses systemd targets to manage boot states, with graphical.target (GUI) as the default for desktop installations and multi-user.target (CLI) for servers. Use `systemctl get-default` to view the current setting and `systemctl set-default <target>`.target to change it.

Getting status:

```shell
systemctl status multi-user.target 
```

Isolating targets:

When you use the isolate command, systemd doesn't just start the new target; it immediately stops every service that isn't required by that new target.

Famous targets:

- `rescue`: Local file systems are mounted, no networking, and root-user only (maintenance mode)
- `emergency`: Only the root file system and in read-only mode, No networking and root-user only(maintenance mode)
- `reboot`
- `halt`: Stops all processes and halt CPU activities
- `poweroff`: Like halt but also sends an ACPI shutdown signal (No lights!)

```shell
systemctl is-system-running
systemctl isolate emergency
systemctl is-system-running
```

Think of isolate as a "reset" for your current session.

- **Dependency Check**: systemd looks at the target you want to reach.

- **Start Phase**: It starts all services required by that target.

- **Stop Phase**: It kills everything else currently running that isn't a dependency of the new target.

Entering Terminal 

```shell
sudo systemctl isolate multi-user.target
```

Entering GUI

```shell
sudo systemctl isolate graphical.target
```

![isolate](../img/isolate.png)

## SysV runlevels

On SysV we were able to define different stages. On a Red Hat-based system we usually had 7:


RedHat:
0. Shutdown
1. Single-user mode (recovery); Also called S or s
2. Multi-user without networking
3. Multi-user with networking
4. to be customized by the admin
5. Multi-user with networking and graphics
6. Reboot

Debian:
0. Shutdown
1. Single-user mode
2. Multi-user mode with graphics
6. Reboot

### Checking status and setting defaults

Check current runlevel(also works on systemd machines)
default is in `/etc/inittab`

```shell
runlevel
```

