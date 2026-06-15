# Git Checkout

The `git checkout` command is used to switch branches or restore working tree files.

## Checking Out Specific Files
You can pull a specific file or directory from another branch without switching your entire working directory to that branch.

### From Another Branch
To checkout a specific file or folder from the `prod` branch:
```bash
git checkout prod -- bottle-connector
```
This will update the `bottle-connector` path in your current branch to match the version in `prod`.

### From a Baseline / Commit
To restore a file to its state at a specific commit or baseline:
```bash
git checkout <commit_hash> -- <file_path>
```

### Discarding Local Changes
If you want to discard local changes and return to the state of the current branch (HEAD):
```bash
git checkout HEAD -- <file_path>
```
