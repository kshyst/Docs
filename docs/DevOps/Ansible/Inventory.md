# Inventory

The **Inventory** file defines the hosts and groups of hosts upon which commands, modules, and playbooks operate. While Ansible looks at `/etc/ansible/hosts` by default, it is best practice to maintain a project-specific inventory file.

## Inventory Structure (INI Format)

The **INI** format is a common way to structure inventory files. It allows you to group hosts logically.

Example `inventory.ini`:

```ini
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

## Groups and Hierarchies

Groups allow you to target multiple hosts at once. You can also create groups of groups using the `:children` suffix.

```ini
[atlanta]
host1
host2

[raleigh]
host3
host4

[southeast:children]
atlanta
raleigh
```

!!! tip
    Using a custom inventory file per project improves portability and prevents accidental changes to global system configurations. Use the `-i` flag to specify your inventory file when running Ansible commands.
