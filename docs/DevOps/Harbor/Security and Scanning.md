# Security and Scanning

Security is a core pillar of Harbor. It provides tools to ensure that images are free from known vulnerabilities and haven't been tampered with.

## Vulnerability Scanning

Harbor integrates with scanners like **Trivy** (the default) to analyze images for security flaws.

### Features:
- **Automatic Scanning**: You can enable "Scan on push" in project settings so every image is scanned immediately after being uploaded.
- **Scheduled Scanning**: System admins can schedule system-wide scans (e.g., daily or weekly).
- **Vulnerability Thresholds**: You can prevent images with a certain severity level (e.g., "High" or "Critical") from being pulled.

## Image Signing (Content Trust)

Image signing allows you to verify the origin and integrity of an image. Harbor supports two main methods:

1. **Notary**: Uses the Docker Content Trust (DCT) mechanism.
2. **Cosign**: A more modern approach that is part of the Sigstore project.

By requiring signed images, you can ensure that only authorized and verified images are deployed to your clusters.

## CVE Exception Lists

Sometimes, a vulnerability is known but cannot be fixed immediately, or it is determined to be a "false positive" or "no-risk" in your specific environment. Harbor allows you to create **CVE Exception Lists** at both the system and project levels to bypass pull blocks for specific CVEs for a set period.

## Security Best Practices in Harbor

- **Enable Scan on Push**: Always ensure images are scanned as soon as they arrive.
- **Configure Deployment Security**: Set policies to block the deployment of images with critical vulnerabilities or those that are unsigned.
- **Use Robot Accounts**: Avoid using personal credentials in CI/CD pipelines.
- **Regularly Audit Logs**: Monitor Harbor's access logs for suspicious activity.
