#!/bin/bash

# turn on bash's job control
set -m

# Start backend and put it in background
crispypod &

# Start frontend
node build-node/index.js

# now we bring the front end back into the foreground
# and leave it there
fg %1