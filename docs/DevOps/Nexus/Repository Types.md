# Repository Types

Nexus uses three primary types of repositories to manage components.

## 1. Proxy Repository
A proxy repository is a repository that is linked to a remote repository (e.g., Maven Central, npmjs.org). 

- **Purpose**: Caches artifacts locally to speed up builds and reduce external bandwidth.
- **Benefits**: If the external registry goes down, you can still access cached artifacts.

## 2. Hosted Repository
A hosted repository is a repository that stores components internal to your organization.

- **Purpose**: Storing internal builds, private libraries, and proprietary software.
- **Examples**: A `maven2 (hosted)` repository for your internal Java libraries, or a `docker (hosted)` repository for your private Docker images.

## 3. Group Repository
A group repository is a powerful feature that allows you to combine multiple repositories (both proxy and hosted) under a single URL.

- **Purpose**: Simplifies configuration for developers. Instead of adding multiple repository URLs to their build tools, they only need to use the Group URL.
- **Ordering**: The order of repositories in a group matters. Nexus searches them in the order they are listed.

---

## Common Formats

| Format | Common Use Case |
| :--- | :--- |
| **Maven** | Java/Kotlin dependencies and artifacts |
| **npm** | JavaScript/Node.js packages |
| **PyPI** | Python packages |
| **Docker** | Container images |
| **NuGet** | .NET packages |
| **Raw** | Any generic file (ISO, ZIP, binaries) |
