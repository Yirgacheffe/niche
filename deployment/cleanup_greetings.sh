#!/bin/sh

#
# Clean up greeting services
#
readonly -a srvs=(a b c d e f g h)
readonly ns=niche-dev

for i in "${srvs[@]}"
do
    # Cleanup greetings service
    kubectl delete -f "./greeting/greetings-${i}-srv.yaml" -n ${ns}
    
done

kubectl get pods -n ${ns}
