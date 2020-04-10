#!/bin/sh

#
# Install greeting service into istio
#
readonly ns=niche-dev
kubectl apply -f ./niche-ingress-vs.yaml -n ${ns}
