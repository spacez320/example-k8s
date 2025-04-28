example-k8s
===========

This produces a very simple Go webserver which is meant to be useful as an example application for
testing a new Kubernetes server.

Usage
-----

You can run things locally (although the point is to run it in Kubernetes).

```
go install src/example
bin/example
```

```
docker build -f build/package/Dockerfile .
```

Beyond that, there are several ways to install this into Kubernetes.

### Kubectl

```
kubectl create namespace example
kubectl run --image spacez320/example-k8s:latest example --namespace example
```

### Helm

```
helm upgrade --create-namespace --install --namespace example example ./build/package/helm
```

### Kustomize

```
kubectl create namespace example
kubectl apply --kustomize ./build/package/kustomize/overlays/default/ --namespace example
```

Development
-----------

Skaffold is available for some local workflows. This has generally been tested on local Minikube.

```
eval $(minikube docker-env)  # If trying to use a local image build instead of Docker Hub
skaffold dev --filename build/package/skaffold/skaffold.yaml
```

It's a good idea to make sure that the Helm Chart and Kustomize setups are producing roughly the
same things.

```
diff --side-by-side \
    <(helm template example ./build/package/helm | sed '/^#/d' | yq --prettyPrint 'sort_keys(..)') \
    <(kubectl kustomize ./build/package/kustomize/overlays/default/ | yq --prettyPrint 'sort_keys(..)')
```

The webserver itself has some tests.

```
cd src/example
go test
```
