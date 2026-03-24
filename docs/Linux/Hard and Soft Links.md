# Hard and Soft Links

A link is basically an additional directory entry for a file or folder.

```shell
ln myfile hardlink
```

Both hardlink and original file have the same inode:

So if we change one, the other would be changed exactly.
```shell
ls -i
```

To create softlink:

```shell
ln -s myfile softlink
```

The result looks like this :
```shell
9607674 -rw-rw-r-- 2 kshyst kshyst 7062 Mar 21 19:31 myfiles.txt
9602633 lrwxrwxrwx 1 kshyst kshyst   11 Mar 23 17:58 softmyfile.txt -> myfiles.txt
9607674 -rw-rw-r-- 2 kshyst kshyst 7062 Mar 21 19:31 yourfiles.txt
```

The softlink will have different inode.

> We can't create hardlink from directories and files in another filesystem.

To unlink:

```shell
unlink softlink
```

find command has type link (l):

```shell
find . -type l
```

`python3` command is always a link to python3.xx which is exact version.

```shell
ls -l /usr/bin/python3
# lrwxrwxrwx 1 root root 10 Nov 12 15:45 /usr/bin/python3 -> python3.12
```
