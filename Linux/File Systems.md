# File Systems

After partitioning disk we need to give each partition a file system.

# File Systems Comparison (Including Legacy & Swap)

| File System | Primary OS | Max File Size | Max Volume Size | Journaling | Best Use Case |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **FAT32** | Universal | 4 GB | 2 TB | No | Small USB flash drives, older devices, high compatibility. |
| **exFAT** | Universal | 16 EB | 128 PB | No | Large external hard drives and SDXC cards needing Mac/Windows compatibility. |
| **NTFS** | Windows | 8 PB | 8 PB | Yes | Windows internal system drives and large Windows-only external drives. |
| **MINIX (v1/v2)** | MINIX / Early Linux | 64 MB (v1) to 4 GB (v2) | 64 MB (v1) to 4 GB (v2) | No | Historical learning, OS development, very early Linux systems. |
| **ext2** | Linux | 2 TB | 32 TB | No | Flash drives or SD cards where write-heavy journaling overhead is unwanted. |
| **ext3** | Linux | 2 TB | 32 TB | Yes | Legacy Linux systems transitioning from ext2 needing crash protection. |
| **ext4** | Linux | 16 TB | 1 EB | Yes | Standard, modern Linux distributions and servers. |
| **Btrfs** | Linux | 16 EB | 16 EB | Yes (CoW)* | Advanced Linux systems, NAS devices, systems needing snapshot capabilities. |
| **APFS** | macOS | 8 EB | 8 EB | Yes (CoW)* | Modern Apple devices (macOS, iOS) optimized for Solid State Drives (SSDs). |
| **HFS+** | macOS | 8 EB | 8 EB | Yes | Older Apple computers with mechanical hard drives (pre-macOS High Sierra). |

* **CoW**: Copy-on-Write. Instead of traditional journaling, these file systems write new data to a new location before updating the pointers, offering excellent protection against data corruption.



### Commands

To format partitions with filesystems:

> When using `-t` with mkfs is actually using another command. each filesystem has a command

```shell
# See all mkfs s
ls -l /sbin/mk*

mkfs # For filesystem
mkswap # For swap
```
