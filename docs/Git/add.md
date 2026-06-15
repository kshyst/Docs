# Git Add

The `git add` command adds a change in the working directory to the staging area.

## Staging Changes

### Add All Changes
To stage all changes in the entire repository, including new files, modified files, and **deleted files**:
```bash
git add -A
```
*(Equivalent to `git add --all`)*

## Comparison of Add Flags

| Command | New Files | Modified Files | Deleted Files |
| :--- | :---: | :---: | :---: |
| `git add .` | Yes | Yes | Yes (in modern Git) |
| `git add -u` | No | Yes | Yes |
| `git add -A` | Yes | Yes | Yes |

- **`git add .`**: Stages changes in the current directory and its subdirectories.
- **`git add -u`**: Stages modified and deleted files, but **not** new (untracked) files.
- **`git add -A`**: Stages everything (new, modified, and deleted) across the entire repository.
