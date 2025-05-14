#!/usr/bin/env bash

# Runs scripts/golint for all touched packages
MOD=$(go list -m)
PKGS=$(echo "$@"| xargs -n1 dirname | sort -u | sed -e "s#^#${MOD}/#")

go run github.com/omni-network/omni/scripts/golint $PKGS
