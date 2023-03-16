#!/bin/bash

if [ "$1" == "test" ]; then
    go test -v
elif [ "$1" == "test-report" ]; then
    go test -v
    node index.js
else
    echo "Usage: run [test|test-report]"
    exit 1
fi

wait
echo "All processes Testing is Completed"
