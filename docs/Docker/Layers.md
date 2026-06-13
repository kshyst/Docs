# Docker Layers

Docker images are composed of a series of read-only layers. Each layer represents an instruction in the image's Dockerfile (e.g., `RUN`, `COPY`, `ADD`). These layers are stacked on top of each other to form the final image.

## How Layers Work

Docker uses a **Union File System** to combine these layers into a single coherent filesystem. Each layer only contains the differences (diffs) from the layer below it.

### Copy-on-Write (CoW) Strategy
Docker utilizes a **Copy-on-Write (CoW)** strategy for maximum efficiency.
- If a layer needs to modify a file from a lower layer, it first copies the file to the current layer and then applies the changes.
- This ensures that the underlying layers remain unchanged and can be shared among multiple images.

The unique combination of these layers is identified by a [Docker Digest](Docker%20Digests.md), which serves as an immutable fingerprint of the image content.

## Types of Layers

- **Base Layer**: The initial layer of the image, typically an operating system like `ubuntu` or `alpine`.
- **Build Layers**: Created by `RUN` instructions, these layers usually involve installing packages or compiling code.
- **Data Layers**: Created by `COPY` or `ADD` instructions, these layers add files from the host machine to the image.
- **Metadata Layers**: Instructions like `ENV`, `EXPOSE`, and `LABEL` add metadata to the image rather than files to the filesystem.
- **Executable Layers**: `CMD` and `ENTRYPOINT` specify what process should run when a container is started from the image.

## Benefits of Layering

- **Storage Efficiency**: Multiple images can share the same base and intermediate layers, significantly reducing disk usage.
- **Faster Builds**: When you modify a Dockerfile and rebuild the image, Docker only rebuilds the layers that have changed and reuses the rest from the cache.
- **Faster Pulls**: When pulling an image, Docker only downloads the layers that you don't already have on your local machine.

## Viewing Image Layers
You can inspect the layers of an image using the `docker history` command:
```bash
docker history my-image:latest
```