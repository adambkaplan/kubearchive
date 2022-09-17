# kubearchive Installation

## Prerequisites

1. Install a container engine onto your machine, such as [docker](https://www.docker.com/) or [podman](https://podman.io).
2. Ensure you have admin access to a Kubernetes cluster configured with an [Ingress controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/).
3. Clone this repository:

   ```sh
   $ git clone https://github.com/adambkaplan/kubearchive.git
   ```

## Build and Push the Image

First, ensure your container engine is logged into container registry that is accessible to your Kubernetes cluster.

Next, build and push the image, using the `IMG` variable to specify where to push the image.

```sh
$ make container-build IMG=myregistry.io/myuser/kubearchive:latest
$ make container-push IMG=myregistry.io/myuser/kubearchive:latest
```

## Configure an Ingress Overlay

The deployment manifests include an Ingress for the `localhost` host domain, which is suitable for local development with KinD clusters.
When deploying on a remote cluster, you should provide a Kustomize overlay as follows:

1. Copy the overlays example to a directory under `config/overlays`:

   ```sh
   $ mkdir -p config/overlays
   $ cp -r config/examples/overlays config/overlays/test
   ```

2. Set the desired value for your host domain in the `apiserver-ingress-patch.yaml` file.

## Deploy kubearchive

Use the `deploy` make target to install `kuberarchive` on your cluster, using the `IMG` variable to specify the image you built previously.

```sh
$ make deploy IMG=myregistry.io/myuser/kubearchive:latest
```

To deploy using an ingress overlay, use the `OVERLAY` to specify the directory containing your overlays:

```sh
$ make deploy IMG=myregistry.io/myuser/kubearchive:latest OVERLAY=config/overlays/test
```
