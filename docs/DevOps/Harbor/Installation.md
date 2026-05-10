# Harbor Installation

Installing Harbor involves several steps, from preparing the environment to configuring the software and running the installer.

## Prerequisites

Before installing Harbor, ensure your system meets the following requirements:

- **Docker**: Version 17.06.0-ce or higher.
- **Docker Compose**: Version 1.18.0 or higher.
- **Hardware**:
    - CPU: 2 CPU
    - RAM: 4 GB
    - Disk: 40 GB

## Installer Types

Harbor offers two types of installers:

1. **Online Installer**: Downloads Harbor images from Docker Hub during the installation process.
2. **Offline Installer**: Includes all Harbor images in the package, making it suitable for environments without internet access.

## Basic Configuration

The main configuration file for Harbor is `harbor.yml`. You must rename `harbor.yml.tmpl` to `harbor.yml` before editing.

### Key Parameters:

- **hostname**: The FQDN of your Harbor instance (e.g., `harbor.example.com`).
- **http**: Basic configuration for HTTP access.
- **https**: Highly recommended for production. Requires paths to your SSL certificate and private key.
- **harbor_admin_password**: The default password for the `admin` user.
- **database**: Configuration for the internal PostgreSQL database.
- **data_volume**: The directory where Harbor stores its data (images, database, logs).

```yaml
hostname: harbor.myexample.com
http:
  port: 80
https:
  port: 443
  certificate: /path/to/certificate
  private_key: /path/to/private_key
```

## Running the Install Script

Once `harbor.yml` is configured, run the `install.sh` script. You can include optional components like Trivy or Notary.

```bash
# Basic installation
sudo ./install.sh

# Installation with Trivy (vulnerability scanning)
sudo ./install.sh --with-trivy

# Installation with Chart Repository (Helm)
sudo ./install.sh --with-chartmuseum
```

After the installation is complete, you can access the Harbor web interface at the hostname you specified.
