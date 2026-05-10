# Harbor Introduction

Harbor is an open-source trusted cloud native registry project that stores, signs, and scans content. It extends the open-source Docker Distribution by adding the functionalities usually required by an enterprise such as security, identity and management.

## Key Features

- **Role Based Access Control (RBAC)**: Users and repositories are organized via "Projects", and a user can have different permissions for different repositories in a project.
- **Vulnerability Scanning**: Harbor scans images regularly and warns users of vulnerabilities. It supports multiple scanners, with Trivy being the default.
- **Image Signing and Content Trust**: Images can be signed using Notary or Cosign to ensure their integrity and origin.
- **Replication**: Images can be replicated (synchronized) between multiple Harbor instances. This is useful for load balancing, high availability, and multi-datacenter scenarios.
- **Identity Integration and Federation**: Harbor can be integrated with external identity providers such as AD/LDAP, OIDC, and SAML.
- **Graphical User Interface (GUI)**: Harbor provides a user-friendly web console for managing repositories, projects, and users.
- **API and CLI**: Provides a full RESTful API and CLI for automation and integration with CI/CD pipelines.

## Why Harbor for Enterprise?

In enterprise environments, security and compliance are paramount. Harbor addresses these needs by:

1. **Isolation**: Projects allow for strict separation of teams and resources.
2. **Security Compliance**: Regular scanning and signing ensure that only "clean" and "trusted" images are deployed to production.
3. **Efficiency**: Replication and caching (Proxy Cache) reduce latency and save bandwidth in distributed environments.
4. **Auditability**: Harbor logs all activities, providing a clear audit trail of who accessed or modified which image.
