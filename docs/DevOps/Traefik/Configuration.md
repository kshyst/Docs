# Traefik Configuration

Traefik distinguishes between two types of configuration: **Static** and **Dynamic**. Understanding this split is crucial for effective management.

## Static Configuration

Static configuration defines the core setup of Traefik—things that do not change frequently and require a restart if modified.

- **Entrypoints**: Defining ports (80, 443).
- **Providers**: Connecting to Docker, Kubernetes, or file-based sources.
- **Certificates Resolvers**: Configuring Let's Encrypt details.
- **Logging**: Setting up access and application logs.

Static configuration can be provided via a file (`traefik.yml` or `traefik.toml`), environment variables, or CLI arguments.

## Dynamic Configuration

Dynamic configuration defines how Traefik handles incoming requests. This part of the configuration is automatically updated without restarting Traefik.

- **Routers**: Matching rules for incoming requests.
- **Middlewares**: Request transformations.
- **Services**: Load balancing and backend definitions.

### Providers
Dynamic configuration is fetched from **Providers**:
- **Docker**: Labels on your containers.
- **Kubernetes**: Ingress or IngressRoute objects.
- **File**: A separate YAML/TOML file that Traefik watches for changes.
- **HTTP**: An external API endpoint.

## Configuration Merging

When Traefik starts, it reads the static configuration. It then connects to the enabled providers to discover the dynamic configuration. Traefik continuously monitors these providers and merges any changes into its internal state in real-time.

!!! info
    If you use a file provider for dynamic configuration, ensure you distinguish it from your static configuration file to avoid confusion.
