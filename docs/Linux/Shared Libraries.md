# Shared Libraries

To determine shared libraries that executable programs depend on and install them when necessary.

## Linking

When we write a program, we use libraries. For example, if you need to read text from standard input, you need to link a library that provides this. Think linking has two forms:

- **Static** Linking is when you add this library to your executable program. In this method, your program size is big because it has all the needed libraries. One good advantage is your program can be run without being dependent on other programs/libraries.
- **Dynamic** Linking is when you just say in your program "We need this and that library to run this program". This way your program is smaller but you need to install those libraries separately. This makes programs more secure (Because libraries can be updated centrally, and more advanced any improvement in a library will improve the whole program), and smaller.


Linux dynamic libraries have names like `libLIBNAME.so.VERSION` and are located at places like `/lib*/` and `/usr/lib*/`. On Windows, we call them Dynamic Linked Libraries (DLLs).

> Dynamic linking is also called shared libraries because all the programs are sharing one library which is separately installed.


Libraries related to system utilities are installed in `/lib` and `/lib64` (for 32bit and 64bit libraries) and libraries installed by other software will be located at `/usr/lib` and `/usr/lib64`.

### ldd

With this you can find the libraries a certain program uses statically or dynamically.

```shell
 ldd /bin/echo
```

## Symbolic Links for Libraries

If you are writing a program and you use udev functions, you will ask for a library called libudev.so.1. But a Linux distro, might call its version of udev library libudev.so.1.4.0.

I will check the same thing on my system. First I'll find where the libudev.so.1 on my system is:

```shell
locate libudev.so.1
/lib/i386-linux-gnu/libudev.so.1
/usr/lib64/libudev.so.1.7.3
```

And the file contains:

```shell
ls -la /lib/i386-linux-gnu/libudev.so.1
lrwxrwxrwx 1 root root    16 Nov 13 23:05 /lib/i386-linux-gnu/libudev.so.1 -> libudev.so.1.4.0
```

This is symbolic link so no mather the version the linking will work.

## Dynamic Library Configs and Cache

Dynamic linking ic configured using a text config file that is located at `/etc/ld.so.conf`. This just include `/etc/ld.so.conf.d/`

To load these libraries faster, `ldconfig` commands processed all these files. It creates `ld.so.cache` to locate files that should be dynamically loaded.

> If you change the ld.so.conf (or sub-directories) you need to run ldconfig. Try it with -v switch to see the progress / data.


To see ldconfig cache:

```shell
ldconfig -p | head
```

## Finding dynamic libraries

When a program needs a shared library, the system will search files in this order:

1. LD_LIBRARY_PATH environment variable
2. Programs PATH
3. `/etc/ld.so.conf` (Which might load more files from `/etc/ld.so.conf.d/` in its beginning or its end)
4. `/lib/`, `/lib64/`, `/usr/lib/`, `/usr/lib64/`

In some cases we may want to override the default system libraries.

- You are running an old software which needs an old version of a library.
- You are developing a shared library and want to test it without installing it
- You are running a specific program (say from opt) which needs to access its own libraries

In these cases, you can point the environment variable **LD_LIBRARY_PATH** to the library you need to use and then run your program. A colon (:) separated list of directories will tell your program where to search for needed libraries before checking the libraries in `/etc/ld.so.cache`.

Example:


```shell
export  LD_LIBRARY_PATH=/usr/lib/myoldlibs:/home/jadi/lpic/libs/
```

## Loading Dynamically

Dynamic Loaders is used to load dynamic libraries needed by an executable. 

Might be called `ld` or `ld-linux`.

```shell
locate ld-linux
/usr/lib64/ld-linux-x86-64.so.2
/usr/share/man/man8/ld-linux.8.gz
/usr/share/man/man8/ld-linux.so.8.gz
```

```shell
/usr/lib64/ld-linux-x86-64.so.2 /usr/bin/ls
Desktop  Documents  Downloads  Music  Pictures  Public  Templates  tmp  Videos
```

You can run any Linux executable even if its executable bit is not set! just run it using ld-linux as we did a few lines above!
