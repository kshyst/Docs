# Introduction

**Ansible** is an open-source automation platform used for IT tasks such as configuration management, application deployment, and intraservice orchestration. It is designed to be minimal in nature, consistent, secure, and highly reliable.

## Key Concepts

- **Agentless**: Ansible uses **SSH** (for Linux/Unix) or **WinRM** (for Windows) to connect to nodes. There is no need to install any agent software on the managed systems.
- **YAML**: Ansible configurations are written in **YAML** (Yet Another Markup Language), which is easy for humans to read and write.
- **Idempotency**: This is a core principle of Ansible. An **Idempotent** operation is one that has no additional effect if it is called more than once with the same input parameters. In Ansible, this means it only performs changes if the current state of the system does not match the desired state.

## Terminology

- **Control Node**: The machine where Ansible is installed. You run Ansible commands and playbooks from this node.
- **Managed Nodes (Hosts)**: The remote systems or "hosts" that Ansible manages.
- **Inventory**: A file (usually INI or YAML) that lists the IP addresses or hostnames of the **Managed Nodes**.
- **Modules**: Small, standalone scripts that Ansible pushes to the managed nodes to perform specific tasks (e.g., `apt`, `yum`, `copy`, `service`).
- **Tasks**: The smallest unit of action in Ansible. A task calls a specific **Module** with certain arguments.
- **Playbooks**: YAML files containing a ordered list of **Tasks** to be executed on specific hosts. This represents your "Infrastructure as Code."

!!! note
    Ansible ensures that the system reaches the desired state defined in your playbooks, regardless of its starting point.
