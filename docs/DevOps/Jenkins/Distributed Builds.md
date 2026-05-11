# Distributed Builds

As your automation needs grow, a single Jenkins server may not have enough resources. Distributed builds allow you to scale Jenkins by adding more compute nodes.

## Adding Nodes and Agents

To add a new execution node:
1. Navigate to **Manage Jenkins** > **Manage Nodes and Clouds**.
2. Select **New Node**.
3. Provide a name and select **Permanent Agent**.

## Connection Methods

How the controller communicates with the agent depends on the environment.

### SSH (Linux/Unix)
This is the most common method for Linux agents. The controller initiates an SSH connection to the agent and starts the agent process.
- **Requirement**: The controller must have SSH access to the agent.
- **Configuration**: Provide the agent's IP/hostname and a credential (SSH key or username/password).

### Inbound Agent (JNLP)
Commonly used for Windows agents or agents behind a firewall. The agent initiates the connection to the controller.
- **Requirement**: The agent machine must be able to reach the controller's TCP port (default 50000).
- **Configuration**: The agent runs a small Java command provided by the Jenkins UI.

## Using Labels

**Labels** are the primary way to control where a job executes.

- **Purpose**: Group agents by their capabilities (e.g., `linux`, `windows`, `docker`, `high-memory`).
- **Pipeline Usage**:
  ```groovy
  pipeline {
      agent { label 'docker' }
      stages {
          stage('Build') {
              steps {
                  sh 'docker build -t my-app .'
              }
          }
      }
  }
  ```

!!! note
    By using labels, you ensure that builds requiring specific tools (like a specific version of Java or a GPU) always land on the correct hardware.
