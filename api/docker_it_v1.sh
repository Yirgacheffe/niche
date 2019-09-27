#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='9/24/2019 13:00' --build-arg IMAGE_TAG_REF='dev-99d4d10' --build-arg VCS_REF='99d4d10' -t 'soline74/niche-api' .
docker tag soline74/niche-api soline74/niche-api:dev-99d4d10
