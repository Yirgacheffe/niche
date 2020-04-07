#!/bin/sh
#
# Build images for all greeting service
#

# --
readonly -a apis=(v1 v2 v3 v4)

for i in "${apis[@]}"
do
    cp -f build_api.sh "${i}"
    cp -f Dockerfile   "${i}"
    
    pushd "${i}"
    # Build go linuxc binary & docket it
    ./build_api.sh
    ./docker_it.sh
    
    # Clean up
    rm -rf Dockerfile
    rm -rf build_api.sh
    rm -rf niche-api
    popd
done

docker image ls | grep 'soline74/niche-api'
