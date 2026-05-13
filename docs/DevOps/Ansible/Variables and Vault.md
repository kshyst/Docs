# Variables and Vault

**Variables** in Ansible allow you to make your playbooks dynamic and reusable. **Ansible Vault** provides a way to encrypt sensitive data like passwords and API keys.

## Using Variables

Variables are referenced using Jinja2 syntax: `{{ variable_name }}`.

```yaml
vars:
  web_package: apache2

tasks:
  - name: Install package
    apt:
      name: "{{ web_package }}"
      state: present
```

### Variable Priority
When the same variable is defined in multiple places, Ansible follows a specific order of precedence (from lowest to highest):

1.  Role defaults (`role/defaults/main.yml`)
2.  Inventory variables
3.  Playbook variables
4.  Role variables (`role/vars/main.yml`)
5.  Command-line extra vars (`-e "var=value"`) — **Always wins**.

### Capturing Output (Register)
The `register` keyword allows you to capture the output of a task and store it in a variable for later use.

```yaml
- name: Check if a file exists
  ansible.builtin.stat:
    path: /etc/myapp.conf
  register: file_check

- name: Do something if the file exists
  debug:
    msg: "The file exists!"
  when: file_check.stat.exists
```

## Ansible Vault

**Ansible Vault** encrypts files or specific variables so they can be safely stored in version control.

### Common Commands

- **Create**: `ansible-vault create secrets.yml`
- **Edit**: `ansible-vault edit secrets.yml`
- **Encrypt existing**: `ansible-vault encrypt my_passwords.yml`
- **View**: `ansible-vault view secrets.yml`
- **Encrypt string**: `ansible-vault encrypt_string 'SuperSecret!' --name 'db_pass'`

### Using Vault in Playbooks

When running a playbook that requires a vault password, use the following flags:

```bash
# Prompt for password
ansible-playbook deploy.yml --ask-vault-pass

# Read password from a file
ansible-playbook deploy.yml --vault-password-file ~/.vault_pass.txt
```

!!! tip
    Group variables (`group_vars/`) and host variables (`host_vars/`) are excellent places to store environment-specific configurations and secrets.
