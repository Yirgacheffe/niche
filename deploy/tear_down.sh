# uninstall
$ helm template install/kubernetes/helm/istio --name istio --namespace istio-system | kubectl delete -f -
$ kubectl delete namespace istio-system
