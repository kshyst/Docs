# Project Guidelines: Kshyst Docs

## Documentation Structure
This project is an **MkDocs** site using the **Material** theme.
- **Root Directory:** Contains configuration files like `mkdocs.yml`, `.gitlab-ci.yml`, and `requirements.txt`.
- **`docs/` Folder:** The core documentation content, organized by categories (e.g., AI, Cyber Security, Linux, Networking).
- **`site_build/`:** The output directory for the built HTML site (ignored by Git).
- **Assets:** Images and static files are typically stored in `docs/assets/` or category-specific `img/` folders.

## CI/CD Workflow (GitLab)
The project uses a two-stage GitLab CI/CD pipeline:

### 1. Build Stage
- **Image:** `python:3.12`
- **Environment:** Configures local DNS and uses Iranian PyPI mirrors (`liara.ir`, `runflare.com`, `hyperclouds.ir`) to ensure connectivity.
- **Process:** Installs dependencies from `requirements.txt` and runs `mkdocs build`.
- **Output:** Generates artifacts in `site_build/`.

### 2. Deploy Stage
- **Image:** `ubuntu:latest`
- **Environment:** Configures local APT mirrors (`arvancloud.ir`) and installs `ansible`, `rsync`, and `openssh-client`.
- **SSH Auth:** Uses a Base64-encoded `$SSH_PRIVATE_KEY` for server access.
- **Process:** Executes the Ansible playbook to push changes to the production server.

## Ansible Deployment Flow
The deployment is orchestrated via `ansible/playbook.yml`:
- **Inventory:** Targets hosts in the `doc_servers` group (defined in `ansible/inventory.ini`).
- **Directory Management:** Ensures `/var/www/docs` exists on the target server.
- **Sync:** Copies the contents of `site_build/` to the web server's application directory.
- **Nginx Config:** Deploys an Nginx configuration from `ansible/templates/docs.j2` to `/etc/nginx/sites-available/docs` and links it to `sites-enabled`.
- **Handlers:** Restarts the Nginx service upon configuration or content updates.

## Gemini CLI Best Practices
- **Updates:** When adding new documentation categories, update `mkdocs.yml` if manual navigation is required (though `navigation.indexes` and `navigation.tabs` are enabled).
- **Mirrors:** Always respect the mirror configurations in CI/CD when suggesting changes to dependencies or build steps.
