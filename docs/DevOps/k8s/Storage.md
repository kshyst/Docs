# Storage in Kubernetes

On-disk files in a container are ephemeral, which presents some problems for anything but trivial applications when running in containers. Kubernetes provides several abstractions to manage persistent data.

## 1. Volumes
At its core, a **Volume** is just a directory, possibly with some data in it, which is accessible to the containers in a Pod. How that directory comes to be, the medium that backs it, and its contents are determined by the particular volume type used.
- `emptyDir`: Temporary volume that is deleted when the Pod is removed.
- `hostPath`: Mounts a file or directory from the host node's filesystem into your Pod.

## 2. PersistentVolumes (PV)
A **PersistentVolume** is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes. It is a resource in the cluster just like a node is a cluster resource.

## 3. PersistentVolumeClaims (PVC)
A **PersistentVolumeClaim** is a request for storage by a user. It is similar to a Pod. Pods consume node resources and PVCs consume PV resources.

## 4. StorageClasses
A **StorageClass** allows administrators to describe the "classes" of storage they offer. Different classes might map to quality-of-service levels, or to backup policies, or to arbitrary policies determined by the cluster administrators.

### Workflow Example:
1. Admin creates a **StorageClass**.
2. User creates a **PersistentVolumeClaim** requesting a specific StorageClass and size.
3. K8s automatically creates a **PersistentVolume** that matches the PVC.
4. User creates a **Pod** that uses the PVC as a volume.

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: standard
```
