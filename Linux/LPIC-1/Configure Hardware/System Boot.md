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

**dmesg**: Shows the boot process logs during boot. Kernel stores the logs into "Kernel Ring Buffer". The syslog daemon collects the boot logs and stores them in `/var/log/dmesg`.

**journalctl**:

- `-k`: for Kernel logs
- `-b` for boot logs
- `-u kernel` all the previous


