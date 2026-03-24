# Ansible

- For automation on configuring and installing software on servers
- Uses SSH or WinRM to do it so no need to install anything on the servers
- Uses YAML
- Idempotent. Means it only do changes if the desired state is not met.

## Terminology

- **Control Node**: The machine where ansible is installed
- **Managed Nodes (Hosts)**: The servers we tend to manage
- **Inventory**: A file that lists IP addresses or hostnames of the managed nodes
- **Modules**: Small standalone apps. Ansible pushes these to servers to the actual work like `apt`, `cp`, etc.
- **Tasks**: Actions we want to perform which will call specific modules
- **Playbooks**: YAML files containing a list of Tasks to be executed on certain Hosts. This is your “Infrastructure as Code.”

## How it Works

### Step 1: The Inventory file

Ansible needs to know what machines to talk to. By default, Ansible looks at `/etc/ansible/hosts`, but it is best practice to create your own inventory file for each project.

Create custom inventory file: `inventory.ini`:

```yml
# This is an INI formatted inventory file

# A single host not in any specific group
192.168.1.10

# A group of web servers
[webservers]
web1.example.com
web2.example.com
192.168.1.20

# A group of database servers
[dbservers]
db1.example.com
```

### Step 2: Ad-Hoc Commands

Quick commands instead of writing a whole playbook.

```shell
ansible [pattern] -i [inventory_file] -m [module] -a "[module_arguments]"
```

Example: Ping Test

The ping module doesn’t use the standard ICMP ping; it attempts to log in via SSH and verify that Python is available on the target.

```shell
ansible all -i inventory.ini -m ping
```

### Step 3: Playbooks

`setup-web.yml`:

```yaml
---
- name: Setup Web Servers
  hosts: webservers
  become: yes  # This tells Ansible to use sudo for these tasks

  tasks:
    - name: Ensure Apache is installed (Debian/Ubuntu)
      apt:
        name: apache2
        state: present
        update_cache: yes

    - name: Ensure Apache is running and enabled on boot
      service:
        name: apache2
        state: started
        enabled: yes

    - name: Copy index.html to the web server
      copy:
        content: "<h1>Hello from Ansible!</h1>"
        dest: /var/www/html/index.html
```

- ---: Indicates the start of a YAML file.
- name: A human-readable description of what the playbook or task is doing.
- hosts: Matches the group [webservers] in your inventory.ini.
- become: yes: Elevates privileges to root.
- tasks: The list of actions to perform.
- apt, service, copy: These are Ansible Modules.

To execute the playbook:

```shell
ansible-playbook -i inventory.ini setup-web.yml
```

If you run the exact same command a second time, Ansible will report ok instead of changed because of idempotency

## Step 4: Dynamic Variables

Ansible allows you to make your playbooks dynamic using variables. Variables in Ansible are referenced using double curly braces: `{{ variable_name }}`

```yaml
---
- name: Setup Web Servers with Variables
  hosts: webservers
  become: yes
  
  # Define variables for this specific play
  vars:
    web_package: apache2
    welcome_message: "Welcome to my Automated Server!"

  tasks:
    - name: Install the web package
      apt:
        name: "{{ web_package }}"
        state: present

    - name: Deploy custom index file
      copy:
        content: "<h1>{{ welcome_message }}</h1>"
        dest: /var/www/html/index.html
```

## Step Five: Gathering Facts

When a playbook runs, the very first thing you usually see in the output is:

```shell
TASK [Gathering Facts]
```

Ansible automatically connects to the managed nodes and runs a hidden module called setup. This gathers hundreds of variables about the target system (OS, IP addresses, disk space, memory).

To use these facts:

```yaml
    - name: Show the operating system
      debug:
        msg: "This server is running {{ ansible_distribution }} version {{ ansible_distribution_version }}"
```

