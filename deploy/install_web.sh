#!/bin/sh

#
# Install api service into istio
#
readonly ns=niche-dev

# Inject sidecar
istioctl kube-inject -f "./resources/niche-web-srv.yaml" | kubectl apply -n ${ns} -f -
kubectl  apply       -f "./resources/niche-web-dest.yaml" -n ${ns}

# List niche api pods
kubectl get pods -n ${ns} | grep niche-web
