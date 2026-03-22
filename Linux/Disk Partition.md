# Disk Partition

## Block Devices

A non volatile (won't lose data after power out) which its info can be accessed in any order.

Hard disk, CD, Floppy all use block. We define block size when we are formatting them.

To see blocks

```shell
lsblk
```

## Editing Partition Tables

### fdisk

For MBR (Master Boot Record) systems which use BIOS.

MBR systems can only have 4 primary partitions.

Used for viewing, changing and creating partitions.

- `-l`: Lists partitionsi

```shell
# See partitions on a disk
sudo fdisk /dev/nvme0n1 -l
# Enters fdisk
sudo fdisk /dev/nvme0n1
```

fdisk menu useful commands:

- `n`: Adds new partition
- `p`: Prints partitions
- `t`: Change Type of partition
- `d`: Deletes Partitions

### gdisk

For GPT (GUID Partition Table) systems which use UEFI.


### parted

Mostly used for its change partition size feature

you can use `gparted` in graphical envs.

