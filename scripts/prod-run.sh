#!/bin/bash

docker run --rm -d \
  --name tasx \
  --network host \
  --env-file=${1} \
  tasx_app
