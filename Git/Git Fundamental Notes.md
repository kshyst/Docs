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

## Revert

> The git revert command can be considered an 'undo' type command, however, it is not a traditional undo operation. Instead of removing the commit from the project history, it figures out how to invert the changes introduced by the commit and appends a new commit with the resulting inverse content. This prevents Git from losing history, which is important for the integrity of your revision history and for reliable collaboration.

### How it works exactly

> The git revert command is used for undoing changes to a repository's commit history. Other 'undo' commands like, git checkout and git reset, move the HEAD and branch ref pointers to a specified commit. Git revert also takes a specified commit, however, git revert does not move ref pointers to this commit. A revert operation will take the specified commit, inverse the changes from that commit, and create a new "revert commit". The ref pointers are then updated to point at the new revert commit making it the tip of the branch.

### Common options

> -n, --no-commit: This option will revert the changes but will not create a new commit. This is useful for reverting changes before committing them.
> -e, --edit: This option will open the commit message in the default text editor. This allows you to modify the commit message before committing the revert.
> --no-edit: This option will revert the changes and create a new commit with the default commit message. This is the default behavior.

### Revert vs Reset

> Reverting has two important advantages over resetting. First, it doesn’t change the project history, which makes it a “safe” operation for commits that have already been published to a shared repository. For details about why altering shared history is dangerous.


> Second, git revert is able to target an individual commit at an arbitrary point in the history, whereas git reset can only work backward from the current commit. For example, if you wanted to undo an old commit with git reset, you would have to remove all of the commits that occurred after the target commit, remove it, then re-commit all of the subsequent commits. Needless to say, this is not an elegant undo solution.



## Reset

> The git reset command is a complex and versatile tool for undoing changes. It has three primary forms of invocation. These forms correspond to command line arguments --soft, --mixed, --hard. The three arguments each correspond to Git's three internal state management mechanism's, The Commit Tree (HEAD), The Staging Index, and The Working Directory.


### --hard

> This is the most direct, DANGEROUS, and frequently used option. When passed --hard The Commit History ref pointers are updated to the specified commit. Then, the Staging Index and Working Directory are reset to match that of the specified commit. Any previously pending changes to the Staging Index and the Working Directory gets reset to match the state of the Commit Tree. This means any pending work that was hanging out in the Staging Index and Working Directory will be lost.


### --mixed

> This is the default option. When passed --mixed The Commit History ref pointers are updated to the specified commit. Then, the Staging Index is reset to match that of the specified commit. The Working Directory is not affected. This means any pending changes in the Working Directory are left alone. The changes that were in the Staging Index are reset to match the specified commit.


### --soft

> When passed --soft The Commit History ref pointers are updated to the specified commit. The Staging Index and Working Directory are not affected. This means any pending changes in the Staging Index and Working Directory are left alone. The changes that were in the Commit Tree are reset to match the specified commit.


### differences
> git reset --mixed: Resets the HEAD and updates the staging area, but leaves the working directory unchanged.

>git reset --soft: Only resets the HEAD, leaving both the staging area and the working directory unchanged.

>git reset --hard: Resets the HEAD, updates the staging area, and resets the working directory to match the specified commit.


## RM

> The git rm command can be used to remove individual files or a collection of files. The primary function of git rm is to remove tracked files from the Git index. Additionally, git rm can be used to remove files from both the staging index and the working directory. There is no option to remove a file from only the working directory. The files being operated on must be identical to the files in the current HEAD. If there is a discrepancy between the HEAD version of a file and the staging index or working tree version, Git will block the removal. This block is a safety mechanism to prevent removal of in-progress changes.

### Undoing a git rm

> Executing git rm is not a permanent update. The command will update the staging index and the working directory. These changes will not be persisted until a new commit is created and the changes are added to the commit history. This means that the changes here can be "undone" using common Git commands.

```shell

git reset HEAD <file>

git checkout -- <file>
```

## Branches

> Git branches are effectively a pointer to a snapshot of your changes. When you want to add a new feature or fix a bug—no matter how big or how small—you spawn a new branch to encapsulate your changes. This makes it harder for unstable code to get merged into the main code base, and it gives you the chance to clean up your future's history before merging it into the main branch.

> Deleting a branch on shared repository , if you delete a branch on the local it wont be deleted on the shared repository because its now a different branch like b1 and origin/b1

```shell

git push origin --delete crazy-experiment
```

## Rebase

> The git rebase command is used to reapply commits on top of another base tip. This is useful for local branches that are not shared or for rewriting the commit history of a shared branch. The git rebase command is an alternative to the git merge command.

### Interactive Rebase

> The git rebase command can be used with the -i or --interactive flag to interactively reapply commits on top of another base tip. This flag allows you to selectively choose which commits to reapply.

```shell

git rebase -i HEAD~3
```

### --onto

```shell

git rebase --onto main feature1 feature2
```
> from this 
```shell
   o---o---o---o---o  main
        \
         o---o---o---o---o  featureA
              \
               o---o---o  featureB
```

> to this

```shell
                      o---o---o  featureB
                     /
    o---o---o---o---o  main
     \
      o---o---o---o---o  featureA
```


## Reflog

> The git reflog command is a Git utility tool that is used to manage the reference logs in a Git repository. The reference logs are used to store the history of the HEAD and the branch tips. The git reflog command is used to view the reference logs and to manage the reference logs in a Git repository.

```shell

git reflog
```

### How to view reflog for specific branch

```shell

git reflog master
```

### How to undo a git rebase using reflog

> The git reflog command can be used to undo a git rebase operation. The reflog command will display the
> history of the HEAD and the branch tips. The reflog command can be used to find the commit hash of the
> commit before the rebase operation. The git reset command can then be used to reset the HEAD to the commit
> before the rebase operation.

```shell

git reflog
git reset --hard HEAD@{1}
```


## Remote

> The git remote command is a Git utility used to manage remote repositories. The git remote command is used to create, view, and delete connections to other repositories. The git remote command is used to manage the remote repositories that are associated with a Git repository.


### Prune

> deletes any local branch that isn't available in remote repository

```shell

git remote prune origin
```

### Show

> high level information about the remote repository

```shell

git remote show origin
```


## Fetch

> The git fetch command is used to retrieve commits from a remote repository's branch

```shell

git fetch origin
```

### Fetching a specific branch

```shell

git fetch origin master
```

### Fetching all branches

```shell

git fetch --all
```

## Pull

> The git pull command is used to fetch and download content from a remote repository and immediately update the local repository to match that content. The git pull command is actually a combination of two other commands, git fetch followed by git merge. In the first stage of operation git pull will execute a git fetch scoped to the local branch that HEAD is pointed at. Once the content is downloaded, git pull will enter a merge workflow. A new merge commit will be-created and HEAD updated to point at the new commit.

```shell

git pull origin master
```

### Pulling while having new branches or different commits on master and same branches

> in git config you can choose from 3 type of behaviour to use when pulling
> 
> - rebase
> - merge
> - ff or fast-forward

```shell

git config --global pull.rebase true

git config --global pull.ff only

git config --global pull.merge true
#or
git config --global pull.rebase false
```


## Push

> The git push command is used to upload local repository content to a remote repository. Pushing is how you transfer commits from your local repository to a remote repo. It's the counterpart to git fetch, but whereas fetching imports commits to local branches, pushing exports commits to remote branches.

```shell

git push origin master
```

> to push all branches and tags use :

```shell

git push origin --all
```

### Pushing to bare repositories

> A bare repository is a repository that does not have a working directory. It is used to store the history of a project without the need for a working directory. Bare repositories are typically used as a central repository for a project that is shared among multiple developers.


### Force pushing

> Git prevents you from overwriting the central repository’s history by refusing push requests when they result in a non-fast-forward merge. So, if the remote history has diverged from your history, you need to pull the remote branch and merge it into your local one, then try pushing again.
> 
> he --force flag overrides this behavior and makes the remote repository’s branch match your local one, deleting any upstream changes that may have occurred since you last pulled.

```shell

git push origin master --force
```

### Push deleting a branch

```shell

git branch -D branch_name
git push origin :branch_name
```

## Squash

> Git squash is the process of taking a series of commits and merge them into a single commit. This technique is especially useful when a feature branch has numerous small, incremental commits that clutter the commit history. By squashing these commits, you can present a more polished and cohesive history to your collaborators.

> merge the changes in feature branch into current branch

```shell

git merge --squash feature-branch
```

> Squashing using rebase 

```shell

git rebase -i HEAD~3
```

> this code will open an editor with the last 3 commits

```shell

pick 1a2b3c4 Commit message 1
pick 5e6f7g8 Commit message 2
pick 9h0i1j2 Commit message 3
```

> change the commits to squash

```shell

pick 1a2b3c4 Commit message 1
squash 5e6f7g8 Commit message 2
squash 9h0i1j2 Commit message 3
```

> save and close the editor

> this will open another editor to edit the commit message

### Squashing from detached head

> If you are in a detached HEAD state, you can squash commits using the git rebase command. The git rebase command is used to reapply commits on top of another base tip. This is useful for local branches that are not shared or for rewriting the commit history of a shared branch.
> the new commit wont be referenced by any branch and is on a detached head state

```shell

## Detached Head

> A detached HEAD in Git is the state a repository is in when the HEAD points to a specific commit instead of a branch. This means that any changes made in a detached HEAD state will not be referenced by any branch and will be lost if the HEAD moves to another commit.

> To save the commits in a detached HEAD state, you can create a new branch that points to the current commit. This will allow you to reference the commits later.

```shell

git checkout -b new-branch
#or
git switch -c new-branch
```

## About Merge

> The git merge command is used to combine changes from one branch into another. This command is typically used to merge changes from a feature branch into the main branch. The git merge command is a powerful tool that can be used to combine changes from different branches.

> - when merging branch x into branch y it means that = git merge = while you are checked out on branch y
> - the head of the branch y will shift into the merge commit but the head of branch x will not change

### Fast forward merge

> A fast-forward merge is a type of merge operation in Git that allows you to move the branch pointer to a new commit without creating a new merge commit. This type of merge is possible when the branch being merged has not diverged from the branch you are merging into. In this case, Git can simply move the branch pointer forward to the new commit.

![ff merge](img/ffmerge.png)

> for not fast forwarding while fast forwarding is possible use :

```shell

git merge --no-ff feature-branch
```

### 3-way merge

> A 3-way merge is a type of merge operation in Git that creates a new merge commit to combine changes from two branches. This type of merge is used when the branches being merged have diverged from each other. In a 3-way merge, Git compares the two branches and the common ancestor commit to create a new merge commit that combines the changes from both branches.

```shell

git merge feature-branch
```

### Merge conflicts

> If the two branches you're trying to merge both changed the same part of the same file, Git won't be able to figure out which version to use. When such a situation occurs, it stops right before the merge commit so that you can resolve the conflicts manually.

>The great part of Git's merging process is that it uses the familiar edit/stage/commit workflow to resolve merge conflicts. When you encounter a merge conflict, running the git status command shows you which files need to be resolved.

#### there are 2 types of conflicts

> - failed to start the merging process :
>> A merge will fail to start when Git sees there are changes in either the working directory or staging area of the current project. Git fails to start the merge because these pending changes could be written over by the commits that are being merged in. When this happens, it is not because of conflicts with other developer's, but conflicts with pending local changes. The local state will need to be stabilized using git stash, git checkout, git commit or git reset.
> - Git fails during the merge
>> A failure DURING a merge indicates a conflict between the current local branch and the branch being merged. This indicates a conflict with another developers code. Git will do its best to merge the files but will leave things for you to resolve manually in the conflicted files.

#### Tools to use when merge fails to start

> - git checkout # to discard changes in working directory
> - git reset # to discard changes in staging area

#### Tools to use when merge fails during the merge

> - git mergetool # to open a visual merge tool
> - git diff # to see the differences between the files
> - git merge --abort # to abort the merge
> - git add # to stage the resolved files
> - git commit # to commit the resolved files
> - git reset # to unstage the resolved files

### Merge Strategies

> git merge will automatically select a merge strategy unless explicitly specified. The git merge and git pull commands can be passed an -s (strategy) option. The -s option can be appended with the name of the desired merge strategy. If not explicitly specified, Git will select the most appropriate merge strategy based on the provided branches. The following is a list of the available merge strategies.

> - recursive: This operates on two heads. Recursive is the default merge strategy when pulling or merging one branch. Additionally this can detect and handle merges involving renames, but currently cannot make use of detected copies. This is the default merge strategy when pulling or merging one branch.

> - resolve:This can only resolve two heads using a 3-way merge algorithm. It tries to carefully detect cris-cross merge ambiguities and is considered generally safe and fast.

> - octopus: The default merge strategy for more than two heads. When more than one branch is passed octopus is automatically engaged. If a merge has conflicts that need manual resolution octopus will refuse the merge attempt. It is primarily used for bundling similar feature branch heads together.

> - ours: The Ours strategy operates on multiple N number of branches. The output merge result is always that of the current branch HEAD. The "ours" term implies the preference effectively ignoring all changes from all other branches. It is intended to be used to combine history of similar feature branches.

> - subtree: This is an extension of the recursive strategy. When merging A and B, if B is a child subtree of A, B is first updated to reflect the tree structure of A, This update is also done to the common ancestor tree that is shared between A and B.



## Cherry-picking

> git cherry-pick is a powerful command that enables arbitrary Git commits to be picked by reference and appended to the current working HEAD. Cherry picking is the act of picking a commit from a branch and applying it to another. git cherry-pick can be useful for undoing changes. For example, say a commit is accidently made to the wrong branch. You can switch to the correct branch and cherry-pick the commit to where it should belong.

### Why use cherry-picking

> git cherry-pick is a useful tool but not always a best practice. Cherry picking can cause duplicate commits and many scenarios where cherry picking would work, traditional merges are preferred instead.

#### Team collaboration

> Often times a team will find individual members working in or around the same code. Maybe a new product feature has a backend and frontend component. There may be some shared code between to two product sectors. Maybe the backend developer creates a data structure that the frontend will also need to utilize. The frontend developer could use git cherry-pick to pick the commit in which this hypothetical data structure was created. This pick would enable the frontend developer to continue progress on their side of the project.

#### Bug hotfixes

> When a bug is discovered it is important to deliver a fix to end users as quickly as possible. For an example scenario,say a developer has started work on a new feature. During that new feature development they identify a pre-existing bug. The developer creates an explicit commit patching this bug. This new patch commit can be cherry-picked directly to the main branch to fix the bug before it effects more users.

#### Undoing changes and restoring lost commits

> Sometimes a feature branch may go stale and not get merged into main. Sometimes a pull request might get closed without merging. Git never loses those commits and through commands like git log and git reflog they can be found and cherry picked back to life.

### cherry pick with --no-commit 

> The --no-commit option will execute the cherry pick but instead of making a new commit it will move the contents of the target commit into the working directory of the current branch.
