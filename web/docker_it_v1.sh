#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='9/24/2019 17:00' --build-arg VERSION='dev-81fbdd3' --build-arg VCS_REF='81fbdd3' -t 'soline74/niche-web' .
docker tag soline74/niche-web soline74/niche-web:dev-81fbdd3
