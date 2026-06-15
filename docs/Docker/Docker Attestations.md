# Docker Attestations & Manifest Mismatches

With modern Docker (BuildKit and Buildx), images often include **attestations**. While these provide valuable metadata, they can sometimes cause "infinite recreation loops" in automation tools like **WUD (What's Up Docker)** or **Watchtower** due to digest mismatches.

## What are Build Attestations?

Build attestations are metadata attached to an image during the build process. There are two primary types:

1.  **Provenance**: Describes how the image was built (source code, build arguments, etc.).
2.  **SBOM (Software Bill of Materials)**: Lists the packages and dependencies installed in the image.

## The Manifest List Wrapper

When Docker attaches attestations, it wraps your image in a **Manifest List** (or Image Index) inside the registry. This happens even if you are building for a single platform (e.g., `linux/amd64`).

- **The Tag**: Points to the top-level **Manifest List Digest**.
- **The Image**: Inside the list, there is a **Platform-Specific Digest** for the actual image.

### The Issue: The Digest Mismatch Loop

This structure causes a conflict with some deployment tools:

1.  **Pulling**: When your server pulls the image, it resolves the index and often stores the top-level **Index Digest** in its local `RepoDigests`.
2.  **Comparing**: Monitoring tools (like WUD) might compare the local digest against the **Platform-Specific Digest** in the registry.
3.  **Conflict**: Because these two digests are different (Index vs. Image), the tool thinks a "new" version is available.
4.  **Loop**: The tool triggers a re-pull and container recreation, finds the "mismatch" again, and repeats the process indefinitely.

## The Fix: Disabling Default Attestations

If you don't need provenance or SBOM metadata and want to avoid this loop, you can tell Docker to build a standard single-platform image without the index wrapper.

### Method 1: Environment Variable
Set this variable in your CI/CD environment (e.g., GitLab CI or GitHub Actions):

```bash
export BUILDX_NO_DEFAULT_ATTESTATIONS=1
docker buildx build -t my-image:latest --push .
```

### Method 2: Build Flags
You can explicitly disable attestations in your build command:

```bash
docker buildx build \
  --provenance=false \
  --sbom=false \
  -t my-image:latest --push .
```

## When to Use This Fix

| Scenario | Recommendation |
| :--- | :--- |
| **High Security / Compliance** | Keep attestations enabled and update your monitoring tools to handle Manifest Lists. |
| **Standard Production** | Disable attestations if they cause recreation loops in Watchtower/WUD. |
| **Local Development** | Usually not an issue as attestations are primarily a registry-side feature. |

!!! warning
    Disabling attestations removes valuable security metadata. Only use this fix if you are experiencing the specific "digest mismatch" recreation loop and cannot update your monitoring software.
