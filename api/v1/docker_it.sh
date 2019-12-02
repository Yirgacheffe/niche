#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='12/01/2019 10:23' --build-arg IMAGE_TAG_REF='dev-9248113' --build-arg VCS_REF='9248113' -t 'soline74/niche-api' .
docker tag soline74/niche-api soline74/niche-api:dev-9248113
