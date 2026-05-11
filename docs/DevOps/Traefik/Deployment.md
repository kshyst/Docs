# Deploying Traefik

The most common way to deploy Traefik in a containerized environment is using Docker Compose. This allows Traefik to monitor the Docker daemon and automatically configure itself for other containers.

## Docker Compose Example

Create a `docker-compose.yml` file for Traefik:

```yaml
version: '3.8'

services:
  traefik:
    image: traefik:v3.0
    container_name: traefik
    command:
      - "--api.insecure=true" # Enable the dashboard (insecure mode)
      - "--providers.docker=true" # Enable the Docker provider
      - "--providers.docker.exposedbydefault=false" # Only expose containers with traefik.enable=true
      - "--entrypoints.web.address=:80" # Define HTTP entrypoint
    ports:
      - "80:80"
      - "8080:8080" # Dashboard port
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro # Allow Traefik to listen to Docker
      - ./traefik.yml:/etc/traefik/traefik.yml:ro # Static configuration
    restart: always

networks:
  default:
    name: traefik_public
```

## Exposing a Service

Once Traefik is running, you can expose other containers simply by adding labels:

```yaml
services:
  my-app:
    image: nginx
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.myapp.rule=Host(`myapp.example.com`)"
      - "traefik.http.routers.myapp.entrypoints=web"
```

!!! note
    In a production environment, you should never use `--api.insecure=true`. Instead, use a router with authentication and HTTPS to secure the dashboard.
