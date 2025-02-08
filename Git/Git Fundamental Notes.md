# Some Notes About Git

### commit -a
> Commit a snapshot of all changes in the working directory. This only includes modifications to tracked files (those that have been added with git add at some point in their history).

### commit --amend
> will modify the last commit. Instead of creating a new commit, staged changes will be added to the previous commit. This command will open up the system's configured text editor and prompt to change the previously specified commit message.

## Log

> ou can view all commits across all branches by executing :

```shell

git log --branches=*
```

## diff

> Diffing is a function that takes two input data sets and outputs the changes between them.


#### example of diffing 2 commits
```shell

git diff HEAD^ HEAD
```

> this diffs the last commit with the one before it.
> HEAD^ is the commit before the last commit.

### log

```shell

diff --git a/Git/Git Fundamental Notes.md b/Git/Git Fundamental Notes.md
new file mode 100644
index 0000000..d3f70b9
--- /dev/null
+++ b/Git/Git Fundamental Notes.md      
@@ -0,0 +1,12 @@
+# Some Notes About Git
+
+### commit -a
+> Commit a snapshot of all changes in the working directory. This only includes modifications to tracked files (those that have been added with git add at some point in their history).
+
+### commit --amend
+> will modify the last commit. Instead of creating a new commit, staged changes will be added to the previous commit. This command will open up the system's configured text editor and prompt to change the previously specified c
ommit message.
+
+
+### diff
+
+> Diffing is a function that takes two input data sets and outputs the changes between them.
\ No newline at end of file
diff --git a/Python/VirtualEnvironments.md b/Python/VirtualEnvironments.md
index 84a0cc7..df5679b 100644
--- a/Python/VirtualEnviro
```

> 1- Comparison input

```shell

diff --git a/Git/Git Fundamental Notes.md b/Git/Git Fundamental Notes.md
```

> 2- Metadata

```shell
index 0000000..d3f70b9
```

> 3- Markers for changes

```shell

--- /dev/null
+++ b/Git/Git Fundamental Notes.md  
```

> 4- Diff chunks

```shell

@@ -0,0 +1,12 @@
+# Some Notes About Git
+
+### commit -a
+> Commit a snapshot of all changes in the working directory. This only includes modifications to tracked files (those that have been added with git add at some point in their history).
+
+### commit --amend
```

## Stash

> The git stash command takes your uncommitted changes (both staged and unstaged), saves them away for later use, and then reverts them from your working copy. 

```shell
git stash
git stash list
git stash apply
git stash drop
git stash pop
```

### what will be stashed and what not

> will stash
> - changes that have been added to your index (staged changes)
> - changes made to files that are currently tracked by Git (unstaged changes)

> will not stash
> - new files in your working directory that Git has not yet started tracking
> - files that have been ignored by Git

> add -u for stashing untracked files

```shell

git stash -u
```
> add -a for stashing all files

```shell

git stash -a
```

### Popping Stashes

> by default it works like a stack
> but you can specify which stash to pop

```shell
git stash pop stash@{2}
```

### show

> shows the changes in a stash

```shell
git stash show -p stash@{2}
```

### partial stash

> you can stash only some files


> - / for seaching a hunk by regex
> - ? help
> - q quit
> - y stage this hunk
> - n do not stage this hunk
> - s split the hunk

```shell
git stash -p
```

### Create a branch from stash

> If the changes on your branch diverge from the changes in your stash, you may run into conflicts when popping or applying your stash. Instead, you can use git stash branch to create a new branch to apply your stashed changes to:

```shell

git stash branch add-stylesheet stash@{1}
Switched to a new branch 'add-stylesheet'
On branch add-stylesheet
Changes to be committed:

    new file:   style.css

Changes not staged for commit:

    modified:   index.html

Dropped refs/stash@{1} (32b3aa1d185dfe6d57b3c3cc3b32cbf3e380cc6a)
``` 

## Tags

> Tags are a way to mark specific points in a repository’s history as being important. Used for versioning basically.

### Annotated Tags

> Annotated tags are stored as full objects in the Git database. They’re checksummed; contain the tagger name, email, and date; have a tagging message; and can be signed and verified with GNU Privacy Guard (GPG).

```shell
git tag -a v1.0 -m "version 1.0"
```

### Lightweight Tags

> Lightweight tags are just pointers to specific commits. They are created with the -l flag.

```shell
git tag v1.0
```

### Listing Tags

```shell
git tag
```

### Tagging a Specific Commit

```shell

git tag -a v1.2 9fceb02
```

### Pushing Tags

```shell

git push origin v1.2
```

### Deleting Tags

```shell

git tag -d v1.2
```

### Checkout a Tag

```shell

git checkout v1.2
```

## Blame

> The git blame command is a versatile troubleshooting utility that has extensive usage options. The high-level function of git blame is the display of author metadata attached to specific committed lines in a file. This is used to examine specific points of a file's history and get context as to who the last author was that modified the line. This is used to explore the history of specific code and answer questions about what, how, and why the code was added to a repository.

```shell

git blame README.md
    82496ea3 (kevzettler     2018-02-28 13:37:02 -0800  1) # Git Blame example
    82496ea3 (kevzettler     2018-02-28 13:37:02 -0800  2)
    89feb84d (Albert So      2018-03-01 00:54:03 +0000  3) This repository is an example of a project with multiple contributors making commits.
    82496ea3 (kevzettler     2018-02-28 13:37:02 -0800  4)
    82496ea3 (kevzettler     2018-02-28 13:37:02 -0800  5) The repo use used elsewhere to demonstrate `git blame`
    82496ea3 (kevzettler     2018-02-28 13:37:02 -0800  6)
    89feb84d (Albert So      2018-03-01 00:54:03 +0000  7) Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod TEMPOR incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum
    89feb84d (Albert So      2018-03-01 00:54:03 +0000  8)
    eb06faed (Juni Mukherjee 2018-03-01 19:53:23 +0000  9) Annotates each line in the given file with information from the revision which last modified the line. Optionally, start annotating from the given revision.
    eb06faed (Juni Mukherjee 2018-03-01 19:53:23 +0000 10)
    548dabed (Juni Mukherjee 2018-03-01 19:55:15 +0000 11) Creating a line to support documentation needs for git blame.
    548dabed (Juni Mukherjee 2018-03-01 19:55:15 +0000 12)
    548dabed (Juni Mukherjee 2018-03-01 19:55:15 +0000 13) Also, it is important to have a few of these commits to clearly reflect the who, the what and the when. This will help Kev get good screenshots when he runs the git blame on this README.
```

## Clean

> The git clean command is a Git utility tool that is used to remove untracked files from the working directory. This command is useful for removing unwanted files that are not being tracked by Git. It is also used to remove directories that are not being tracked by Git.

```shell

git clean -n

git clean -f
git clean --force
git clean -df  # remove directories that are not being tracked by Git
git clean -dn # dry run
git clean -xf # remove ignored files

```

### Interactive Clean

> The git clean command can be used with the -i or --interactive flag to interactively remove untracked files from the working directory. This flag allows you to selectively choose which files to remove.

```

git clean -di
Would remove the following items:
  untracked_dir/  untracked_file
*** Commands ***
    1: clean                2: filter by pattern    3: select by numbers    4: ask each             5: quit                 6: help
What now>
```