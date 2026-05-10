# Image Management

Managing images in Harbor involves authenticating, tagging, and pushing/pulling images using the Docker CLI or other container tools.

## Logging In

To push or pull private images, you must first log in to your Harbor instance.

```bash
docker login harbor.example.com
```
Enter your username and password or your robot account's credentials.

## Tagging Images

Images must be tagged with the Harbor hostname and the project name before they can be pushed.

**Format:** `<harbor_hostname>/<project_name>/<repository_name>:<tag>`

```bash
# Example: Tagging a local alpine image for Harbor
docker tag alpine:latest harbor.example.com/library/alpine:v1.0
```

## Pushing Images

Once tagged, push the image to Harbor:

```bash
docker push harbor.example.com/library/alpine:v1.0
```

## Pulling Images

To pull an image from Harbor:

```bash
docker pull harbor.example.com/library/alpine:v1.0
```

## Harbor as a Helm Chart Repository

Harbor can also store and manage Helm charts. This requires the `chartmuseum` component to be installed.

### Using the Helm CLI with Harbor:

1. **Add the Harbor repo to Helm:**
   ```bash
   helm repo add harbor https://harbor.example.com/chartrepo/my-project
   ```

2. **Push a chart (requires the helm-push plugin):**
   ```bash
   helm cm-push my-chart-0.1.0.tgz harbor
   ```

3. **Pull/Install a chart:**
   ```bash
   helm install my-release harbor/my-chart
   ```

Harbor's UI allows you to browse through both container images and Helm charts within the same project interface.
