# Localization

## Time

Commands for seeing times:

- `date`
- `cal`
- `timedatectl`

```shell
date +'%Y %m %H  %M'

# Shows time settings
timedatectl

# Shows speed of command
time ls
```

Most distros ask about time configs in installation. To change them
later use these:

```shell
# Older
tzconfig

# Newer
tzselect
```

To customize shells timezone add this to `.profile`:

```shell
TZ='Indian/Maldives'; export TZ
```

When setting a timezone config, what actually happens is `/etc/localtime` which is a symbolic link will be linked to your selected timezone.

Timezones are in `/usr/share/zoneinfo/*` for example `/usr/share/zoneinfo/Africa/Lagos`.

So if you want to change timezone do: 

```shell
ln -s /usr/share/zoneinfo/Africa/Lagos /etc/localtime
```

Delete the previous link before creating the new one.

We also have `/etc/localtime`. This file is used to describe timezone info to the
system. But `/etc/timezone` is a text file used by programs to know your zone.

### Hardware Clock

How does system keep showing right clock after power off?

With these packages installed:

```shell
sudo apt install util-linux util-linux-extra
```

Run:

```shell
hwclock
```

It actually shows the localtime after reading timezone so `hwclock` is in UTC.

> Windows change the hardware clock to localtime but linux use UTC so when you 
> dual boot, changing os will cause in non-sense device time.

If you use `time -s` to set time, it will now affect `hwclock`.

### NTP (Network Time Protocol)

NTP (Network Time Protocol) is one of the oldest protocols still in use on the Internet (invented in 1985). Its sole purpose is to synchronize the clocks of computers over a network to within a few milliseconds of Coordinated Universal Time (UTC).

It connects to NTP servers which use the best atomic clocks on the planet. Most famous server
is `pool.ntp.org` which contains many pools.

#### ntpdate

```shell
sudo ntpupdate pool.ntp.org
```

Then to set system clock to the last synced time:

```shell
hwclock -w
```

#### ntpd

Instead of setting manually you can use a daemon service called `ntpd`.

```shell
apt install ntp
systemctl start ntp
```

> You can't use both `ntpdate` and the service.


#### ntpq

To query the ntp service use `ntpq`. It connects to ntp server with a commandline

It can troubleshoot the service

```shell
# Shows connection to pools with details
# If it shows + behind a server name, it is good to use
ntpq -p
```

#### chrony

Newer NTP protocol implementation

Configs in `chrony.conf`

To change time in a commandline style use `chronyc`


## Locale

This command shows locale variables
```shell
locale
```

You can change your system language by setting `LANG` in `.profile`