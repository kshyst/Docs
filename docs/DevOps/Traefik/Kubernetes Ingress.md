# Traefik in Kubernetes

While Traefik supports the standard Kubernetes **Ingress** object, it also provides its own Custom Resource Definitions (CRDs) which offer more power and flexibility.

## Ingress vs. IngressRoute

The standard Ingress object is limited in how it handles advanced features like TCP/UDP routing or complex middlewares. Traefik's `IngressRoute` CRD provides a more "Traefik-native" way to configure routing.

### IngressRoute Example

```yaml
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: myapp-route
spec:
  entryPoints:
    - websecure
  routes:
    - match: Host(`myapp.example.com`) && PathPrefix(`/api`)
      kind: Rule
      services:
        - name: my-api-service
          port: 80
      middlewares:
        - name: auth-middleware
  tls:
    certResolver: myresolver
```

## Deployment via Helm

The recommended way to install Traefik in Kubernetes is via the official Helm chart:

```bash
helm repo add traefik https://traefik.github.io/charts
helm repo update
helm install traefik traefik/traefik
```

## Key CRDs
- **IngressRoute**: HTTP/HTTPS routing.
- **IngressRouteTCP / IngressRouteUDP**: L4 routing.
- **Middleware**: Define reusable middlewares as K8s objects.
- **TLSOption**: Configure TLS versions and cipher suites globally.

!!! info
    Using CRDs allows you to manage Traefik configuration using standard `kubectl` commands and GitOps workflows while accessing features not available in standard Ingress.
