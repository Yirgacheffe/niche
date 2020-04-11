#!/bin/sh
#
# Curl wttr.in testing for outbound traffic
#

SLEEP_SOURCE_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})

echo "Start access wttr.in from POD: ${SLEEP_SOURCE_POD}"
kubectl exec ${SLEEP_SOURCE_POD} -c sleep -- curl -sL -D - http://wttr.in/tianjin
