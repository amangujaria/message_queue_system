#!/bin/sh

if [ "$1" == "server" ]
then
    go build server.go
    ./server
elif [ "$1" == "client" ]
then
    go build client.go
    ./client
#    for i in $(yes | sed $2q); do ./client & done
fi
