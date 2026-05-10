# Harbor Best Practices

For a production-ready Harbor deployment, follow these best practices to ensure security, reliability, and performance.

## 1. HTTPS and SSL/TLS
Always run Harbor over HTTPS.
- Use valid certificates from a trusted CA (like Let's Encrypt).
- If using self-signed certificates, ensure all Docker clients and Kubernetes nodes have the CA certificate added to their trusted store.

## 2. External Storage
By default, Harbor stores images on the local filesystem. For production:
- Use **S3-compatible storage** (AWS S3, MinIO, Google Cloud Storage, Azure Blob Storage).
- This provides better durability, scalability, and allows for easier Harbor upgrades or migrations.

## 3. Database Backups
Harbor uses PostgreSQL to store metadata (users, projects, roles, scan results).
- Regularly back up the PostgreSQL database.
- Use external managed database services (like AWS RDS) if possible for easier management and high availability.

## 4. High Availability (HA)
For critical environments, deploy Harbor in an HA configuration:
- Use a load balancer in front of multiple Harbor nodes.
- Use a shared external database (PostgreSQL).
- Use a shared external key-value store (Redis) for session management and job queuing.
- Use shared external storage (S3).

## 5. Regular Updates
- Stay updated with the latest Harbor releases to benefit from new features and security patches.
- Always check the release notes and migration guides before upgrading.

## 6. Monitoring and Logging
- Integrate Harbor with monitoring tools like **Prometheus** and **Grafana**. Harbor provides a `/metrics` endpoint.
- Ship Harbor logs to a centralized logging system (like ELK stack or Graylog) for better auditability and troubleshooting.

## 7. Use Robot Accounts for CI/CD
Never hardcode personal admin credentials in CI/CD scripts. Create specific robot accounts for each project or pipeline with the minimum required permissions.
