kubectl apply -f niche-ns.yaml

kubectl label namespace niche-dev istio-injection=enabled

# kubectl api-resources

# kubectl get customresourcedefinition | grep istio.io


# download
http -d https://github.com/istio/istio/releases/download/1.3.5/istio-1.3.5-osx.tar.gz

# without tiller
$ kubectl create namespace istio-system
helm template . --name istio-init --namespace istio-system | kubectl apply -f -
kubectl get crds | grep 'istio-io' | wc -l

# uninstall
$ helm template install/kubernetes/helm/istio --name istio --namespace istio-system | kubectl delete -f -
$ kubectl delete namespace istio-system

# deploy resource
kubectl apply -n niche-dev -f niche-greetings-a-dep.yaml

# deply external mesh resource
k apply -f rabbitmq-se.yaml -n niche-dev




# misc command
kg ns -L istio-injection

# uninstall resource
kubectl delete -n niche-dev -f niche-greetings-a-dep.yaml
