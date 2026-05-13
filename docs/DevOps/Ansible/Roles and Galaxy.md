# Roles and Galaxy

**Roles** provide a framework for fully independent and reusable collections of variables, tasks, files, templates, and modules. **Ansible Galaxy** is a community hub for sharing and downloading these roles.

## Role Directory Structure

A role follows a standardized directory structure:

```text
my_role/
├── defaults/ # Default variables (lowest priority)
│   └── main.yml
├── files/    # Static files
├── handlers/ # Handlers
│   └── main.yml
├── meta/     # Metadata (author, dependencies)
│   └── main.yml
├── tasks/    # Main list of tasks
│   └── main.yml
├── templates/ # Jinja2 templates
└── vars/     # Role-specific variables (high priority)
    └── main.yml
```

Initialize a new role using:
```bash
ansible-galaxy init role_name
```

## Using Roles

To use a role, include it in your playbook:

```yaml
---
- name: Deploy Nginx
  hosts: webservers
  roles:
    - role_name
```

## Making Roles Reusable

By using templates and variables, you can make roles adaptable to different environments.

### Defaults (`defaults/main.yml`)
```yaml
---
http_port: 80
```

### Templates (`templates/nginx.conf.j2`)
```jinja2
server {
    listen {{ http_port }};
    ...
}
```

### Overriding Variables
```yaml
- name: Deploy on custom port
  hosts: webservers
  vars:
    http_port: 8080
  roles:
    - nginx_role
```

## Ansible Galaxy

Search for and install community-contributed roles.

```bash
# Search for a role
ansible-galaxy search firewall

# Install a role
ansible-galaxy install geerlingguy.nginx
```

### managing Dependencies (`requirements.yml`)

The best practice is to define project dependencies in a `requirements.yml` file.

```yaml
---
roles:
  - name: geerlingguy.nginx
    version: 3.2.0
  - src: geerlingguy.redis
    version: 1.6.0
```

Install all listed roles:
```bash
ansible-galaxy install -r requirements.yml -p ./roles
```

!!! tip
    Always pin the version of roles in `requirements.yml` to ensure consistent and reproducible deployments across different environments.
