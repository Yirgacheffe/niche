# Deployment shells
k apply -f ./k8s/niche-ns.yaml
k label namespace niche-dev istio-injection=enabled


# deploy secret
k apply -f niche-secret.yaml -n niche-dev

# deploy resource
k apply -n niche-dev -f niche-greetings-a-dep.yaml

# deploy greeting service in istio
k apply -f niche-greetings-d-dep.yaml -n niche-dev
k apply -f niche-greetings-d-svc.yaml -n niche-dev


# deploy external mesh resource ( rabbitmq & mongodb)
k apply -f rabbitmq-se.yaml -n niche-dev
