# Boot

### Boot Process

1. Motherboards firmware does the POST (Power On Self Test)
2. Motherboard loads the bootloader (Loads on memory and is done because new OSs are too complicated)
3. Bootloader loads the linux kernel based on its configs and commands
4. Kernel loads and prepares the system (root filesystem) and runs the initial programs
5. Init program start the services like webservers, networking, GUI, etc.

### Motherboard Firmware types

#### BIOS

Basic Input Output System

- Older
- Limited to on sector of dist and needs multistage boot loader
- Can start the bootloader from internal/external HDD, CD/DVD, USB Flash, Network Server
- If booting from HDD, Master Boot Record **(MBR)** will be used.

#### UEFI

Unified Extensible Firmware Interface

- Modern
- Specifies a special disk partition for the bootloader called EFI System Partition(ESP)
- ESP is FAT and mounted on `/boot/efi` and bootloader files has `/efi` extension.

> Check `/sys/firmware/efi` to see firmware


### Bootloader

Initializes the minimum hardware needed to boot the system. Then locates and runs the OS.

We can point UEFI to run any OS we want but typically under GNU/Linux systems we use GRUB bootloader.

### Kernel

Core of the OS. Bootloader loads the kernel in the memory and runs is. But the kernel need some initial info to start; Things like drivers which are necessary to work with the hardware. Those are stored in `initrd` or `initramfs` alongside the kernel and used during the boot.

You can send parameters to kernel during the boot using the Grub configs. For example, sending a 1 or S will result the system booting in single-user mode (recovery). Or you can force your graphics to work in 1024×768x24 mode by passing vga=792 to the Kernel during the boot.

Kernel logs are in `/var/log/messages`. Some systems is in `/var/log/syslog`.

**dmesg**: Shows the boot process logs during boot. Kernel stores the logs into "Kernel Ring Buffer". The syslog daemon collects the boot logs and stores them in `/var/log/dmesg`.

> Kernel Ring Buffer is a buffer in kernel and dedicated to it.

**journalctl**:

- `-k`: for Kernel logs
- `-b` for boot logs
- `-u kernel` all the previous

Also, can be find in `/var/log/boot` or `/var/log/boot.log`.

### Init

After kernels initialization, it runs the **Initialization Daemon Process** and takes care of starting other daemons.

#### Init Systems

- **SysVinit** based on Unix system V. Outdated. Maybe on older machines.
- **upstart** as replacement for SysVinit and made by canonical. Discontinued due to wide adoption of Systemd. Can still be found in ChromeOS.
- **Systemd** Hated by linux elites for not following Unix principles. Widley adopted.

```shell
which init

readlink -f /sbin/init

ps -p 1
```

You can see the process hierarchy using:
```shell
pstree
```


### systemd

#### Units
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

