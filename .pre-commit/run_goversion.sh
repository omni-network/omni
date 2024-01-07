#!/usr/bin/env bash

# run_goversion.sh ensures that the local go version matches the go.mod minor version.
# Minor version matching is fine since local go isn't used to build production binaries, that is done by CI.

# minor returns the minor version from the provided string.
# e.g. go1.14.2 -> 1.14
function minor() (
  REGEX="([0-9]+\.[0-9]+)\.[0-9]+"
  if [[ ! $1 =~ $REGEX ]]; then
     echo "Failed parsing minor version from $1"
     exit 1
  fi

  echo "${BASH_REMATCH[1]}"
)

ACTUAL=$(minor "$(go version)")
EXPECTED=$(minor "$(go list -m -f '{{.GoVersion}}')")

if [[ "$ACTUAL" != "$EXPECTED" ]]; then
   echo "Go version mismatch; $EXPECTED vs $ACTUAL"
   echo "Please update local go installation to match go.mod minor version"
   exit 1
fi
