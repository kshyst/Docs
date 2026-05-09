# Introduction to Kubernetes

Kubernetes (often abbreviated as **K8s**) is an open-source platform designed to automate deploying, scaling, and operating application containers.

## What is Kubernetes?

Originally developed by Google and now maintained by the Cloud Native Computing Foundation (CNCF), Kubernetes provides a container-centric management environment. It orchestrates computing, networking, and storage infrastructure on behalf of user workloads.

## Why Use Kubernetes?

In the era of microservices, managing hundreds or thousands of containers manually is impossible. Kubernetes solves several key challenges:

- **Service Discovery and Load Balancing**: K8s can expose a container using a DNS name or its own IP address. If traffic to a container is high, K8s can load balance and distribute the network traffic.
- **Storage Orchestration**: Automatically mount a storage system of your choice, such as local storage, public cloud providers, and more.
- **Automated Rollouts and Rollbacks**: You can describe the desired state for your deployed containers, and it can change the actual state to the desired state at a controlled rate.
- **Bin Packing**: You provide K8s with a cluster of nodes that it can use to run containerized tasks. You tell K8s how much CPU and memory each container needs, and it fits containers onto your nodes to make the best use of your resources.
- **Self-healing**: K8s restarts containers that fail, replaces containers, kills containers that don't respond to your user-defined health check, and doesn't advertise them to clients until they are ready to serve.

## Brief History

- **Borg**: Google's internal cluster management system that inspired Kubernetes.
- **2014**: Google open-sourced Kubernetes.
- **2015**: Kubernetes v1.0 was released, and Google partnered with the Linux Foundation to form the CNCF.
- **Present**: Kubernetes has become the de-facto standard for container orchestration.
