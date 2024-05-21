#!/bin/bash

# turn on bash's job control
set -m

# Start backend and put it in background
/crispypod/crispypod serve --http=0.0.0.0:8080 &

# Start frontend
node start-server.js

# now we bring the front end back into the foreground
# and leave it there
fg %1