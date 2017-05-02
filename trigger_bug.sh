#!/bin/bash

rm -rf vendor

if $( which govendor > /dev/null 2>&1 ); then
    govendor init
    govendor fetch +missing
else
    echo "govendor not found in PATH"
fi
