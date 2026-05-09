# Kubectl Cheatsheet

`kubectl` is the command-line tool used to interact with Kubernetes clusters.

## Context and Configuration
```bash
# Get current context
kubectl config current-context

# Switch to another context
kubectl config use-context <context-name>

# List all contexts
kubectl config get-contexts
```

## Resource Creation and Management
```bash
# Create a resource from a file
kubectl apply -f <filename>.yaml

# List all pods in the current namespace
kubectl get pods

# List all pods in all namespaces
kubectl get pods -A

# Get detailed information about a pod
kubectl describe pod <pod-name>

# Delete a resource
kubectl delete -f <filename>.yaml
kubectl delete pod <pod-name>
```

## Debugging and Interaction
```bash
# View logs for a pod
kubectl logs <pod-name>

# Stream logs
kubectl logs -f <pod-name>

# Execute a command in a container
kubectl exec -it <pod-name> -- /bin/bash

# Port forward to a local port
kubectl port-forward <pod-name> 8080:80
```

## Scaling and Updates
```bash
# Scale a deployment
kubectl scale deployment <deployment-name> --replicas=5

# Check rollout status
kubectl rollout status deployment <deployment-name>

# Undo a rollout
kubectl rollout undo deployment <deployment-name>
```

## Node and Cluster Info
```bash
# Get nodes
kubectl get nodes

# Get cluster info
kubectl cluster-info

# List all resources
kubectl get all
```
