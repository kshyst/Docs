# Installation Guide

Nexus Repository Manager is commonly deployed using Docker for ease of management and portability.

## Prerequisites

- Docker installed
- Docker Compose (recommended)
- Minimum 4GB RAM (8GB+ recommended for production)

## Quick Start with Docker

To run a basic instance of Nexus 3:

```bash
docker run -d \
  -p 8081:8081 \
  --name nexus \
  -v nexus-data:/nexus-data \
  sonatype/nexus3
```

## Production Setup with Docker Compose

Using Docker Compose allows for better management of volumes and environment variables.

### `docker-compose.yml`

```yaml
version: '3.8'

services:
  nexus:
    image: sonatype/nexus3:latest
    container_name: nexus
    restart: always
    ports:
      - "8081:8081"   # Web UI & API
      - "8082:8082"   # Optional: Port for Docker Hosted Repository
      - "8083:8083"   # Optional: Port for Docker Proxy Repository
    volumes:
      - nexus-data:/nexus-data
    environment:
      - INSTALL4J_ADD_VM_PARAMS=-Xms2g -Xmx2g -XX:MaxDirectMemorySize=2g

volumes:
  nexus-data:
```

## Post-Installation

1. **Access the Web Interface**: Navigate to `http://<server-ip>:8081`.
2. **Initial Admin Password**:
   Nexus generates a random password on the first run. Retrieve it with:
   ```bash
   docker exec nexus cat /nexus-data/admin.password
   ```
3. **Setup Wizard**: Log in with username `admin` and the retrieved password. You will be prompted to:
   - Change the admin password.
   - Configure Anonymous Access (recommended: disable for private repos).
   - Enable Pro-active Monitoring (optional).
