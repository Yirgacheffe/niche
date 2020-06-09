#!/bin/sh

#
# Clean up niche ingress
#
readonly ns=niche-dev
kubectl delete -f ./niche-ingress-vs.yaml -n ${ns}
