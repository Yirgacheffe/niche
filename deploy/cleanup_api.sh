#!/bin/sh

#
# Clean up niche services
#
readonly ns=niche-dev

# Cleanup greetings service
kubectl delete -f "./services/niche-api-srv.yaml"  -n ${ns}

kubectl delete -f "./services/niche-api-vs.yaml"   -n ${ns}
kubectl delete -f "./services/niche-api-dest.yaml" -n ${ns}

kubectl get pods -n ${ns} | grep niche-api