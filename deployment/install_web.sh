#!/bin/sh

#
# Install web service into istio
#
readonly ns=niche-dev

# Inject sidecar
istioctl kube-inject -f "./web/niche-web-srv.yaml" | kubectl apply -n ${ns} -f -
kubectl  apply       -f "./web/niche-web-dest.yaml" -n ${ns}

# List niche web pods
kubectl get pods -n ${ns} | grep niche-web
