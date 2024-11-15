# The fun and scary (for some) part of the workshop

Running on k8s!

As you might expect, Wasm can run on baremetal and on kubernetes. Checkout [https://wasmcloud.com/docs/deployment/deployment-guide/](wasmCloud's Deployment Guides) for your favorite places to run.

The scripts for this workshop are taken from [wasmcloud-contrib](https://github.com/wasmCloud/wasmCloud-contrib/tree/main/deploy/k8s/kustomize) and added here so that workshop folks only need to look to one place.

## wasmCloud deployment with kustomize

This example contains configuration files and scripts to deploy wasmCloud on Kubernetes using Kustomize.
Kustomize is a configuration management tool that allows you to customize raw, template-free YAML files for multiple purposes, leaving the original YAML untouched and usable as is.

## Dependencies

To deploy wasmCloud on a local kind cluster with ingress, you will need to have the following prerequisites installed:

1. [kind](https://kind.sigs.k8s.io/docs/user/quick-start/): Kind is a tool for running local Kubernetes clusters using Docker container "nodes". Follow the quick start guide to install kind.
2. [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/): Kustomize is a powerful configuration management tool for Kubernetes that allows you to customize and manage YAML configurations without using templates. It is integrated into kubectl and can be used to manage different configurations for different environments seamlessly. To install kubectl, which includes Kustomize, follow the [installation guide](https://kubectl.docs.kubernetes.io/installation/kustomize/).

Once you have installed these prerequisites, you will be ready to deploy wasmCloud on your local kind cluster with Ingress-enabled.

## Deployment

```bash
./scripts/kind-cluster.sh
```

This script runs through the following steps:

1. Create a Kubernetes cluster using `kind` with the configuration specified in `deploy/kind/cluster.yaml`.
2. Deploy [NGNIX ingress controller](https://github.com/kubernetes/ingress-nginx)
3. Deploy Nats
4. Build and apply the Kubernetes manifests using `kustomize` with Helm enabled, using the configuration located at `./deploy/dev`.
5. Apply the wasmCloud host configuration located at `./deploy/dev/wasmcloud-host-config.yaml`.

### Using wash locally

In order to run wash locally, we need a connection to NATS.

```bash
kubectl port-forward svc/nats 4222:4222 4223:4223
```

In another terminal tab:

```bash
wash get inventory


  Host ID                                                    Friendly name
  NDJXJT6UVJPXFKQCX6RXKGV4PCZWCEAQLYEROOBUP5SKJ544C5S3SWRM   green-moon-6757

  Host labels
  hostcore.osfamily                                          unix
  cluster-type                                               kind
  hostcore.os                                                linux
  hostcore.arch                                              aarch64
  kubernetes                                                 true

  Component ID                                               Name                   Max count
  rust_hello_world-http_component                            http-hello-world       1

  Provider ID                                                Name
  path_based_routing-httpserver                              http-server-provider
```