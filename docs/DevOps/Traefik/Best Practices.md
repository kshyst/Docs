# Traefik Best Practices

To ensure a secure and high-performance Traefik deployment, follow these best practices for production environments.

## 1. Secure the Dashboard
By default, the Traefik dashboard is accessible on port 8080 and is unprotected. In production:
- Disable `api.insecure`.
- Create a dedicated **Router** for the dashboard on port 443 (HTTPS).
- Apply **BasicAuth** or **ForwardAuth** middleware to restrict access.

```yaml
# Example: Securing dashboard via Docker labels
labels:
  - "traefik.http.routers.dashboard.rule=Host(`traefik.example.com`)"
  - "traefik.http.routers.dashboard.service=api@internal"
  - "traefik.http.routers.dashboard.middlewares=admin-auth"
  - "traefik.http.routers.dashboard.tls.certresolver=myresolver"
```

## 2. Structured Logging
Enable access logs in **JSON format**. This makes them much easier to parse with tools like **Grafana Loki**.

```yaml
# traefik.yml (Static)
accessLog:
  filePath: "/var/log/traefik/access.log"
  format: json
  fields:
    names:
      StartUTC: drop
```

## 3. Metrics Integration
Enable the Prometheus provider to export metrics for visualization in Grafana.

```yaml
# traefik.yml (Static)
metrics:
  prometheus:
    entryPoint: metrics
    addRoutersLabels: true
```

## 4. Resource Limits and Health Checks
- **Health Checks**: Always define health checks for your services to ensure Traefik doesn't route traffic to unhealthy containers.
- **Resources**: Set CPU and memory limits for the Traefik container to prevent it from consuming all host resources during high traffic.

## 5. Use Labels Wisely
- Only enable Traefik on containers that need external access (`exposedByDefault: false`).
- Use descriptive names for your routers and services to make debugging easier in the dashboard.

## 6. Certificate Backups
Backup your `acme.json` file regularly. If you lose this file, you may hit Let's Encrypt rate limits while trying to re-generate all certificates for your domains.

!!! danger
    Never expose the Docker socket (`/var/run/docker.sock`) to the public internet. Ensure Traefik is the only service that can access it, and use `:ro` (read-only) whenever possible.
