# Variable and Vault

## Variable Priority

From lowest to highest priority:

- Role defaults (`role/defaults/main.yml`)
- Inventory variables
- Playbook variables
- Role variables (`role/vars/main.yml`)
- Command-line extra vars (-e "my_var=value") - This always wins.

## Register

Capturing output of a task and storing it in a variable:

```yaml
- name: Check if a file exists
  ansible.builtin.stat:
    path: /etc/myapp.conf
  register: file_check

- name: Do something if the file exists
  ansible.builtin.debug:
    msg: "The file exists!"
  when: file_check.stat.exists == True
```

## Group Variables

```yaml
# In your vars file:
db_config:
  port: 5432
  user: admin
  max_connections: 100

# In your task/template:
# You access it via dot notation: {{ db_config.port }}
```

## Vault

Ansible Vault encrypts files or specific variables so you can safely store them in version control.

```shell
# Create vault
ansible-vault create group_vars/all/secrets.yml

# Edit vault
ansible-vault edit group_vars/all/secrets.yml

# Encrypt an existing plain text file
ansible-vault encrypt my_passwords.yml

# View vault
ansible-vault view group_vars/all/secrets.yml

# Just one variable
ansible-vault encrypt_string 'SuperSecretPassword!' --name 'db_password'
```

```yaml
# vars/main.yml
db_user: admin
db_password: !vault |
          $ANSIBLE_VAULT;1.1;AES256
          32326566373766646533366366623661643431616239323330386638316132636130316361313733
          ... (rest of the encrypted string)
```

When using vault, deploying playbook asks its pass:

```shell
ansible-playbook deploy.yml --ask-vault-pass
```

To read the vault pass from  a file:
```shell
ansible-playbook deploy.yml --vault-password-file ~/.vault_pass.txt
```

## Extra

You can do simple math in j2:

```yaml
 {{ db_config.max_connections / 2 }}
```