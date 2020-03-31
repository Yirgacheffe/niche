#!/bin/sh

#
# Clean up niche services
#
readonly ns=niche-dev

# Cleanup greetings service
kubectl delete -f "./resources/niche-api-srv.yaml"     -n ${ns}
kubectl delete -f "./resources/niche-api-dest.yaml"    -n ${ns}
kubectl delete -f "./resources/niche-api-vs-5050.yaml" -n ${ns}

kubectl get pods -n ${ns} | grep niche-api
