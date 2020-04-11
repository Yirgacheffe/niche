#!/bin/sh

#
# Install sleep tools for testing
#

# Inject sidecar
# Located to 'default' namespace
istioctl kube-inject -f "./sleep.yaml" | kubectl apply -f -

# List pods
kubectl get pods
