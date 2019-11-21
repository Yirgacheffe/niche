# ------------------------------------------------------------------------------------------------

# Command for testing
kubectl port-forward -n niche-dev svc/niche-greetings-d 8090:8080
kubectl port-forward --namespace default svc/codefly-rabbitmq 15672:15672

# Command for restart a pods
kubectl replace --force -f xxxx.yaml

kubectl scale deployment esb-admin --replicas=0 -n {namespace}
kubectl scale deployment esb-admin --replicas=1 -n {namespace}

kubectl delete pod {podname} -n {namespace}
kubectl get pod {podname} -n {namespace} -o yaml | kubectl replace --force -f -

# kubectl api-resources

# kubectl get customresourcedefinition | grep istio.io


# misc command
kg ns -L istio-injection