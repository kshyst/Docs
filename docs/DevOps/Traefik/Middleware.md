# Traefik Middleware

Middlewares are a powerful way to transform requests before they reach your services or to modify responses before they are sent back to clients.

## How It Works

Middlewares are attached to **Routers**. When a request matches a router's rule, it passes through the chain of middlewares in the order they are defined.

## Common Middlewares

### 1. RedirectScheme (HTTP to HTTPS)
Automatically redirects all HTTP traffic to HTTPS.
```yaml
# Dynamic configuration example
http:
  middlewares:
    https-redirect:
      redirectScheme:
        scheme: https
        permanent: true
```

### 2. BasicAuth
Secures a router with a simple username and password.
```yaml
# Passwords must be hashed (use 'htpasswd')
http:
  middlewares:
    admin-auth:
      basicAuth:
        users:
          - "admin:$apr1$H6uskkkW$IgSST60n9G.p8tL8S7B.v1"
```

### 3. Compress
Enables Gzip compression to reduce bandwidth and improve load times.
```yaml
http:
  middlewares:
    test-compress:
      compress: {}
```

### 4. IPWhiteList
Restricts access to specific IP ranges.
```yaml
http:
  middlewares:
    internal-only:
      ipWhiteList:
        sourceRange:
          - "127.0.0.1/32"
          - "192.168.1.0/24"
```

## Using Middlewares with Docker Labels

You can apply these middlewares directly via Docker labels:

```yaml
services:
  app:
    image: my-app
    labels:
      - "traefik.http.routers.app.middlewares=https-redirect@file,admin-auth@docker"
```

!!! tip
    The `@docker` or `@file` suffix tells Traefik which provider the middleware is defined in. If it's defined on the same container, the suffix is optional.
