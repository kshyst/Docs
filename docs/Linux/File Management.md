# File Management


### ls

```shell
# long, sorted by time, reverse sort, human readable
ls -ltrh
```

### file

Tells file format

```shell
file salam.txt
# prints the mime format
file -i salam.txt
```

### dd

cooler copy

```shell
dd if=tasks.txt of=my-tasks
# write zeros from zero device
dd if=/dev/zero of=zeros count=100 bs=1
od zeros
# Boot a flash
sudo dd if=ubuntu.iso of=/dev/sbc bs 2048
```

> Using this you are directly writing on blocks so if you mess
> something you have to reformat the device


### find

```shell
find . -iname "*hello*"
# Only dirs
find . -iname "*hello*" -type d
#Based on size. more than 100 kilobytes
find . -iname "*hello*" -type d -size +100k
# Find empty
find . -empty
```

Types: `d` or `f`

Time flags for find:
- `-newer`: newer than the referenced file
- `-amin`: Access time
- `-cmin`: Status changed (for example the permissions changed)
- `-mmin`: Modified minutes

Same for atime, ctime and mtime

> -amin 40 means exactly 40 minutes, -amin +40 more than 40 minutes ago, -amin -40 less than 40 minutes ago

```shell
find . -newer salam.txt
```

More flags:

- `-ls`: Shows the ls for the files
- `-print`: Prints idk
- `-exec`: Runs something on it

```shell
find . -mmin +100000 -exec echo "yooooo {}" \;
# Actually useful usage :
find . -name "*.htm" -exec mv "{} {}l" \;
find . -empty -exec rm '{}' \; 
```

Search for certain permissions:
```shell
sudo find / -perm -u+s
```

### locate

Faster find. It creates a database using `updatedb` everyday from all the files and dirs in the system. 

The update package is called `plocate` on debian systems.

The configuration is under `/etc/updatedb.conf` or `/etc/syconfig/locate`

```shell
plocate docker
# same as
find . "*docker*"

sudo updatedb
```
