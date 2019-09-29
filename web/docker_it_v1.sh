#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='9/29/2019 10:14' --build-arg VERSION='dev-ac756d1' --build-arg VCS_REF='ac756d1' -t 'soline74/niche-web' .
docker tag soline74/niche-web soline74/niche-web:dev-ac756d1
