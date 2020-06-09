#!/bin/sh
#
# Build images for all greeting service
#

# --
readonly -a vers=(v1 v2 v3 v4)

for i in "${vers[@]}"
do
    cp -f build.sh   "${i}"
    cp -f Dockerfile "${i}"
    
    pushd "${i}"
    # Build go linuxc binary & docket it
    ./build.sh
    ./docker_it.sh
    
    # Clean up
    rm -rf Dockerfile
    rm -rf build.sh
    rm -rf niche-configs
    popd
done

docker image ls | grep 'soline74/niche-configs'
