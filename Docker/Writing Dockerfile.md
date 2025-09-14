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

### Running the Container

```shell
docker run 
```