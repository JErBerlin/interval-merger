#!/usr/bin/env bash
# File: remove_nl.sh 
#!/bin/bash

# Check if input and output file are the same
if [ "$1" == "$2" ]; then
    # Remove newline characters and insert a space between lines
    sed -i ':a;N;$!ba;s/\n/ /g' "$1"
else
    # Remove newline characters and insert a space between lines
    tr -d '\n' < "$1" | sed 's/\n/ /g' > "$2"
fi
