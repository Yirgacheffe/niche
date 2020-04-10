#!/bin/sh

#
# Clean up niche services
#
readonly ns=niche-dev
kubectl delete -f ./niche-ingress-vs.yaml -n ${ns}
