# Jenkins Installation

Jenkins can be installed in several ways. This guide covers the most common methods for getting started.

## Running with Docker (Recommended)

Running Jenkins in a Docker container is the fastest and most isolated way to test and evaluate Jenkins.

### Prerequisites
- Docker installed on your machine.

### Execution
Run the following command to start a Jenkins container:

```bash
docker run -p 8080:8080 -p 50000:50000 \
  -v jenkins_home:/var/jenkins_home \
  --name jenkins-server \
  jenkins/jenkins:lts
```

- `-p 8080:8080`: Maps the Jenkins UI port.
- `-p 50000:50000`: Maps the agent communication port.
- `-v jenkins_home:/var/jenkins_home`: Creates a persistent volume for your data.

## Standalone WAR File

For environments where Docker is not available, you can run Jenkins as a standalone Java application.

### Prerequisites
- Java Runtime Environment (JRE) or JDK 11 or 17.

### Execution
1. Download the latest Jenkins WAR file from [jenkins.io](https://www.jenkins.io/download/).
2. Open a terminal and run:

```bash
java -jar jenkins.war --httpPort=8080
```

Jenkins will start and unpack its files into `~/.jenkins/` by default.

## Setup Wizard

After starting Jenkins for the first time, navigate to `http://localhost:8080`.

### Initial Admin Password
To unlock Jenkins, you need the initial administrator password.
- **Docker**: Check the container logs using `docker logs jenkins-server`.
- **WAR**: The password will be printed to the console and also stored in:
  `/var/jenkins_home/secrets/initialAdminPassword` (or `~/.jenkins/secrets/initialAdminPassword`).

### Suggested Plugins
On the **Customize Jenkins** page, it is highly recommended to select **Install suggested plugins**. This will install common plugins for Git, Pipelines, and SSH, which are essential for most users.

!!! warning
    Ensure you create a strong administrator user during the setup process and do not rely on the initial admin password for long-term use.
