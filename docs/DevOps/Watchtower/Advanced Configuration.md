# Advanced Configuration

Watchtower provides advanced features for complex environments and specific needs.

## Label-Based Filtering
Instead of monitoring all containers, you can use labels to opt-in specific containers.

1. Start Watchtower with the `--label-enable` flag:
   ```bash
   docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --label-enable
   ```
2. Add the `com.centurylinklabs.watchtower.enable=true` label to the containers you want to update:
   ```yaml
   services:
     myapp:
       image: myorg/myapp:latest
       labels:
         - com.centurylinklabs.watchtower.enable=true
   ```

## Scheduling with Cron
If you prefer updates to happen at specific times (e.g., at night), use the `WATCHTOWER_SCHEDULE` environment variable with a Cron expression.
```bash
# Update every day at 3 AM
docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock -e WATCHTOWER_SCHEDULE="0 0 3 * * *" containrrr/watchtower
```

## Notifications
Watchtower can send notifications when updates are performed. Supported services include Slack, Discord, Email, and Gotify.

### Slack Example
```bash
-e WATCHTOWER_NOTIFICATIONS=slack \
-e WATCHTOWER_NOTIFICATION_SLACK_HOOK_URL="https://hooks.slack.com/services/..."
```

### Discord Example
```bash
-e WATCHTOWER_NOTIFICATIONS=discord \
-e WATCHTOWER_NOTIFICATION_DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/..."
```

## Private Registries
To update images from private registries, Watchtower needs credentials. The easiest way is to mount your local Docker config file:
```bash
docker run -d \
  --name watchtower \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /home/user/.docker/config.json:/config.json \
  containrrr/watchtower
```
Alternatively, use environment variables:
- `REPO_USER=myuser`
- `REPO_PASS=mypassword`

## Environment Variables Summary
- `WATCHTOWER_SCHEDULE`: Cron expression for updates.
- `WATCHTOWER_NOTIFICATIONS`: Type of notification service.
- `WATCHTOWER_LABEL_ENABLE`: Only update containers with the specific label.
- `WATCHTOWER_TIMEOUT`: Timeout for stopping containers (default 10s).
