#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='12/05/2019 16:35' --build-arg IMAGE_TAG_REF='dev-b74ee79' --build-arg VCS_REF='b74ee79' -t 'soline74/niche-api' .
docker tag soline74/niche-api soline74/niche-api:dev-b74ee79
