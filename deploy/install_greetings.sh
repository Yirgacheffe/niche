#!/bin/sh

#
# Install greeting service into istio
#
readonly -a srvs=(a b c d e f g h)
readonly ns=niche-dev

for i in "${srvs[@]}"
do
    # inject istio proxy
    istioctl kube-inject -f "./resources/greetings-${i}-srv.yaml" | kubectl apply -n ${ns} -f -    
done

kubectl get pods -n ${ns}
