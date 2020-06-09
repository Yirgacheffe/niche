#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='03/31/2020 11:05' --build-arg IMAGE_TAG_REF='dev-e1be533' --build-arg VCS_REF='e1be533' -t 'soline74/niche-configs' .
docker tag soline74/niche-configs soline74/niche-configs:dev-e1be533
