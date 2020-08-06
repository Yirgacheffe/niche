#!/bin/sh

#
# Install niche ingress virtual service
#
readonly ns=niche-dev
kubectl apply -f ./niche-ingress-gw.yaml -n ${ns}
