#!/bin/sh
#
# Generate some traffic wifh fortio load testing
#

fortio load -t 10s http://localhost:80/search?city=TJ
