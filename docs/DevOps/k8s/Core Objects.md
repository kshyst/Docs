# Core Kubernetes Objects

Kubernetes works by managing objects. An object is a persistent entity in the system.

## 1. Pods
A **Pod** is the smallest deployable unit of computing that you can create and manage in Kubernetes. A Pod is a group of one or more containers, with shared storage and network resources, and a specification for how to run the containers.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
  - name: nginx
    image: nginx:latest
```

## 2. Deployments
A **Deployment** provides declarative updates for Pods and ReplicaSets. You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```

## 3. Services
A **Service** is an abstract way to expose an application running on a set of Pods as a network service.

### Types of Services:
- **ClusterIP** (default): Exposes the Service on a cluster-internal IP. This makes the Service only reachable from within the cluster.
- **NodePort**: Exposes the Service on each Node's IP at a static port (the `NodePort`).
- **LoadBalancer**: Exposes the Service externally using a cloud provider's load balancer.
- **ExternalName**: Maps a Service to the contents of the `externalName` field (e.g. `foo.bar.example.com`).

## 4. ConfigMaps
A **ConfigMap** is an API object used to store non-confidential data in key-value pairs. Pods can consume ConfigMaps as environment variables, command-line arguments, or as configuration files in a volume.

## 5. Secrets
A **Secret** is an object that contains a small amount of sensitive data such as a password, a token, or a key. Using a Secret means you don't need to include confidential data in your application code.
- `Opaque`: Generic secret.
- `kubernetes.io/service-account-token`: Service account token.
- `kubernetes.io/dockerconfigjson`: Docker credentials.
- `kubernetes.io/tls`: TLS certificate and key.
