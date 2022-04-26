#!/bin/sh

set -e

FILE=/ethermint/config/genesis.json

if [ ! -f "$FILE" ] && [ "$1" = 'dlv' ]; then
    ./init.sh

    exec "$@" "--"
elif [ "$1" = 'dlv' ]; then
    exec "$@" "--"
fi

exec "$@"
