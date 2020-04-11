#!/bin/sh

#
# Clean up sleep service
#

# Cleanup sleep service
kubectl delete -f "./sleep.yaml" 

# List pods
kubectl get pods
