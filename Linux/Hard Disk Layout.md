# Design Hard Disk

## FHS (File system Hierarchy Standard)

![](img/fhs.png)

`media` is for known mountables which system can automatically mount. `mnt` is used for user configured mounts.

`opt` when user installs something like apps.

Most used thing in `var` is `/var/log/`

## Disk Types

ATA is `Advanced Technology Attachment`. The standard interface to connect storage devices like hard drives, SSDs, CD/DVD and etc to a computer's motherboard.

- **PATA (Parallel ATA)**: The old standard and also known as **IDE**. Data send it parallel meaning multiple bits at once. Linux namings are `/dev/hda` and `/dev/hdb`.
- **SATA (serial ATA)**: The modern standard that replaced PATA. Sends data serially meaning one bit at a time but at much higher frequencies. Linux namings are `/dev/sda`, `/dev/sdb`. Even though it's an ATA, the linux machine uses SCSI subsystem driver to manage SATA drives which is why we get the `sd` prefix. 
- **SCSI (Small Computer System Interface)**: The Enterprise standard pronounced "Skuzzy". Originally designed for servers and high-end workstations to handle many devices. Modern servers now use `SAS (Serial Attached SCSI)`.
- **NVME (Non-Volatile Memory Express)**: Doesn't use cables and directly talks to the PCIe bus. In linux appears as `/dev/nvme0nx`.

## Partitions

You have to PARTITION the disks, that is to create smaller parts on a big disk. These are self-contained sections on the main drive. OS sees these as standalone disks. We recognize them as: - `/dev/sda1` (first partition of the first SCSI disk) - `/dev/hdb3` (3rd partition on the second disk)

BIOS systems were using MBR and could have up to 4 partitions on each disk, although instead of creating 4 Primary partitions, you could create an Extended partition and define more Logical partitions inside it.

Using MBR:

- `/dev/sda3` is the 3rd primary partition on the first disk
- `/dev/sdb5` is the first logical partition on the second disk
- `/dev/sda7` is the 3rd logical partition of the first physical disk

>If you define an extended partition on a BIOS system, that will be /dev/sdx5 (1-4 for primary, and the first extended will be 5).

The new standard is GPT.

UEFI systems use `GUID Partition Table (GPT)` which supports 128 partitions on each device.

Linux systems can mount this partitions on different paths. Say you can have a separated disk with one huge partition for your `/home` and another one for your `/var/logs/`.

```shell
fdisk /dev/sda
Welcome to fdisk (util-linux 2.25.1).
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.


Command (m for help): p
Disk /dev/sda: 298.1 GiB, 320072933376 bytes, 625142448 sectors
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0x000beca1

Device     Boot     Start       End   Sectors   Size Id Type
/dev/sda1  *         2048  43094015  43091968  20.6G 83 Linux
/dev/sda2        43094016  92078390  48984375  23.4G 83 Linux
/dev/sda3        92080126 625141759 533061634 254.2G  5 Extended
/dev/sda5        92080128 107702271  15622144   7.5G 82 Linux swap / Solaris
/dev/sda6       107704320 625141759 517437440 246.8G 83 Linux
```

To see partition type:

The PTTYPE column shows partitions type of each partition.

```shell
lsblk -o NAME,SIZE,PTTYPE,PARTTYPENAME
```

Some Commands:

- `parted`: See info about partition but doesn't understand GPT.
- `fdisk`: Same but better.

### LVM (Logical Volume Manager)

In many cases, you need to resize your partitions or even install new disks and add them to your current mount points; Increasing the total size. LVM is designed for this.

LVM helps you create one partition from different disks and add or remove space from/to them. The main concepts are:

It uses 3-layer abstraction:

- `Physical Volume (PV)`: `sudo pvs`, A whole drive or a partition. It is better to define partitions and not use whole disks - unpartitioned.
- `Volume Groups (VG)`: `sudo vgs`, This is the collection of one or more PVs. OS will see the vg as one big disk. PVs in one VG, can have different sizes or even be on different physical disks.
- `Logical Volumes (LV)`: `sudo lvs` OS will see lvs as partitions. You can format an LV with your OS and use it.


### Design Hard disk layout

#### swap 

Swap in Linux works like an extended memory. The Kernel will page memory to this partition/file. It is enough to format one partition with swap file system and define it in /etc/fstab

> Note: There is no strict formula for swap size. People used to say "double the ram but not more than 8GB". On recent machines with SSDs, some say "RAM + 2" (Hibernation + some extra ) or "RAM * 2" depending on your usage.

#### /boot

Older Linux systems were not able to handle HUGE disks during boot (say Terabytes) so the /boot partition was separated. this separation comes in handy in situations such as recovering broken systems. you can even make /boot read-only. Most of the time, having 100MB for /boot is enough. This partition can be on a different disk or a separated partition.

This partition should be accessible by BIOS/UEFI during boot (No network drive).

On UEFI systems, there is a `/boot/efi` mount point called EFI (Extensible Firmware Interface) system partition or ESP. This contains the bootloader and kernel and should be accessible by the UEFI firmware during boot.


### Swaps
