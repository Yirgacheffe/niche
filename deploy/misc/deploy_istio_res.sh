#!/bin/bash

# deploy external mesh resource ( rabbitmq & mongodb)
k apply -f rabbitmq-se.yaml -n niche-dev
