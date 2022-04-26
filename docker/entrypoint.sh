#!/bin/sh

set -e 

FILE=/ethermint/config/genesis.json

if [ ! -f "$FILE" ] && [ "$1" = 'ethermintd' ]; then
    ./init.sh

    exec "$@" "--"
elif [ "$1" = 'ethermintd' ]; then
    exec "$@" "--"
fi

exec "$@"
