# Vagrant
Vagrant is an open-source tool that helps us to automate the creation and management of Virtual Machines. In a nutshell, we can specify the configuration of a virtual machine in a simple configuration file, and Vagrant creates the same Virtual machine using just one simple command. It provides command-line interfaces to automate such tasks. 
### What is a Virtual Machine?
Virtual Machine is a machine that does not exist physically but can be used just like a physical computer. Any task that can be done on a physical machine can also be executed in a virtual machine. But Virtual Machine is built on top of a physical system, and multiple virtual machines can be created in a single physical computer. All the virtual machines share the same hardware, but each of them might have a separate operating system. The physical system that hosts all the virtual machines is called the Host Computer. The medium that separates the Host Computer hardware and the virtual environments is something called Hypervisor, or Hyper-V.

![V](https://media.geeksforgeeks.org/wp-content/uploads/20210522190257/VM2.png)

Each Virtual Machine should have its own configuration like operating system, CPUs, RAM, Hard Disk Memory, networking, etc. And the creation of such VMs, manually configuring all the properties is really a hectic task. In this scenario, Vagrant comes into the picture. 
## Terminologies Of Vagrant
- **Vagrant Box**: The basic unit of Vagrant setup is Vagrant Box. Just like Docker Image, Vagrant Box is a self-contained image of the Operating System. More specifically, it is a packaged Virtual Machine. Instead of installing an operating system and all the software components inside a VM manually, 
- **Vagrantfile**: Vagrantfile is a configuration file that contains all the information about the Vagrant Box. It contains the configuration of the VM, like the operating system, CPU, RAM, Hard Disk Memory, networking, etc.

## Installing Vagrant/Packer on Ubuntu/Debian
### Add the HashiCorp GPG key.
```bash
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
```
### Add the official HashiCorp Linux repository.
```bash
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
```
### Update and install.
```bash
sudo apt-get update
sudo apt-get install vagrant
sudo apt-get install packer
```

### Download and Install VirutalBox
[Virtual Box for Linux](https://www.virtualbox.org/wiki/Linux_Downloads)
# Add public box in vagrant
### Using Public Boxes
### Adding a bento box to Vagrant
```bash
vagrant box add --provider virtualbox bento/ubuntu-22.04
vagrant box add --provider virtualbox bento/debian-12
```

### Create a Vagrantfile
```bash
vagrant init bento/ubuntu-22.04
```

### Using a bento box in a Vagrantfile
```bash
Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-20.04"
end
```

```bash
Vagrant.configure("2") do |config|
  config.vm.box = "bento/debian-12"
end
```

# Building Boxes
### Requirements: install packer, vagrant and virtualbox

### clone bento project
```bash
git clone https://github.com/chef/bento.git
```
### To build an Ubuntu 18.04 box for only the VirtualBox provider
```bash
cd packer_templates/ubuntu
packer build -only=virtualbox-iso ubuntu-22.04-amd64.json
```
