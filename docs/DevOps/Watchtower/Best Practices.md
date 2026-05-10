# Best Practices

While Watchtower simplifies container management, following best practices ensures your production environment remains stable.

## Use Proper Image Tags
Avoid using the `latest` tag in production. Instead, use specific version tags or semi-stable tags (e.g., `v1.2`, `v1.2-stable`). This allows you to control which major versions are updated automatically.

## Exclude Critical Containers
If some containers should never be updated automatically (e.g., a database where a minor version update might require a manual migration), you can explicitly exclude them.

### Method 1: Labels
If you are using `--label-enable`, simply don't add the label to critical containers.
Alternatively, use the disable label:
`com.centurylinklabs.watchtower.enable=false`

### Method 2: Command Line
Exclude specific containers when starting Watchtower:
```bash
docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --stop-signal SIGTERM my-critical-db
```

## Handle Graceful Shutdowns
Ensure your applications handle signals correctly. Watchtower sends a stop signal (default `SIGTERM`) to containers before stopping them.
- Use the `--stop-signal` flag if your application requires a different signal.
- Increase the timeout if your application needs more than 10 seconds to shut down:
  `--timeout 30s`

## Monitor Watchtower Logs
Always check Watchtower's logs to ensure updates are completing successfully.
```bash
docker logs -f watchtower
```

## Use Cleanup Regularly
Old Docker images can consume significant disk space. Always use the `--cleanup` flag to remove unused images after an update.

## Test in Staging First
Before deploying Watchtower to production, test it in a staging or development environment to ensure that your containers restart correctly and that no volume data is lost during the update process.
