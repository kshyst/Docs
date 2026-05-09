
## 1. The Core Architecture: Pull-Based vs. Push-Based

The biggest technical shift in GitOps is moving away from a CI tool (like Jenkins or GitHub Actions) "pushing" code to your cluster.

### The Pull-Based Model (Standard GitOps)

- **The Agent:** You install an operator (like **ArgoCD** or **Flux**) inside your Kubernetes cluster.
    
- **The Loop:** This agent constantly "polls" your Git repo and compares it to what is actually running in the cluster.
    
- **Drift Detection:** If someone manually changes a service using a command like `kubectl edit`, the GitOps agent will see the "drift" and automatically overwrite it to match what is in Git.
    
- **Security:** Your CI system doesn't need "Admin" keys to your cluster. The cluster reaches _out_ to Git, keeping your credentials internal.
    

### The Push-Based Model (Traditional)

- Your CI pipeline finishes building an image and then executes a script to update the cluster.
    
- **Risk:** If the script fails halfway, your environment is in an inconsistent state. GitOps solves this by ensuring the "Desired State" is always eventually reached.
    

---

## 2. Managing Image Versions (The Workflow)

In a microservices setup, you typically have two types of repositories:

1. **App Repo:** Contains the source code (e.g., `auth-service`).
    
2. **Config Repo:** Contains the deployment manifests (YAML files, Helm charts, or Kustomize).
    

### Step-by-Step Version Tracking:

- **Step 1:** You edit code in the **App Repo** and push it.
    
- **Step 2:** CI builds a new image: `my-registry.io/auth-service:v1.2.3`.
    
- **Step 3:** An automated process (like an "Image Updater") modifies the **Config Repo**, changing the image tag from `v1.2.2` to `v1.2.3`.
    
- **Step 4:** The GitOps Controller sees the change in the **Config Repo** and pulls the new image into the cluster.
    

---

## 3. Testing in GitOps: Environment Branches vs. Folders

How do you test a new image before it hits production? There are two main styles:

### A. The Folder Strategy (Recommended)

You have one branch (`main`), but separate folders for environments:

Plaintext

```
/deployments
  /staging
    auth-service-v1.2.3-rc1.yaml
  /production
    auth-service-v1.2.2.yaml
```

- **Testing:** You update the image version in the `/staging` folder first.
    
- **Promotion:** Once verified, you copy that version string to the `/production` folder. This ensures the exact same image SHA is moved forward.
    

### B. The Branching Strategy

You have a `staging` branch and a `production` branch.

- **Testing:** You merge your feature into `staging`. The GitOps agent for the staging cluster sees the update and deploys it.
    
- **Promotion:** You create a Pull Request from `staging` to `production`.
    

---

## 4. Top Tools for 2026

|**Tool**|**Style**|**Best For**|
|---|---|---|
|**ArgoCD**|Visual/Centralized|Teams that want a high-level UI to see the "health" of all microservices at once.|
|**FluxCD**|Lightweight/Native|Teams that prefer a "set it and forget it" approach that feels like part of Kubernetes.|
|**Kargo**|Multi-stage Promotion|Specifically designed to manage the "promotion" of images from Dev → Staging → Prod.|
|**Crossplane**|Infrastructure GitOps|Extends GitOps to manage cloud resources (RDS, S3) alongside your microservices.|

---

## 5. Why use GitOps for Microservices?

- **Instant Rollbacks:** If `v1.2.3` of your Payment Service breaks the API, you don't need to re-run a 20-minute CI build. You just `git revert` the last commit in your Config Repo, and the cluster rolls back in seconds.
    
- **Observability:** You can look at your Git history and see exactly who changed which service version and when.
    
- **Consistency:** It eliminates the "it worked in staging but not in prod" problem because the manifests are identical, differing only by the environment-specific variables you define.
    

> **Note:** GitOps requires discipline. You must **never** use the `:latest` tag for your images. Always use specific versions (SemVer) or the Git Commit SHA so the GitOps controller knows exactly which immutable "slice" of code it is deploying.