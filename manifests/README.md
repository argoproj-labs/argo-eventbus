# Installation Manifests

Following commands will install Argo EventBus controller into namespace
`argo-eventbus`.

```sh
kubectl create ns argo-eventbus

# For cluster scope installation
kubectl apply -f ./install.yaml

# For namespace scope installation
kubectl apply -f ./namespace-install.yaml
```

If you desire to install it into a different namespace, change the namespace in
[cluster-install/kustomization.yaml](cluster-install/kustomization.yaml) or
[namespace-install/kustomization.yaml](namespace-install/kustomization.yaml),
run command below to regenerate manifests before running `kubectl apply`.

```sh
make manifests
```

## Kustomize

You can use `./cluster-install` or `./namespace-install` as Kustomize remote
base. If you want to install in a different namespace, give a namespace you
desired in your kustomization.yaml.
