# Writing a Dockerfile

A `Dockerfile` is a text document that contains all the commands a user could call on the command line to assemble an image.

## Basic Directives

- `FROM`: Specifies the base image to build from (e.g., `FROM ubuntu:latest`). It defines the base layer.
- `RUN`: Executes commands in a new layer on top of the current image and commits the results.
- `COPY`: Copies files or directories from the host machine into the container's file system.
- `CMD`: Provides defaults for an executing container. There can only be one `CMD` instruction in a `Dockerfile`.

### Example Dockerfile

```dockerfile
FROM ubuntu:latest

# Update and install dependencies
RUN apt-get update && apt-get install -y python3

# Set the working directory
WORKDIR /app

# Copy the application code
COPY . /app

# Start the application
CMD ["python3", "/app/app.py"]
```

### Building the Image
```bash
# Build an image from the current directory
docker build -t my-python-app .

# -t: Specifies a name and optionally a tag in the 'name:tag' format
# -f: Specifies the path to the Dockerfile if it's not named 'Dockerfile' or is in a different directory
```

## Dockerfile Instructions

### EXPOSE
The `EXPOSE` instruction informs Docker that the container listens on the specified network ports at runtime. It serves primarily as documentation between the person who builds the image and the person who runs the container.
```dockerfile
EXPOSE 8080
```

### ENTRYPOINT vs CMD
`ENTRYPOINT` allows you to configure a container that will run as an executable. `CMD` provides default arguments for the `ENTRYPOINT`.

```dockerfile
ENTRYPOINT ["python", "manage.py", "runserver"]
CMD ["0.0.0.0:8080"]
```

### VOLUME
Creates a mount point with the specified name and marks it as holding externally mounted volumes from native host or other containers.
```dockerfile
VOLUME ["/data"]
```

### USER
Sets the user name (or UID) and optionally the user group (or GID) to use when running the image and for any `RUN`, `CMD` and `ENTRYPOINT` instructions that follow it.
```dockerfile
USER nobody
```

### STOPSIGNAL
Sets the system call signal that will be sent to the container to exit.
```dockerfile
STOPSIGNAL SIGTERM
```

### SHELL
Allows the default shell used for the shell form of commands to be overridden.
```dockerfile
SHELL ["/bin/bash", "-c"]
```

### ONBUILD
Adds a trigger instruction to the image to be executed when the image is used as the base for another build.
```dockerfile
ONBUILD RUN echo "This will run when this image is used as a base."
```

### LABEL
Adds metadata to an image.
```dockerfile
LABEL version="1.0" description="My custom image"
```

### ARG
Defines a variable that users can pass at build-time to the builder with the `docker build` command using the `--build-arg <varname>=<value>` flag.
```dockerfile
ARG VERSION=20.04
FROM ubuntu:${VERSION}

ARG APP_HOME=/app
WORKDIR $APP_HOME
```

## Multi-stage Builds
Multi-stage builds are useful for optimizing Dockerfiles while keeping them easy to read and maintain. They allow you to use multiple `FROM` statements in a single Dockerfile.

```dockerfile
# Stage 1: Build
FROM node:16 AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2: Production
FROM node:16-slim AS production
WORKDIR /app
ENV NODE_ENV=production
ENV PORT=3000
COPY --from=build /app/package*.json ./
RUN npm install --only=production
COPY --from=build /app/dist /app/dist

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD curl -f http://localhost:3000/health || exit 1

CMD ["node", "dist/app.js"]
```

## Multi-platform Builds
Using `docker buildx`, you can build your application for multiple architectures (e.g., AMD64, ARM64).

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t my-app .
```
> **Note**: If building for a different architecture than the host, Docker uses `QEMU` for simulation.

## Best Practices

- **Use Build Cache**: Order your instructions from least frequent to most frequent changes to leverage Docker's layer caching.
- **Pin Base Images**: Use specific versions or digests (SHA-256) instead of `latest` to ensure reproducibility.
- **Use .dockerignore**: Exclude unnecessary files (like `node_modules` or `.git`) to reduce image size and build time.
- **Minimize Layers**: Combine related `RUN` commands using `&&` and use multi-stage builds to keep the final image small.
- **Enable Pipefail**: Use `set -o pipefail` in shell scripts to ensure that pipe failures are caught.
  ```dockerfile
  RUN set -o pipefail && wget -O - https://example.com | wc -l > /count
  ```