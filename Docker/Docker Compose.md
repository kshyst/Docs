# Docker Compose

## Example of docker compose file

```yaml
version: '3.8'
services:
  database:
    image: postgres:latest
    container_name: blog-db
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: admin123
      POSTGRES_DB: blog_db
    volumes:
      - blog-db-data:/var/lib/postgresql/data
    networks:
      - blog-net
    ports:
      - "5432:5432"

  backend:
    image: python:3.10
    container_name: blog-backend
    working_dir: /app
    volumes:
      - ./backend:/app
    environment:
      DB_HOST: database
      DB_USER: admin
      DB_PASSWORD: admin123
      DB_NAME: blog_db
    networks:
      - blog-net
    ports:
      - "5000:5000"
    command: sh -c "pip install -r requirements.txt && python app.py"

  frontend:
    image: nginx:latest
    container_name: blog-frontend
    volumes:
      - ./frontend:/usr/share/nginx/html:ro
    networks:
      - blog-net
    ports:
      - "8080:80"

volumes:
  blog-db-data:

networks:
  blog-net:
```

### Some Explaining

```yml
version: '3.8'

services:
  web:
    image: nginx
    ports:
      - "8080:80"
    environment:
      NODE_ENV: production
  database:
    image: postgres
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: secret
    volumes:
      - db_data:/var/lib/postgresql/data
```

- `version`: the version of docker compose we are using
- `services`: an object that defines various services we are using
- web and database are service names.
- image defines image, ports define list of port bindings, environment defines env variables
- volumes define where should the volumes be stored
- `restart:` always, no, on-failure, unless-stopped. First one always restarts when container is stopped and on-failure only restarts when container failed.
- `logging:` 
```    
logging:
  driver: "json-file"
  options:
  max-size: "10m"
```
- `healthcheck:` Lets you define a command for health checking on certain interval.
```yaml
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost/health"]
      interval: 30s
      retries: 3
```
- `build:` Builds the specified image with the given dockerfile, context should be the path to where the dockerfile is and dockerfile is the name of the dockerfile, args are the arguments for building
```yaml
    build:
      context: ./path/to/context
      dockerfile: Dockerfile.custom
      args:
        - APP_VERSION=1.0.0
```
- `external_links:` Lets you connect to containers outside of the compose, useful when you wanna use containers on different machine
```yaml
    external_links:
      - "some_external_container"
```
- `platform:` Specify platform
```yml
    platform: linux/amd64 
```
- `command:` Runs default command
```yaml
    command: ["nginx", "-g", "daemon off;"] 
```
- `tmpfs:` Temp storage defining
```yaml
    tmpfs:
      - /tmp 
```
- `secrets:` Defining secrets
```yml
services:
  web:
    secrets:
      - my_secret
secrets:
  my_secret:
    file: ./my_secret.txt
```
- `ulimits:` Define limits on resources, nproc is the number of processing units.
```yml
    ulimits:
      nproc: 1024
```
- `cpus:` Define number of processing units for the container
```yml
  cpus: "0.5"
```
- `mem_limit:` Limit memory
```yaml
    mem_limit: 512m 
```
- `sysctls:` Kernel parameters
```yml
    sysctls:
      net.ipv4.ip_forward: 1
```
- `pid:` Set specific pid for the container in the host machine
```yaml
    pid: host 
```

## Commands

### up

```shell
docker compose up
```