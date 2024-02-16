#!/usr/bin/env bash

# This scripts generates all the protobufs using 'buf generate'.
# Cosmos however uses a mix of gogo, pulsar and orm plugins for code generation.
# So we manually call buf generate for each type of plugin.

function bufgen() {
    TYPE=$1 # Either orm,pulsar,proto
    FILE=$2 # Path to proto file to generate
    OUTDIR=$3 # Output directory

    echo "  ${TYPE}: ${FILE}"

    buf generate \
      --template="scripts/buf.gen.${TYPE}.yaml" \
      --output="${OUTDIR}" \
      --path="${FILE}"
}

# Ensure we are in the root of the repo, so  ${pwd}/go.mod must exit
if [ ! -f go.mod ]; then
  echo "Please run this script from the root of the repository"
  exit 1
fi

echo "Generating pulsar protos for cosmos module config"
for FILE in halo2/*/module/*.proto halo/halopb/*/*.proto
do
  bufgen pulsar "${FILE}" "."
done

echo "Generating gogo protos for cosmos module types"
for FILE in halo2/*/types/*.proto
do
  bufgen gogo "${FILE}" "."
done

echo "Generating orm protos for cosmos keeper orm"
for FILE in halo/aggregate/*/*.proto
do
  bufgen orm "${FILE}" "."
done
