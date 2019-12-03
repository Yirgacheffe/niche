#!/bin/sh

# docker and tag it
docker build --build-arg BUILD_DATE='12/03/2019 11:19' --build-arg VERSION='dev-5c869b5' --build-arg VCS_REF='5c869b5' -t 'soline74/niche-web' .
docker tag soline74/niche-web soline74/niche-web:dev-5c869b5
