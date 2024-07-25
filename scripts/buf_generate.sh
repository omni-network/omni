#!/usr/bin/env bash

# This scripts generates all the protobufs using 'buf generate'.
# Cosmos however uses a mix of gogo, pulsar and orm plugins for code generation.
# So we manually call buf generate for each type of plugin.

function bufgen() {
    TYPE=$1 # Either orm,pulsar,proto
    DIR=$2 # Path to dir containing protos to generate

    # Skip if ${DIR}/*.proto does not exist
    if ! test -n "$(find "${DIR}" -maxdepth 1 -name '*.proto')"; then
      return
    fi


    echo "  ${TYPE}: ${DIR}"

    buf generate \
      --template="scripts/buf.gen.${TYPE}.yaml" \
      --path="${DIR}"
}

# Ensure we are in the root of the repo, so  ${pwd}/go.mod must exit
if [ ! -f go.mod ]; then
  echo "Please run this script from the root of the repository"
  exit 1
fi

echo "Generating pulsar protos for cosmos module config"
for DIR in halo/*/module/ octane/*/module/
do
  bufgen pulsar "${DIR}"
done

echo "Generating gogo protos for cosmos module types"
for DIR in halo/*/types/ octane/*/types/ halo/genutil/genserve
do
  bufgen gogo "${DIR}"
done

echo "Generating orm protos for cosmos keeper orm"
for DIR in halo/*/keeper/ octane/*/keeper/ monitor/xmonitor/emitcache
do
  bufgen orm "${DIR}"
done
