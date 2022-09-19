# kubearchive

Save Kubernetes objects and logs off cluster.

This project has largely been inspired by [Tekton Results](https://github.com/tektoncd/results).

## What does this do?

`kubearchive` watches objects on your cluster, and stores its data in a database.
When properly configured, `kubearchive` can also store logs for that object's containers.

## Try It Out!

Follow the [installation instructions](docs/installation.md) to build and deploy Kubearchive.

## Reference

Use the [REST interface guide](docs/reference/rest-interface.md) to understand how to query and fetch data from Kuberarchive.

Refer to the [Architecture](docs/architecture/architecture.md) diagram to understand how Kubearchive works.
