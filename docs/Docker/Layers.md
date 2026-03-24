# Docker layers

Docker layers are changes in images from the previous layer. Every command in docker file creates a layer.

Each layer contains:

- Added or deleted files 
- Environmental variables
- Information about the commands ran

Docker uses Copy-on-Write COW for creating layers. This will make the changes on the previous layer without copying it.

Layer types:

- `Base Layers` like ubuntu, python and ... which specifies the os and environment of the container.
- `Build Layers` these are built from `RUN` commands and are mostly for installing packages and changing the system files.
- `Copy Layers` built from `ADD` and `COPY` commands
- `Environmental Layers` built from `ENV` command. The envs specified will only be available to the layers after this.
- `Command layers` built by `CMD` and `ENTRYPOINT`

## How layers work

Each layer is built and stored separately and when something changes, only the corresponding layer will change and others are used from cache.  