# kubearchive

Save Kubernetes objects and logs off cluster.

This project has largely been inspired by [Tekton Results](https://github.com/tektoncd/results).

## What does this do?

`kubearchive` watches objects on your cluster, and stores its data in a database.
When properly configured, `kubearchive` can also store logs for that object's containers.

## Installation

Run the following:

```sh
$ make deploy
```
