#!/bin/sh

#
# Clean up niche services
#
readonly ns=niche-dev
kubectl delete -f ./niche-http-gw.yaml -n ${ns}
