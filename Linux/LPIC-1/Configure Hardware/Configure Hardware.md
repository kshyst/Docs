# Configure Hardware

### Firmware

A software on hardware which runs it. Keyboard, Motherboard ... has firmware.

**BIOS**: Basic input output system. Is a firmware. Very old and outdated. Boots the computer from the first partition from hard disk.

**UEFI**: Unified Extensible Firmware Interface. Uses a specific disk partition for boot (EFI System Partition(ESP))
and uses FAT file system. On linux its on  `/boot` and files are `.efi`.

### I/O Interfaces

**PCI**: Peripheral Component Interconnect. Letting hardware boards to be added to motherboard.

**GPIO**: General Purpose Input Output. Pins on raspi and etc.


**sysfs**: Is a pseudo filesystem. In kernel memory. Shows info about kernel subsystems, hardware devices and associated device drivers. Mounted in `/sys`. Block devices are under `block` and PCI, USB and etc are under `bus`.

```shell
ls /sys
block  bus  class  dev  devices  firmware  fs  hypervisor  kernel  module  power
```

**udev**: Device manager for linux kernel. Successor of devfsd and hotplug. Operates under `/dev` directory. handles device nodes and userspace events when adding or removing hardware devices including firmware loading for devices.

For example start back process when attaching a USB drive.

**dbus**: D-Bus is a message bus system. For apps to talk to each other. Addition to interprocess communication. Helps coordinate a processes lifecycle. Makes it easy to code a single instance app or daemon and launch them when their service is needed.

**proc**: Where kernel keeps its settings and properties. Under `/proc`. This dir is created on ram and files might have write access.

- IRQs (Interrupt Requests)
- I/O ports (locations in memory where CPU can talk with devices)
- DMA (Direct Memory Access, faster than IO)
- Processes
- Network Settings

```shell
cd /proc
cat cpuinfo

# ipv4 settings on device
cd /proc/sys/net/ipv4

# power settings
cd /proc/acpi
```