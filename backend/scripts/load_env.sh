#!/bin/bash

if [ ! -f '.env' ]; then
    echo ".env file does not exist at the root directory" >&2
    exit 1
fi

while IFS= read -r line; do
    if [[ "$line" =~ ^# || -z "$line" ]]; then
        continue
    fi
    export "$line"
done < ".env"

echo "Environment variables have been loaded"
