# Basic Usage

Watchtower is highly configurable through command-line arguments and environment variables. Here are the most common ways to control its behavior.

## Common Flags

### Run Once
If you want Watchtower to check for updates, perform them, and then immediately exit, use the `--run-once` (or `-run-once`) flag. This is useful for manual maintenance tasks or cron jobs.
```bash
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --run-once
```

### Cleanup
By default, Docker keeps old images on the disk after an update. Use the `--cleanup` (or `-c`) flag to remove old images after successfully updating a container.
```bash
docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --cleanup
```

### Polling Interval
You can define how often Watchtower checks for updates using the `--interval` flag (specified in seconds). The default is 86400 seconds (24 hours).
```bash
# Check every 30 minutes
docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --interval 1800
```

## Monitoring Specific Containers
By default, Watchtower monitors all running containers. You can restrict it to specific containers by listing their names as arguments:
```bash
docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower nginx redis
```

## Environment Variables
Most flags have corresponding environment variables. For example:
- `WATCHTOWER_CLEANUP=true`
- `WATCHTOWER_POLL_INTERVAL=1800`
- `WATCHTOWER_RUN_ONCE=true`
