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
    # deploy app-a
    # deploy app-b
else
    apps=$(git-path-changes matches -m app-a -m app-b $COMMIT1 HEAD)
    for app in apps
    do
        # deploy app
    done
fi
```

# Installing

```bash
go get github.com/jamiemccrindle/gitpathchanges
```

# Commands

## match

List which matching files and directories have changed between 2 commits.

```
NAME:
   git-path-changes match - List which matching files and directories have changed between 2 commits

USAGE:
   git-path-changes match [command options] <commit1> <commit2>

OPTIONS:
   --path value, -p value   The path to your git repository (default: ".")
   --match value, -m value  The files or directories to check for changes
   --help, -h               show help (default: false)
```

## has-matches

Return success if matching files have changed between 2 commits otherwise fail. Useful to test whether a file or directory has changed between commits.

```
NAME:
   git-path-changes has-matches - Return success if matching files have changed between 2 commits otherwise fail

USAGE:
   git-path-changes has-matches [command options] <commit1> <commit2>

OPTIONS:
   --path value, -p value   The path to your git repository (default: ".")
   --match value, -m value  The files or directories to check for changes
   --help, -h               show help (default: false)
```

## directories

List all files that have changed between 2 commits.

```
NAME:
   git-path-changes directories - List all files that have changed between 2 commits

USAGE:
   git-path-changes directories [command options] <commit1> <commit2>

OPTIONS:
   --path value, -p value  The path to your git repository (default: ".")
   --help, -h              show help (default: false)
```