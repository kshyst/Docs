# Jenkins Architecture

Jenkins uses a distributed architecture designed to handle high-volume workloads and provide high availability.

## Controller-Agent Model

The Jenkins architecture is based on a **Controller-Agent** (formerly known as Master-Slave) model.

### The Jenkins Controller
The **Controller** is the brain of Jenkins. It is responsible for:
- Serving the HTTP UI and API.
- Scheduling build jobs.
- Managing user configuration and security.
- Managing the plugin environment.
- Orchestrating the execution of builds on agents.

### Jenkins Agents
**Agents** are small Java executables that run on remote machines. Their primary responsibility is to execute the build jobs dispatched by the controller.
- Agents can run on different operating systems (Linux, Windows, macOS).
- They provide the actual computing resources (CPU, RAM, Disk) for the builds.
- They can be scaled horizontally to handle more concurrent builds.

## Workload Distribution

Jenkins achieves scalability by offloading build execution to agents.

1. **Job Trigger**: A job is triggered (manually, via webhook, or on a schedule).
2. **Scheduling**: The Controller determines which agent is available and suitable (based on labels).
3. **Execution**: The Controller dispatches the task to the selected Agent.
4. **Reporting**: The Agent executes the commands and streams the console output back to the Controller in real-time.

## The Workspace

The **Workspace** is a dedicated directory on the agent's file system where the build actually takes place.

- **Source Code**: Jenkins clones the source code into the workspace.
- **Build Artifacts**: Any compiled binaries or reports are generated here.
- **Isolation**: Each job has its own workspace, ensuring that different builds do not interfere with each other's files.

!!! note
    By default, Jenkins does not clean up workspaces after a build. It is a best practice to use the `cleanWs()` step in pipelines to keep the agent's disk space under control.
