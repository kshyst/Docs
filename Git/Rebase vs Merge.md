# Rebase vs Merge

## The Golden Rule Of Rebasing

> The golden rule of git rebase is to never use it on public branches.

![Golden Rule](https://wac-cdn.atlassian.com/dam/jcr:2908e0e6-f74b-4425-b5d2-f5eca8cfcd99/05%20Rebasing%20the%20main%20branch.svg?cdnVersion=2554)


> The rebase moves all of the commits in main onto the tip of feature. The problem is that this only happened in your repository. All of the other developers are still working with the original main. Since rebasing results in brand new commits, Git will think that your main branch’s history has diverged from everybody else’s.


> The only way to synchronize the two main branches is to merge them back together, resulting in an extra merge commit and two sets of commits that contain the same changes (the original ones, and the ones from your rebased branch). Needless to say, this is a very confusing situation.

> So, before you run git rebase, always ask yourself, “Is anyone else looking at this branch?” If the answer is yes, take your hands off the keyboard.


## Force pushing the new main branch

> If you try to push the rebased main branch back to a remote repository, Git will prevent you from doing so because it conflicts with the remote main branch. 

> Force Pushing the main overwrites the remote main branch to match the rebased one from your repository and makes things very confusing for the rest of your team. So, be very careful to use this command only when you know exactly what you’re doing.

## Merge-Base

> The merge-base is the common ancestor of the two branches. It is the commit that both branches were forked from.

```shell

git merge-base main feature
```

## Summary

> And that’s all you really need to know to start rebasing your branches. If you would prefer a clean, linear history free of unnecessary merge commits, you should reach for git rebase instead of git merge when integrating changes from another branch.

> On the other hand, if you want to preserve the complete history of your project and avoid the risk of re-writing public commits, you can stick with git merge. Either option is perfectly valid, but at least now you have the option of leveraging the benefits of git rebase.