#!/bin/sh

#
# Clean up niche services
#
readonly ns=niche-dev

# Cleanup greetings service
kubectl delete -f "./services/niche-web-srv.yaml"  -n ${ns}
kubectl delete -f "./services/niche-web-dest.yaml" -n ${ns}

kubectl get pods -n ${ns} | grep niche-web