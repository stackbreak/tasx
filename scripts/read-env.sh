#!/bin/bash

function read_env() {
  local envFile=$1
  local isComment='^[[:space:]]*#'
  local isBlank='^[[:space:]]*$'
  while IFS= read -r line; do
    [[ $line =~ $isComment ]] && continue
    [[ $line =~ $isBlank ]] && continue
    key=$(echo "$line" | cut -d '=' -f 1)
    value=$(echo "$line" | cut -d '=' -f 2-)
    eval "export ${key}=\"$(echo \${value})\""
  done < <( cat "$envFile" )
}
