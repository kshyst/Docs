# Introduction to Watchtower

Watchtower is an open-source tool designed to automate the process of updating Docker containers. It ensures that your running containers are always using the latest version of their respective images without requiring manual intervention.

## How It Works

Watchtower operates as a Docker container itself. It monitors your running containers by periodically polling the image registry (such as Docker Hub, GitHub Container Registry, or private registries).

### Image Detection Mechanism
A common question is whether Watchtower detects updates based on labels or upload times. In reality, it uses **Image Digests**:

- **Digests over Labels**: Watchtower does not rely on image labels or creation timestamps. Instead, it compares the unique **SHA256 digest** (hash) of the local image with the digest of the image available in the remote registry for the same tag (e.g., `latest`).
- **Detection**: If the remote digest differs from the local one, Watchtower knows a new version is available, even if the tag remains the same.

When Watchtower detects an update:
1. It gracefully shuts down the running container.
2. It pulls the new image (verified by the new digest).
3. It restarts the container using the same options it was originally started with (e.g., environment variables, volume mounts, network settings).

## Why Use Watchtower?

- **Automated CI/CD**: Watchtower bridges the gap between image building and deployment. Once a new image is pushed to the registry, Watchtower ensures it is deployed automatically.
- **Maintenance Reduction**: Eliminates the need to manually `docker pull` and `docker restart` containers across multiple environments.
- **Security**: Keeps your applications up-to-date with the latest security patches bundled in newer image versions.
- **Simplicity**: It requires minimal configuration to get started and integrates seamlessly into existing Docker environments.
