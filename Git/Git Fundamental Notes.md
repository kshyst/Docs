# Some Notes About Git

### commit -a
> Commit a snapshot of all changes in the working directory. This only includes modifications to tracked files (those that have been added with git add at some point in their history).

### commit --amend
> will modify the last commit. Instead of creating a new commit, staged changes will be added to the previous commit. This command will open up the system's configured text editor and prompt to change the previously specified commit message.


### diff

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