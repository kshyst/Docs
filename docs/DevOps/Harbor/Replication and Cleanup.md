# Replication and Cleanup

Harbor provides features to manage image distribution and storage consumption effectively.

## Replication

Replication allows you to synchronize images between Harbor and other registries (including other Harbor instances, Docker Hub, ECR, GCR, etc.).

### Replication Modes:
- **Push-based**: The local Harbor instance pushes images to a remote registry.
- **Pull-based**: The local Harbor instance pulls images from a remote registry.

### Use Cases:
- **Disaster Recovery**: Replicate images to a secondary Harbor instance in a different region.
- **Load Balancing**: Distribute images to local registries near your compute clusters to speed up deployments.
- **Proxy Cache**: Use Harbor as a pull-through cache for public registries (like Docker Hub) to save bandwidth and avoid rate limits.

## Storage Management and Cleanup

As more images and versions are pushed, storage usage grows. Harbor provides mechanisms to reclaim space.

### 1. Tag Retention Policies
You can define rules to automatically delete old or unused tags.
*Example: "Keep only the last 10 tags for each repository" or "Keep tags pushed within the last 30 days".*

### 2. Garbage Collection (GC)
Deleting an image tag in Harbor doesn't immediately delete the data from the underlying storage (the blobs). To actually free up space, you must run **Garbage Collection**.

- **Manual GC**: Triggered by a system admin from the UI or API.
- **Scheduled GC**: Can be scheduled to run periodically (e.g., every Sunday at 3 AM).
- **Dry Run**: It's highly recommended to perform a "Dry Run" before actual GC to see how much space will be reclaimed and ensure no critical data is lost.

> **Note**: During Garbage Collection, it is recommended to set Harbor to "Read-Only" mode to prevent data inconsistency if someone tries to push an image while GC is running.
