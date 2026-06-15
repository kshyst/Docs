# Docker Digests

A **Docker Digest** is a content-addressable identifier for a Docker image. It is a SHA-256 hash of the image's manifest (which includes the layers and configuration). Unlike tags, which are mutable, digests are **immutable**.

## Why Use Digests?

Using digests instead of tags provides several advantages in production environments:

- **Immutability**: A tag like `latest` or `v1.0` can be updated to point to a different image. A digest always refers to the exact same content.
- **Security**: Since the digest is a hash of the content, you can be sure that the image you pull is exactly what you expect. This prevents "image poisoning" or accidental updates.
- **Reproducibility**: Ensuring that every environment (Development, Staging, Production) uses the exact same image bit-for-bit.

## Tags vs. Digests

| Feature | Tags | Digests |
| :--- | :--- | :--- |
| **Example** | `nginx:latest` | `nginx@sha256:0d17...` |
| **Mutability** | Mutable (can change) | Immutable (constant) |
| **Reliability** | Medium | High |
| **Usage** | Development / Versioning | Production / CI/CD |

## How to Find an Image Digest

### 1. Using Docker CLI
You can see the digests of images on your local system by adding the `--digests` flag:

```bash
docker images --digests
```

To get the digest for a specific image:
```bash
docker inspect --format='{{index .RepoDigests 0}}' <image_name>
```

### 2. From Docker Hub
When you visit an image page on Docker Hub (e.g., [nginx](https://hub.docker.com/_/nginx)), you can find the digest listed under the **Tags** tab for each specific version.

## How to Use Digests

### In the CLI
You can pull or run an image using its digest:

```bash
docker pull nginx@sha256:0d17b565c37bc54b565a37e1e672da079c6b77259f939226174a7321a4f0050c
```

### In a Dockerfile
Pinning your base image with a digest is a best practice for production Dockerfiles:

```dockerfile
# Instead of: FROM node:18
FROM node:18@sha256:56db3642398544d03e54b680da3f5383568971f49a37c449339a0f6225642d97

WORKDIR /app
COPY . .
RUN npm install
CMD ["npm", "start"]
```

!!! tip
    While using digests is great for stability, remember that you won't automatically receive security patches. You should have a process to regularly update your digests when new base image versions are released.

## Troubleshooting: Digest Mismatches

In modern Docker builds (BuildKit/Buildx), you might encounter situations where your local digest doesn't match the registry digest even if no code has changed. This is often caused by **Attestations** and the use of **Manifest Lists**.

If you are experiencing "infinite recreation loops" in tools like Watchtower or WUD, see [Docker Attestations & Manifest Mismatches](Docker%20Attestations.md).
