# Docker v1 vs. v2

In the Docker ecosystem, "v1 vs. v2" typically refers to two major transitions: the **Registry API** and **Docker Compose**.

---

## 1. Docker Registry API (v1 vs. v2)

The Registry API is the protocol used by the Docker client to communicate with a registry (like Docker Hub).

### Registry v1 (Legacy)
- **Architecture**: Relied on a simple file-based storage mapping.
- **Identification**: Images were identified by IDs that weren't necessarily tied to the actual content in a verifiable way.
- **Performance**: Slower pulls/pushes due to less efficient layer management.
- **Status**: **Deprecated** and removed from modern Docker environments.

### Registry v2 (Modern)
- **Architecture**: Introduced **Content-Addressable Storage**.
- **[Digests](Docker%20Digests.md)**: Images are identified by a SHA-256 hash of their content. This ensures that the content cannot be tampered with.
- **Manifest Lists**: Enabled multi-platform images (e.g., pulling the same tag on ARM and AMD64 works automatically).
- **Performance**: Faster, more reliable, and supports concurrent layer uploads.

---

## 2. Docker Compose (v1 vs. v2)

This refers to the tool used to manage multi-container applications.

### Compose v1 (`docker-compose`)
- **Language**: Written in **Python**.
- **Installation**: Usually installed as a separate standalone binary (`/usr/local/bin/docker-compose`).
- **Command**: `docker-compose up` (with a hyphen).
- **Status**: Reached End-of-Life (EOL) in June 2023.

### Compose v2 (`docker compose`)
- **Language**: Rewritten in **Go** (the same language as Docker Engine).
- **Installation**: Integrated as a **Docker CLI plugin**.
- **Command**: `docker compose up` (using a space, no hyphen).
- **Features**: 
    - Better performance and resource management.
    - Full compatibility with the Compose Specification.
    - Support for modern BuildKit features by default.
    - Consistent behavior with the core Docker CLI.

---

## Summary Comparison

| Feature | Version 1 (Legacy) | Version 2 (Modern) |
| :--- | :--- | :--- |
| **Registry API** | ID-based, slow, insecure | Content-addressable, fast, [digests](Docker%20Digests.md) |
| **Compose Tool** | `docker-compose` (Python) | `docker compose` (Go plugin) |
| **Maintenance** | Deprecated / EOL | Active / Recommended |
| **Performance** | Basic | Optimized (BuildKit/Parallelism) |

## How to Check Your Version

To check your Docker Compose version:
```bash
docker compose version
```

To see if you are using Registry v2 (standard for all modern registries):
Most modern registries (Docker Hub, GHCR, ECR) only support V2. You can check the `Docker-Distribution-Api-Version` header in a registry response, which should be `registry/2.0`.
