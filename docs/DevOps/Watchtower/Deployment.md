# Deployment

Watchtower is most effective when deployed as a background service. It requires access to the Docker socket to monitor and manage other containers.

## Docker CLI
To quickly start Watchtower using the command line:

```bash
docker run -d \
  --name watchtower \
  -v /var/run/docker.sock:/var/run/docker.sock \
  containrrr/watchtower
```

## Docker Compose (Recommended)
Using Docker Compose is the preferred method as it allows for easier configuration management and persistence.

### Basic Setup
```yaml
services:
  watchtower:
    image: containrrr/watchtower
    container_name: watchtower
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - WATCHTOWER_CLEANUP=true
      - WATCHTOWER_POLL_INTERVAL=3600
    restart: always
```

### Complete Example with Notifications
```yaml
services:
  watchtower:
    image: containrrr/watchtower
    container_name: watchtower
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - WATCHTOWER_CLEANUP=true
      - WATCHTOWER_NOTIFICATIONS=slack
      - WATCHTOWER_NOTIFICATION_SLACK_HOOK_URL=https://hooks.slack.com/services/XXX/YYY/ZZZ
      - WATCHTOWER_NOTIFICATION_SLACK_IDENTIFIER=watchtower-prod
    restart: always
```

## Security Note
Mounting `/var/run/docker.sock` gives Watchtower full control over your Docker daemon. Ensure that only trusted users have access to the host and that you are using the official `containrrr/watchtower` image.
