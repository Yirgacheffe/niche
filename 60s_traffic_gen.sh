#!/bin/sh
#
# Generate some traffic wifh fortio load testing
#

hey -z 1m -c 5 http://localhost:80/search?city=TJ