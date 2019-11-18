# uninstall
$ helm template install/kubernetes/helm/istio --name istio --namespace istio-system | kubectl delete -f -
$ kubectl delete namespace istio-system


# uninstall resource
kubectl delete -n niche-dev -f niche-greetings-a-dep.yaml