# Helm: The Package Manager for Kubernetes

Helm is a tool for managing Kubernetes charts. Charts are packages of pre-configured Kubernetes resources.

## Why Use Helm?

- **Manage Complexity**: Charts describe even the most complex apps, provide repeatable application installation, and serve as a single point of authority.
- **Easy Updates**: Take the pain out of updates with in-place upgrades and custom hooks.
- **Simple Sharing**: Charts are easy to version, share, and host on public or private servers.
- **Rollbacks**: Use `helm rollback` to roll back to an older version of a release with ease.

## Core Concepts

- **Chart**: A bundle of information necessary to create an instance of a Kubernetes application.
- **Repository**: A place where charts can be collected and shared.
- **Release**: An instance of a chart running in a Kubernetes cluster.

## Basic Commands

### Installing and Managing Charts
```bash
# Add a repository
helm repo add bitnami https://charts.bitnami.com/bitnami

# Search for a chart
helm search repo nginx

# Install a chart
helm install my-release bitnami/nginx

# List releases
helm list
```

### Upgrading and Uninstalling
```bash
# Upgrade a release
helm upgrade my-release bitnami/nginx --set service.type=NodePort

# Rollback a release
helm rollback my-release 1

# Uninstall a release
helm uninstall my-release
```

### Creating Your Own Charts
```bash
# Create a new chart
helm create my-chart

# Package a chart
helm package ./my-chart

# Lint a chart
helm lint ./my-chart
```
