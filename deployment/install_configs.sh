#!/bin/sh

#
# Install api service into istio
#
readonly ns=niche-dev

# Inject sidecar
istioctl kube-inject -f "./configs/configs-srv.yaml" | kubectl apply -n ${ns} -f -

kubectl apply -f "./configs/configs-dest.yaml"    -n ${ns}
kubectl apply -f "./configs/configs-vs-5050.yaml" -n ${ns}

# List niche configs pod
kubectl get pods -n ${ns} | grep niche-configs
