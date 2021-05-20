# Installation Manifests

Following commands will install Argo EventBus controller into namespace
`argo-eventbus`.

```sh
kubectl create ns argo-eventbus

kubectl apply -f ./install.yaml
```

If you desire to install it into a different namespace, change the namespace in
[cluster-install/kustomization.yaml](cluster-install/kustomization.yaml), run
command below to regenerate install.yaml before doing `kubectl apply`.

```sh
make manifests
```

## Kustomize

You can use `./cluster-install` Kustomize remote base, or use `./base` as the
remote base and give a namespace you desired in your kustomization.yaml.
