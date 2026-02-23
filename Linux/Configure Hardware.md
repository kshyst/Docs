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

### ls in hardware usage

**lspci**: Shows PCI devices that are connected.

**lsusb**: Shows usb devices

**lsblk**: Shows block devices

**lshw**: Hardware

**lsmod**: All OSs need drives to work with hardware. Linux doesn't load all modules to kernel at the same time. Instead, it uses kernel modules. Loadable kernel modules with `.ko` extension and are used to extend the drivers to the linux kernel. `lsmod` to see loaded modules on kernel.

Deleting and installing module:

```shell
# removing
rmmod <modname>

#installing
insmod <modaddress>
#or
modprobe <modname>
```

To load modules on system boot do one of these:

- add their name to the `/etc/modules` file. (deprecated)
- add their config files to the `/etc/modprobe.d/` directory.