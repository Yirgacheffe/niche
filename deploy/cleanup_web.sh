#!/bin/sh

#
# Clean up niche web services
#
readonly ns=niche-dev

# Cleanup greetings service
kubectl delete -f "./resources/niche-web-srv.yaml"  -n ${ns}
kubectl delete -f "./resources/niche-web-dest.yaml" -n ${ns}

kubectl get pods -n ${ns} | grep niche-web