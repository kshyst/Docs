# SSL and Let's Encrypt

One of Traefik's standout features is its native support for ACME (Automatic Certificate Management Environment), allowing it to automatically generate and renew SSL certificates via **Let's Encrypt**.

## Certificate Resolvers

To use Let's Encrypt, you must define a `certificatesResolver` in your **Static Configuration**.

```yaml
# traefik.yml (Static Configuration)
certificatesResolvers:
  myresolver:
    acme:
      email: admin@example.com
      storage: acme.json # File where certificates will be saved
      httpChallenge:
        entryPoint: web # Use port 80 for the HTTP challenge
```

## Challenge Types

### 1. HTTP Challenge (Standard)
Traefik proves ownership of the domain by serving a secret file on port 80.
- **Pros**: Easy to set up.
- **Cons**: Requires port 80 to be open and reachable from the internet.

### 2. DNS Challenge
Traefik proves ownership by creating a temporary TXT record in your DNS provider.
- **Pros**: Supports wildcard certificates (`*.example.com`) and internal services.
- **Cons**: Requires API access to your DNS provider (Cloudflare, DigitalOcean, etc.).

## Enabling SSL on a Router

Once the resolver is configured, tell your router to use it:

```yaml
# Docker Label Example
labels:
  - "traefik.http.routers.myapp.tls.certresolver=myresolver"
  - "traefik.http.routers.myapp.tls.domains[0].main=myapp.example.com"
```

!!! warning
    Ensure the `acme.json` file has `0600` permissions (read/write only for owner). Traefik will refuse to use it otherwise.
