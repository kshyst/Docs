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

