#!/bin/sh

#
# Install api service into istio
#
readonly ns=niche-dev

# Inject sidecar
istioctl kube-inject -f "./services/niche-api-srv.yaml" | kubectl apply -n ${ns} -f -
kubectl apply -f "./services/niche-api-dest.yaml" -n ${ns}

# List niche api pods
kubectl get pods -n ${ns} | grep niche-api
