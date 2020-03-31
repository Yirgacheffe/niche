#!/bin/sh

#
# Install greeting service into istio
#
readonly ns=niche-dev
kubectl apply -f ./niche-http-gw.yaml -n ${ns}
