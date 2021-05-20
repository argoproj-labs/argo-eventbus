# Argo EventBus Install Manifests

If installing with `kubectl apply -f https://...`, remember to use the link to
the file's raw version. Otherwise you will get
`mapping values are not allowed in this context`.

Manifests expect the namespace `argo-eventbus` to exist. If you desire to deploy
it into a different namespace, change the namespace in
[cluster-install](cluster-install/kustomization.yaml).

Cluster-wide install:

```sh
kubectl create ns argo-eventbus

kubectl apply -f ./install.yaml
```

## Kustomize

You can use `./cluster-install` Kustomize remote base, or use `./base` as the
remote base and give a namespace in your kustomization.yaml.
