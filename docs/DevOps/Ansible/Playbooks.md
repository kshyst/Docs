# Playbooks

**Playbooks** are the building blocks of Ansible. They are YAML files that describe the desired state of your systems. Unlike ad-hoc commands, playbooks are version-controlled and can manage complex multi-tier deployments.

## Playbook Structure

A playbook consists of one or more **plays**. A play maps a group of hosts to a set of tasks.

```yaml
---
- name: Setup Web Servers
  hosts: webservers
  become: yes  # Elevate privileges to root

  tasks:
    - name: Ensure Apache is installed
      apt:
        name: apache2
        state: present

    - name: Ensure Apache is running
      service:
        name: apache2
        state: started
```

### Key Components

- **name**: A description of the play or task.
- **hosts**: Specifies the managed nodes from the **Inventory**.
- **become**: Enables privilege escalation.
- **tasks**: A list of modules to execute.
- **handlers**: Special tasks that only run when notified by another task (e.g., restarting a service after a config change).

## Advanced Features

### Loops
Loops allow you to perform a task multiple times with different values.

```yaml
- name: Configure Pip Mirror
  command: "pip config set global.{{ item.key }} {{ item.value }}"
  loop:
    - { key: 'index-url', value: 'https://mirror-pypi.runflare.com/simple' }
    - { key: 'trusted-host', value: 'mirror-pypi.runflare.com' }
```

### Gathering Facts
By default, Ansible runs the `setup` module at the start of every play to gather **Facts** about the managed nodes (OS, IP, CPU, etc.).

```yaml
- name: Show the operating system
  debug:
    msg: "This server is running {{ ansible_distribution }} {{ ansible_distribution_version }}"
```

## Practical Example: FastAPI Deployment

This example demonstrates a comprehensive deployment strategy including dependencies, directories, and systemd services.

```yaml
---
- name: Deploy FastAPI Application
  hosts: api_servers
  become: yes

  vars:
    app_dir: /opt/fastapi_app
    app_port: 8000

  tasks:
    - name: Install system dependencies
      apt:
        name: [python3, python3-pip, python3-venv]
        state: present
        update_cache: yes

    - name: Ensure application directory exists
      file:
        path: "{{ app_dir }}"
        state: directory
        mode: '0755'

    - name: Copy application files
      copy:
        src: files/main.py
        dest: "{{ app_dir }}/main.py"

    - name: Install Python dependencies
      pip:
        requirements: "{{ app_dir }}/requirements.txt"
        virtualenv: "{{ app_dir }}/venv"

    - name: Configure systemd service
      template:
        src: templates/fastapi.service.j2
        dest: /etc/systemd/system/fastapi.service
      notify: Reload systemd

  handlers:
    - name: Reload systemd
      systemd:
        daemon_reload: yes
        name: fastapi
        state: restarted
```

## Running Playbooks

Execute a playbook using the `ansible-playbook` command:

```bash
ansible-playbook -i inventory.ini site.yml -K
```

!!! note
    If you run the same playbook twice, Ansible will report `ok` for tasks where the state is already met, demonstrating **Idempotency**.
