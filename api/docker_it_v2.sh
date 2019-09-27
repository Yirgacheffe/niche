#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='9/25/2019 13:00' --build-arg IMAGE_TAG_REF='dev-f0c7b9e' --build-arg VCS_REF='f0c7b9e' -t 'soline74/niche-api' .
docker tag soline74/niche-api soline74/niche-api:dev-f0c7b9e
