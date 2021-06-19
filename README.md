# Check for path changes between commits

Check which paths have changes between 2 git commits. This can be useful for optimising builds or deployments where
you can conditionally build or deploy based on what has changed between 2 commits.

For example, say you have this folder structure:

```
infra/
app-a/
app-b/
```

You may only want to redeploy everything if your `infra` folder changes but only `app-a` or `app-b` if they change.

You can use this utility to check for those changes e.g.

```bash
COMMIT=<commit_of_live_deployment>
if git-path-changes has-matches -m infra $COMMIT1 HEAD
then
    # deploy infra
    # deploy app1
    # deploy app2
else
    apps=$(git-path-changes matches -m app-a -m app-b $COMMIT1 HEAD)
    for app in apps
    do
        # deploy app
    done
fi
```
