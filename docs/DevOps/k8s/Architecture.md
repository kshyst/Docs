# Kubernetes Architecture

A Kubernetes cluster consists of a set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.


## Control Plane Components

The control plane's components make global decisions about the cluster (for example, scheduling), as well as detecting and responding to cluster events.

### 1. kube-apiserver
The API server is the front end for the Kubernetes control plane. It exposes the Kubernetes API and is designed to scale horizontally.

### 2. etcd
Consistent and highly-available key-value store used as Kubernetes' backing store for all cluster data.

### 3. kube-scheduler
Component that watches for newly created Pods with no assigned node, and selects a node for them to run on. Factors taken into account for scheduling decisions include: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, etc.

### 4. kube-controller-manager
Runs controller processes. Logically, each controller is a separate process, but to reduce complexity, they are all compiled into a single binary and run in a single process.
- **Node controller**: Responsible for noticing and responding when nodes go down.
- **Job controller**: Watches for Job objects that represent one-off tasks, then creates Pods to run those tasks to completion.
- **EndpointSlice controller**: Populates EndpointSlice objects (to provide a link between Services and Pods).
- **ServiceAccount controller**: Create default ServiceAccounts for new namespaces.

## Node Components

Node components run on every node, maintaining running pods and providing the Kubernetes runtime environment.

### 1. kubelet
An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod. The kubelet takes a set of PodSpecs that are provided through various mechanisms and ensures that the containers described in those PodSpecs are running and healthy.

### 2. kube-proxy
A network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept. It maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.

### 3. Container Runtime
The software that is responsible for running containers. Kubernetes supports several container runtimes: **Docker**, **containerd**, **CRI-O**, and any implementation of the Kubernetes CRI (Container Runtime Interface).
