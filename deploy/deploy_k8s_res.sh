#!/bin/bash

# Create Namespace
k apply -f ./k8s/niche-ns.yaml
k label namespace niche-dev istio-injection=enabled

# Create Secret
k apply -f ./k8s/niche-secret.yaml -n niche-dev
