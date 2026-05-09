# Kubernetes Networking

Networking in Kubernetes is designed to ensure that all components (Pods, Services, and Nodes) can communicate with each other.

## The Kubernetes Networking Model

The Kubernetes networking model is based on three fundamental rules:
1. All Pods can communicate with all other Pods without using NAT.
2. All Nodes can communicate with all Pods (and vice-versa) without NAT.
3. The IP that a Pod sees itself as is the same IP that others see it as.

## CNI (Container Network Interface)

The **CNI** is a specification and library for configuring network interfaces in Linux containers. Kubernetes uses CNI plugins to implement its networking model.

### Popular CNI Plugins:
- **Calico**: Provides high-performance, scalable networking and network policy.
- **Flannel**: A simple and easy-to-configure L3 network fabric.
- **Cilium**: Uses eBPF to provide high-performance networking, security, and observability.
- **Weave Net**: Creates a virtual network that connects Docker containers across multiple hosts and enables their automatic discovery.

## Ingress and Ingress Controllers

While Services handle internal traffic or simple external exposure, **Ingress** provides a way to manage external access to Services in a cluster, typically HTTP/HTTPS.

### Ingress Features:
- SSL/TLS termination.
- Name-based virtual hosting.
- Path-based routing (L7).

### Popular Ingress Controllers:
- **NGINX Ingress Controller**: The most widely used controller.
- **Traefik**: Modern HTTP reverse proxy and load balancer.
- **HAProxy Ingress**: High-performance ingress controller based on HAProxy.
- **Istio Ingress Gateway**: For clusters using a Service Mesh.

![Networking Diagram](img/k8s-networking.png)
