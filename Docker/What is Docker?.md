# What is Docker

Docker doesn't need to run a whole OS for each container.

![dockervm](img/docker_vm.png)

We need to create container from docker images.

## What is Container

- Images: Contain codes and dependencies to run an application
- Containers: Each running version of an image.
- Docker Engine: Engine for handling and controlling containers.
- Docker Hub: Online service for getting images.

## How docker works

Containers contact kernel directly so each container can share processing and storage resources with other with other containers and host system.

Containers are isolated logically which makes them not effecting other containers

- `Docker Daemon` interacts with docker client through restful APIs
- `containerd` Is a runtime that manages the life cycle of a container
- `Shim API` Bridge between containerd and runtime
- `OCI Runtime` Controls namespaces, Cgroups(Resource Control Groups)

Containers use kernel features like namespacing and cgroups make resources isolated and manages them.

Containers are just a process running on Host. But is isolated from other process and other isolations like file resources and network resources.