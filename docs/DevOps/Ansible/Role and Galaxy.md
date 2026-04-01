# Ansible Roles

a Role is a standardized, pre-defined directory structure for your Ansible content. Instead of one giant playbook file, you break down tasks, variables, templates, and handlers into their own separate files within this structure.

## Creating Roles

When you create a role, Ansible expects files to be in specific directories. The main ones you’ll use are:
```text
my_role/
├── defaults/
│ └── main.yml # Default variables for the role (lowest priority)
├── files/
│ └── # Static files that are copied as-is
├── handlers/
│ └── main.yml # Handlers (triggered by ‘notify’)
├── meta/
│ └── main.yml # Role metadata (author, license, dependencies)
├── tasks/
│ └── main.yml # The main list of tasks the role will execute
├── templates/
│ └── # Jinja2 template files
└── vars/
└── main.yml # Role-specific variables (higher priority than defaults)
```

```shell
ansible-galaxy init nginx_role
```

## Using the Role

Use your created role inside a playbook.yml

```yaml
---
- name: Deploy Nginx Web Server
  hosts: webservers
  become: yes

  roles:
    - nginx_role
```

When you run this playbook, Ansible will automatically look for a directory named `nginx_role` (either in a `nginx_roles/` subdirectory next to your playbook or in a system-wide configured path) and execute the tasks in `nginx_role/tasks/main.yml`.

## Making Roles Reusable

In `nginx_role/defaults/main.yml`:

```yaml
---
    # defaults file for nginx_role
    http_port: 80
```

Creating templates in `nginx_role/templates/nginx.conf.j2`:

```shell
    server {
        listen {{ http_port }};
        server_name _;

        location / {
            root /usr/share/nginx/html;
        }
    }
```

Using templates in tasks `nginx_role/tasks/main.yml`:

```yaml
    ---
    # tasks file for nginx_role
    - name: Install Nginx
      ansible.builtin.apt:
        name: nginx
        state: present

    - name: Copy Nginx config template
      ansible.builtin.template:
        src: nginx.conf.j2
        dest: /etc/nginx/sites-available/default
      notify: Restart Nginx

    - name: Start and enable Nginx service
      ansible.builtin.service:
        name: nginx
        state: started
        enabled: yes
```

Also in handler we should define restarting nginx for when the templates change.

At last calling our role in playbook:

```yaml
    ---
    - name: Deploy Nginx Web Server on a custom port
      hosts: webservers
      become: yes

      vars:
        http_port: 8080 # This overrides the default '80'

      roles:
        - nginx_role
```

# Galaxy

Community hub for roles:

```shell
# Search for roles related to 'firewall'
ansible-galaxy search firewall

# installing. should be in author.rolename format
ansible-galaxy install geerlingguy.nginx
```

The best practice for managing project dependencies is to define them in a `requirements.yml` file. This is crucial for version control and collaboration.

```yaml
---
roles:
  - name: geerlingguy.nginx
    version: 3.2.0  # Pinning the version is highly recommended!
  
  - src: geerlingguy.redis
    version: 1.6.0
```

Installing all roles under `/roles`

```shell
ansible-galaxy install -r requirements.yml -p ./roles
```