# Writing Dockerfile

- `FROM`: Tells docker to get certain image from registry and also we can specify the version after `:`. It defines the base layer(or base image)
- `RUN`: Is for running bash commands.
- `COPY`: Copies all the file next to the Dockerfile to the specified directory in our new container environment.
- `CMD`: Tells docker what to run to start the project.

```Dockerfile
FROM ubuntu:latest

RUN apt-get update && apt-get install -y python3

COPY . /app

CMD ["python3", "/app/app.py"]
```

### Creating the image

```shell
docker build .
```

- `-t` specifies the name of the image
- `-f` giving the exact dockerfile name and location

### Running the Container

```shell
docker run 
```

## Dockerfile Commands

### Expose

This expose command in dockerfile does nothing. It only has documentational purposes

```Dockerfile
EXPOSE 8080
```

### Entrypoint and CMD

ENTRYPOINT is what should run after we run our container and CMD is some additional arguments added to it.

```Dockerfile
ENTRYPOINT ["python", "manage.py", "runserver"]
CMD ["-h", "0.0.0.0", "-p", "8080"]
```

The default ENTRYPOINT for docker is :

```shell
/bin/sh -c [COMMAND]
```

Which runs the given command after it

Basically if we don't give ENTRYPOINT, it will to the `/bin/sh -c` followed by whatever CMD is in our Dockerfile

### Volume
Is used to create and connect volumes
```dockerfile
VOLUME ["/data"]
```

### User

Specifies the user inside container
```dockerfile
USER nobody
RUN whoami
```

This will print out nobody

### Stop Signal

Specifies the signal that should be used to stop the container

```dockerfile
STOPSIGNAL SIGTERM
```

### Shell

Specifies the shell for container

SHELL ["/bin/bash", "-c"]
RUN echo "Hello from bash"

### On Build

For specifying commands that should only run when we are using this image as a base image.

```dockerfile
ONBUILD RUN echo "This will run when the image is used as a base."
```

### Label

For giving information about the author

```dockerfile
LABEL version="1.0" description="My custom image"
```

### Arg

For setting arguments in the build time.

```dockerfile
ARG VERSION=20.04
FROM registry.gitlab.com/qio/standard/ubuntu:${VERSION}

ARG APP_HOME=/app
WORKDIR $APP_HOME
```

Usage in docker commands is like:

```shell
docker build --build-arg VERSION=22.04 -t myimage .

```

## Using Multi-stage Build

```dockerfile
FROM node:16 AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

RUN npm run build

FROM node:16-slim AS production

WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000

COPY --from=build /app/package.json /app/package-lock.json ./
RUN npm install --only=production

COPY --from=build /app/dist /app/dist

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD curl -f http://localhost:3000/health || exit 1

CMD ["node", "dist/app.js"]
```

In the first part it uses `node:16` for installing requirements and building the application and the `AS build` part shows the stage name.

The second stage we use `node:16-slim` which only has the requirements for running the application and is pretty lighter.

The `ENV` will set environmental variable inside the container.

Then it will copy the lock files from previous stage and then only installs the requirements needed for production. Then it copies the built application to the new stage.

The `HEALTHCHECK` command will run the CMD after it to check if the container works correctly or nay

## Multi-platform build

Sometimes you want to build an application on multiple platforms or environments

Using `docker buildx` you can build your applicaion on multiple platforms

```shell
docker buildx build --platform linux/amd64,linux/arm64,windows/amd64 -t my-app .
```

> If you want to build on arm but your device is on amd, docker will use `QEMU` to simulate arm architecture.

# Best Practices On Writing Dockerfile

### Use cache build

Avoid doing unnecessary changes and docker will use cache to build

### Pinning base images

Using exact version and Digest instead of latest

A Docker image digest is a unique, cryptographic identifier (SHA-256 hash) representing the content of a Docker image. Unlike tags, which can be reused or changed, a digest is immutable and ensures that the exact same image is pulled every time. This guarantees consistency across different environments and deployments.

