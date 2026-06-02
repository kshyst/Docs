# Docker Registry Setup in Nexus

Using Nexus as a private Docker Registry is one of its most common use cases. It allows you to host private images and proxy Docker Hub.

## 1. Create a Blob Store
Before creating the repository, it's good practice to create a dedicated blob store for Docker.
- Go to **Create blob store**.
- Type: **File**.
- Name: `docker-blob`.

## 2. Create a Hosted Docker Repository
To store your internal images:
- Go to **Repository > Repositories > Create repository**.
- Select **docker (hosted)**.
- **HTTP Port**: Set a dedicated port (e.g., `8082`). This port must be exposed in your Docker container/firewall.
- **Enable Docker V1 API**: Optional, usually not needed for modern Docker.
- **Blob store**: Select `docker-blob`.

## 3. Create a Proxy Docker Repository
To cache images from Docker Hub:
- Select **docker (proxy)**.
- **Remote storage**: `https://registry-1.docker.io`.
- **Docker Index**: Use Docker Hub (`https://index.docker.io/`).
- **HTTP Port**: Set another dedicated port (e.g., `8083`) if you want to pull directly through the proxy.

## 4. Using the Registry

### Login
```bash
docker login <your-nexus-ip>:8082
```

### Push an Image
```bash
docker tag my-app:latest <your-nexus-ip>:8082/my-app:latest
docker push <your-nexus-ip>:8082/my-app:latest
```

### Pull an Image
```bash
docker pull <your-nexus-ip>:8082/my-app:latest
```

## Troubleshooting

### Insecure Registry
If you are not using HTTPS, you must add your Nexus registry to the Docker `daemon.json` file:

```json
{
  "insecure-registries": ["<your-nexus-ip>:8082", "<your-nexus-ip>:8083"]
}
```
Then restart Docker: `systemctl restart docker`.
