#!/usr/bin/env bash

# Generate custom enum bindings for some specific contracts only.
# this script should be run from /contracts directory.

gen() {
  input=$1
  output=$2
  package=$3
  prefix=$4

  go run github.com/omni-network/omni/scripts/solenumgen \
    --input ${input} \
    --output ${output} \
    --package ${package} \
    --prefix ${prefix}
}

gen solve/src/interfaces/ISolverNetInbox.sol ../lib/contracts/solvernet/enum_gen.go solvernet Order
