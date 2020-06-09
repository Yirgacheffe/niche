#!/bin/sh
#
# Check Istio outbound traffic policy
#

kubectl get configmap istio -n istio-system -o yaml | grep -o "mode: REGISTRY_ONLY"


# 
# Using following command to restore the settings
#
# kgcm istio -n istio-system -o yaml | sed 's/mode: REGISTRY_ONLY/mode: ALLOW_ANY/g' | k replace -n istio-system -f -
#
