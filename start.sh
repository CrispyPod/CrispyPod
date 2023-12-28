#!/bin/sh
export GIN_MODE=release

# start crispypod
crispypod &

# start front end
node build-node/index.js &

wait -n

exit $?