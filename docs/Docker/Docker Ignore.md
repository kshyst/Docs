# Docker Ignore

The `.dockerignore` file is used to exclude files and directories from the build context when creating a Docker image. This helps reduce the build time and the size of the final image by preventing unnecessary files from being sent to the Docker daemon.

## Why use .dockerignore?

- **Reduced Build Context Size**: Faster uploads to the Docker daemon.
- **Smaller Image Size**: Prevents sensitive or unnecessary files (like `.git`, `node_modules`, or local logs) from being copied into the image.
- **Security**: Ensures that secrets, environment files, and local configuration don't accidentally end up in the production image.

## Example .dockerignore File

```text
# Dependency directories
node_modules/
vendor/

# Log files
*.log

# Local configuration
*.config
.env

# OS-specific files
.DS_Store
Thumbs.db

# Test and Documentation
test/
*.md

# Exceptions
!Dockerfile
!package.json
```

## Pattern Matching Rules

- `*`: Matches zero or more characters.
- `?`: Matches a single character.
- `**`: Matches any number of directories.
- `!`: An exclusion prefix that allows certain files to be included even if they match a previous pattern.