# Ad-Hoc Commands

**Ad-hoc commands** are quick, one-line commands used to perform tasks without writing a full playbook. They are ideal for simple operations like checking connectivity, restarting services, or managing packages.

## Basic Syntax

```bash
ansible [pattern] -m [module] -a "[arguments]" [options]
```

- **pattern**: The hosts or groups to target (e.g., `all`, `webservers`).
- **-m**: The **Module** to use (e.g., `ping`, `shell`, `apt`).
- **-a**: The arguments required by the module.
- **-i**: Path to the **Inventory** file.

## Common Options

- `-b`, `--become`: Run operations with privilege escalation (usually **sudo**).
- `-K`, `--ask-become-pass`: Prompt for the privilege escalation password.
- `-u`, `--user`: Connect as a specific user.
- `-f`, `--forks`: Number of parallel processes to use (default is 5).
- `-B`, `--background`: Run the task in the background (asynchronous).
- `-P`, `--poll`: Poll the background job for status.
- `--limit`: Restrict the targeted hosts to a specific subset.

## Practical Examples

### Connectivity Test
The `ping` module verifies SSH connectivity and the availability of Python on the target.
```bash
ansible all -i inventory.ini -m ping
```

### Service Management
Restarting Nginx on all webservers:
```bash
ansible webservers -b -m service -a "name=nginx state=restarted"
```

### Package Installation
Installing `htop` using 10 parallel forks:
```bash
ansible db_servers -f 10 -b -m apt -a "name=htop state=present"
```

### File and Directory Management
Ensuring a directory exists with specific permissions:
```bash
ansible app_servers -b -m file -a "path=/opt/myapp state=directory mode=0755 owner=deploy" -K
```

### Gathering Facts
Extracting specific system information:
```bash
# Get all IP addresses on eth0 across all servers
ansible all -m setup -a "filter=ansible_eth0" -i inventory.ini
```

### Asynchronous Tasks
Running a long-running script in the background for 3600 seconds without polling:
```bash
ansible all -B 3600 -P 0 -m script -a "long_running_script.sh"
```

!!! warning
    While ad-hoc commands are powerful for quick fixes, they are not version-controlled or easily reproducible. For complex configurations, always prefer **Playbooks**.
