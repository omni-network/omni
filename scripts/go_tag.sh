#!/bin/bash

# go_tag.sh return the Nth latest go git tag; e.g. v0.15.0.
# It filters out any JS SDK tags like `@omni-network/react@0.2.1`
# For the current tag, use `go_tag.sh 0`.
# For the previous tag, use `go_tag.sh 1`.

N=$1
count=0
for tag in $(git tag --sort=-version:refname); do # Sort by *semver*
    # Return Nth match
    if [[ $count -eq $N ]]; then
        echo "$tag"
        exit 0
    fi

    ((count++))
done

# No match found
exit 1
