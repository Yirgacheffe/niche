#!/bin/sh

#
# Clean up niche api services
#
readonly ns=niche-dev

# Cleanup greetings service
kubectl delete -f "./configs/configs-srv.yaml"     -n ${ns}
kubectl delete -f "./configs/configs-dest.yaml"    -n ${ns}
kubectl delete -f "./configs/configs-vs-5050.yaml" -n ${ns}

kubectl get pods -n ${ns} | grep niche-configs
