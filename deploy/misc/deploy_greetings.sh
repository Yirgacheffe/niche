#!/bin/bash

# deploy resource
kubectl apply -f greetings-a-srv.yaml -n niche-dev
kubectl apply -f greetings-d-srv.yaml -n niche-dev

