#!/bin/sh

# start crispypod
crispypod &

# start front end
npm run preview &
# node build/index.js &

wait -n

exit $?

