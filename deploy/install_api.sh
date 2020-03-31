#!/bin/sh

#
# Install api service into istio
#
readonly ns=niche-dev

# Inject sidecar
istioctl kube-inject -f "./resources/niche-api-srv.yaml" | kubectl apply -n ${ns} -f -

kubectl apply -f "./resources/niche-api-dest.yaml"    -n ${ns}
kubectl apply -f "./resources/niche-api-vs-5050.yaml" -n ${ns}

# List niche api pods
kubectl get pods -n ${ns} | grep niche-api
