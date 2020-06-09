#!/bin/sh

#
# Create niche app namespace & secrets
#
readonly ns=niche-dev

# Create namespace
kubectl apply -f ./niche-ns.yaml
kubectl get ns

# Install secrets
kubectl apply -f ./niche-secrets.yaml -n ${ns}
