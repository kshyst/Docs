# Installation and Setup

There are many ways to set up a Kubernetes cluster depending on your needs.

## 1. Minikube (Local Development)

Minikube is a tool that lets you run Kubernetes locally. Minikube runs a single-node Kubernetes cluster on your personal computer (including Windows, macOS and Linux) so that you can try out Kubernetes or for daily development work.

**Installation**:
```bash
# macOS
brew install minikube

# Linux
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

**Starting Minikube**:
```bash
minikube start
```

## 2. K3s (Lightweight/Edge)

K3s is a highly available, certified Kubernetes distribution designed for production workloads in resource-constrained, remote locations or on IoT devices.

**Quick start**:
```bash
curl -sfL https://get.k3s.io | sh -
# Check for node readiness
sudo k3s kubectl get node
```

## 3. Kubeadm (Manual Setup)

Kubeadm is a tool that performs the actions necessary to get a minimum viable, secure cluster up and running in a way that is conformant with Kubernetes best practices.

**Basic steps**:
1. Install a container runtime (e.g., Docker or containerd).
2. Install `kubeadm`, `kubelet`, and `kubectl` on all nodes.
3. Initialize the control plane: `kubeadm init`.
4. Join worker nodes to the cluster: `kubeadm join <control-plane-host>:<control-plane-port> --token <token> ...`
5. Install a CNI (Container Network Interface) plugin.

## 4. Managed Services

Cloud providers offer managed Kubernetes services that handle the control plane management and maintenance for you.

- **GKE (Google Kubernetes Engine)**: The original managed K8s service. Deeply integrated with Google Cloud.
- **EKS (Amazon Elastic Kubernetes Service)**: Highly scalable and secure.
- **AKS (Azure Kubernetes Service)**: Simplified deployment and management on Microsoft Azure.

### Benefits of Managed Services:
- Automatic updates and patching.
- High availability of the control plane.
- Integrated monitoring and logging.
- Integrated IAM (Identity and Access Management).
