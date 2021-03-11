#!/bin/sh

cd /projectdir

if [ $1 = "watch" ]
then
  reflex -s -r '(\.go$|go\.mod)' go run $PKG_CMD
elif [ $1 = "debug" ]
then
  reflex -s -r '(\.go$|go\.mod)' -- dlv debug --headless --api-version=2 --continue --accept-multiclient --listen=:2345 --log=false $PKG_CMD
else
  exit 1
fi
