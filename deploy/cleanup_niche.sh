#!/bin/sh

#
# Clean up niche services
#
readonly -a srvs=(a b c d e f g h)
readonly ns=niche-dev

for i in "${srvs[@]}"
do
    # Cleanup greetings service
    kubectl delete -f "./services/greetings-${i}-srv.yaml" -n ${ns}
    
done

kubectl get pods -n ${ns}

kubectl delete -f ./niche-http-gw.yaml -n ${ns}
