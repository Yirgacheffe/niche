#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='04/07/2020 13:56' --build-arg IMAGE_TAG_REF='dev-519a1a0' --build-arg VCS_REF='519a1a0' -t 'soline74/niche-api' .
docker tag soline74/niche-api soline74/niche-api:dev-519a1a0
