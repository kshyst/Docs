# Playbooks

A playbook template to use:

`deploy-fastapi.yml`:

```yaml
---
- name: Deploy FastAPI Application
  hosts: api_servers
  become: yes  # Run as root

  # Define variables for our project
  vars:
    app_dir: /opt/fastapi_app
    app_port: 8000

  tasks:
    - name: Update apt cache and install system dependencies
      apt:
        name:
          - python3
          - python3-pip
          - python3-venv
        state: present
        update_cache: yes

    - name: Ensure application directory exists
      file:
        path: "{{ app_dir }}"
        state: directory
        mode: '0755'

    - name: Copy FastAPI app (main.py) to server
      copy:
        src: files/main.py
        dest: "{{ app_dir }}/main.py"

    - name: Copy requirements.txt to server
      copy:
        src: files/requirements.txt
        dest: "{{ app_dir }}/requirements.txt"

    - name: Create virtual environment and install dependencies
      pip:
        requirements: "{{ app_dir }}/requirements.txt"
        virtualenv: "{{ app_dir }}/venv"
        virtualenv_command: python3 -m venv

    - name: Template the systemd service file
      template:
        src: templates/fastapi.service.j2
        dest: /etc/systemd/system/fastapi.service

    - name: Reload systemd daemon to recognize new service
      systemd:
        daemon_reload: yes

    - name: Enable and start FastAPI service
      systemd:
        name: fastapi
        state: restarted
        enabled: yes

    - name: Open firewall port for FastAPI (UFW)
      ufw:
        rule: allow
        port: "{{ app_port }}"
        proto: tcp
```

Use `-K` to always be asked for password when become : yes

```shell
ansible-playbook -i inventory.ini deploy-fastapi.yml -K
```

### Writing loops


```yaml
    - name: Configure Pip Mirror
      command: "pip config set global.{{ item.key }} {{ item.value }}"
      loop:
        - { key: 'index-url', value: 'https://mirror-pypi.runflare.com/simple' }
        - { key: 'trusted-host', value: 'mirror-pypi.runflare.com' }
```