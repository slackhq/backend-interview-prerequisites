#!/bin/bash

set -e
set -x

project_root="$( cd "$( dirname "${BASH_SOURCE[0]}" )/" && pwd )"
cd "$project_root"

# delete current db and logs
rm -f datastore.db 2>/dev/null
rm -f ./server.log

# create new logs
touch ./server.log

# start the server in the background and give it a second to attach to the network interface
./setup -db datastore.db -schema schema.sql > ./server.log 2>&1 &
sleep 1

tail -F ./server.log &

# kill the server when this script completes
trap 'kill $(jobs -p)' EXIT

# run the client
cd ./client

# clean the test cache so tests are ran each time
go clean -testcache
go test .
