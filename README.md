example-k8s
===========

This produces a very simple Go webserver which is meant to be useful as an
example application for testing a new Kubernetes server.

Installation
------------

### Kubectl

      kubectl create namespace example
      kubectl run --image spacez320/example-k8s:latest example --namespace example

### Helm

      helm upgrade --install --namespace example example ./helm
