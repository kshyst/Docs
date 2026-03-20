# Linux as a Virtualization Guest

## Terms
- Virtual machine
- Linux container
- Application container
- Guest drivers
- SSH host keys
- D-Bus machine id

---

## Hypervisor

To check and see if your host operating system or CPU, supports using hypervisors check for the `vmx` (for Intel CPUs) or `svm` (for AMD CPUs) in your `/proc/cpuinfo` in flags.

> You may need to turn the hypervisor option On using your BIOS or UEFI.

Based on your CPU you should have kvm or kvm-amd kernel modules loaded.

```shell
lsmod | grep -i kvm
sudo modprobe kvm
```

> If you see hypervisor in your /proc/cpuinfo it means that you are inside a virtualized Linux machine

---

![](img/hypervisor.png)

## Type 2 Hypervisors

These hypervisors run on a conventional operating system (OS) just as other computer programs do. A guest operating system runs as a process on the host. Type-2 hypervisors abstract guest operating systems from the host operating system.

In other words, a type 2 hypervisor is the software between the guest and host. It completely runs on the host OS and provides virtualization to the guest.

Two of the most famous Type 2 hypervisors are `VirtualBox (from Oracle)` and `VMware`.

## Type 1 Hypervisors

hese hypervisors run directly on the host's hardware to control the hardware and manage guest operating systems. For this reason, they are sometimes called bare-metal hypervisors. The first hypervisors, which IBM developed in the 1960s, were native hypervisors. These included the test software SIMMON and the CP/CMS operating system, the predecessor of IBM z/VM.

Some of the most famous Type 1 hypervisors are KVM, Xen & Hyper-V. KVM is built-in since Linux Kernel version 2.6.20.

---

## Guest-Specific Configs

 If we are cloning a machine or sometimes creating them from templates, at least we need to change these on each machine before booting them:

- Host Name
- NIC MAC Address
- NIC IP (If not using DHCP)
- Machine ID (delete the /etc/machine-id and /var/lib/dbus/machine-id and run dbus-uuidgen --ensure. These two files might be soft links to each other)
- Encryption Keys like SSH Fingerprints and PGP keys
- HDD UUIDs
- Any other UUIDs on the system

---

## IaaS

As the name implies, Infrastructure as a Service or IaaS means offloading parts of your infrastructure to other companies. This means buying services like electricity, cooling, and even running hypervisors to another company and just renting your VirtualMachine from them. This makes your life easier because things like "Adding New Hards" now only mean paying a bit more for more HDD on your machine; Instead of buying HDDs and installing them, ... This is called cloud! You might even be able to move your machine from one continent to another one just with one click.

Like AWS, Azure, Google Cloud

There are programs like `cloud-init` which help you initialize your cloud machine with ease.