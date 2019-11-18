#!/bin/bash

# deploy resource
k apply -n niche-dev -f niche-greetings-a-dep.yaml

# deploy greeting service in istio
k apply -f niche-greetings-d-dep.yaml -n niche-dev
k apply -f niche-greetings-d-svc.yaml -n niche-dev
