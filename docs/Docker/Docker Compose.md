# Docker Compose

Docker Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application's services, networks, and volumes.

## Example Docker Compose File

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

## Key Configuration Directives

- `version`: The version of the Docker Compose file format.
- `services`: Defines the different containers that make up your application.
- `image`: Specifies the image to start the container from.
- `ports`: Defines port mappings between the host and the container (`host:container`).
- `environment`: Sets environment variables inside the container.
- `volumes`: Defines mount points for persistent data.
- `restart`: Configures the restart policy (e.g., `always`, `no`, `on-failure`, `unless-stopped`).
- `logging`: Configures logging drivers and options.
  ```yaml
  logging:
    driver: "json-file"
    options:
      max-size: "10m"
  ```
- `healthcheck`: Defines a command to check the health of the service.
  ```yaml
  healthcheck:
    test: ["CMD", "curl", "-f", "http://localhost/health"]
    interval: 30s
    retries: 3
  ```
- `build`: Specifies configuration for building an image from a Dockerfile.
  ```yaml
  build:
    context: ./path/to/context
    dockerfile: Dockerfile.custom
    args:
      - APP_VERSION=1.0.0
  ```
- `external_links`: Connects to containers outside of the current Compose project.
- `platform`: Specifies the target platform (e.g., `linux/amd64`).
- `command`: Overrides the default command defined in the image.
- `tmpfs`: Mounts a temporary file system in RAM.
- `secrets`: Grants services access to sensitive data (secrets).
- `ulimits`: Configures resource limits (e.g., number of processes).
- `cpus` and `mem_limit`: Sets hard limits on CPU and memory usage.
- `sysctls`: Configures kernel parameters for the container.
- `pid`: Sets the PID namespace (e.g., `host` to share the host's PID stack).
- `extra_hosts`: Adds hostname mappings to the container's `/etc/hosts` file. This is useful for accessing external services or the host machine by a custom domain name.
  ```yaml
  extra_hosts:
    - "external.service.local:192.168.1.100"
    - "api.internal:10.0.0.50"
    # Special value to map a hostname to the host's bridge IP
    - "host.docker.internal:host-gateway"
  ```

## Common Commands

### Manage Stack
```bash
# Start the application stack
docker compose up

# Start in detached mode
docker compose up -d

# Stop and remove the stack
docker compose down
```

### Logs and Debugging
```bash
# View live logs
docker compose logs -f

# Check configuration for errors
docker compose config
```

### Execute Commands
```bash
# Run a command inside a running service container
docker compose exec web ls /usr/share/nginx/html
```

### Build Images
```bash
# Build or rebuild services
docker compose build
```