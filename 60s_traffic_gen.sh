#!/bin/sh
#
# Generate some traffic with hey
#

hey -z 1m -c 5 http://localhost:80/search?city=TJ
