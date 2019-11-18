# Istio
http -d https://github.com/istio/istio/releases/download/1.3.5/istio-1.3.5-osx.tar.gz

# REF: https://istio.io/docs/setup/install/helm/
# Option 1

# Install without tiller
kubectl create namespace istio-system

helm template install/kubernetes/helm/istio-init --name istio-init --namespace istio-system | kubectl apply -f -
kubectl get crds | grep 'istio-io' | wc -l

helm template install/kubernetes/helm/istio --name istio --namespace istio-system

# uninstall
helm template install/kubernetes/helm/istio --name istio --namespace istio-system | kubectl delete -f -
kubectl delete namespace istio-system
