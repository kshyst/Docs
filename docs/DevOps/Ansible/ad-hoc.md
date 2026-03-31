# ad-hoc

```shell
ansible [pattern] -m [module] -a "[arguments]" [options]
```

- `-b` or `--become`: Run operations with become (usually sudo).
- `-K` or `--ask-become-pass`: Prompt for the privilege escalation password.
- `-u` or `--user`: Connect as a specific user.

Example of restarting webserver:

```shell
ansible webservers -b -m service -a "name=nginx state=restarted"
```

Ansible runs 5 servers at a time by default. To increase concurrency:

```shell
ansible all -f 20 -m command -a "df -h"
```

If you are running a task that takes a long time (like a system update or a large script), the SSH connection might time out. You can run ad-hoc commands asynchronously.

- `-B <seconds>`: Maximum time the job is allowed to run.
- `-P <seconds>`: How often to poll the job for status. (Set to 0 to dont get anything)

Sometimes you want to target a group, but exclude a specific node or only target one node within that group for a quick test.

- `--limit`: Restrict the targeted hosts.

```shell
ansible webservers --limit web01.example.com -b -m yum -a "name=* state=latest"
```

## Examples:

```shell
# Get all IP addresses on eth0 across all servers
ansible all -m setup -a "filter=ansible_eth0" -i inventory.ini

# Ensure a directory exists with specific permissions
ansible app_servers -b -m file -a "path=/opt/myapp state=directory mode=0755 owner=deploy" -K

# Push a new DNS configuration
ansible all -b -m copy -a "src=/local/path/resolv.conf dest=/etc/resolv.conf owner=root mode=0644"

# Lock a departed employee's account
ansible all -b -m user -a "name=jdoe state=absent remove=yes"

# Install package using 10 forks
ansible db_servers -f 10 -b -m apt -a "name=htop state=present" 
```