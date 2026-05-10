# Deploying Grafana

Grafana can be deployed in various environments. The most common methods involve using Docker for easy management and scalability.

## Default Configuration

-   **Port**: `3000`
-   **Default Credentials**: `admin` / `admin` (You will be prompted to change this on first login)

## Docker CLI

You can quickly start a Grafana instance using the `docker run` command:

```bash
docker run -d -p 3000:3000 --name=grafana grafana/grafana-oss
```

## Docker Compose

For a more robust deployment, especially when including other tools like Prometheus or Loki, Docker Compose is recommended. This setup includes volume persistence for the Grafana SQLite database.

```yaml
version: '3.8'

services:
  grafana:
    image: grafana/grafana-oss:latest
    container_name: grafana
    ports:
      - "3000:3000"
    volumes:
      - grafana-storage:/var/lib/grafana
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    restart: unless-stopped

volumes:
  grafana-storage:
```

### Persistence
By default, Grafana stores its data (users, dashboards, etc.) in a local SQLite database located at `/var/lib/grafana`. Using a Docker volume ensures this data survives container restarts and upgrades.

## Environment Variables
Grafana can be configured using environment variables. Common ones include:

- `GF_SECURITY_ADMIN_PASSWORD`: Sets the admin password.
- `GF_SERVER_DOMAIN`: The domain name of the Grafana server.
- `GF_INSTALL_PLUGINS`: A comma-separated list of plugins to install on startup.
