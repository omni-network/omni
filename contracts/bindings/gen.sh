#!/usr/bin/env bash
# Generate bindings for the given solidity contracts, to be

DIR=${DIR:-./bindings}
PKG=${PKG:-./bindings}

# generate bindings for the given contract
# params:
#  $1: contract name (ex. 'OmniPortal')
gen_binding() {
  name=$1

  # convert to lower case to respect golang package naming conventions
  name_lower=$(echo ${name} | tr '[:upper:]' '[:lower:]')

  temp=$(mktemp -d)
  forge inspect ${name} abi > ${temp}/${name}.abi
  forge inspect ${name} bytecode > ${temp}/${name}.bin

  abigen \
    --abi ${temp}/${name}.abi \
    --bin ${temp}/${name}.bin \
    --pkg ${PKG} \
    --type ${name} \
    --out ${DIR}/${name_lower}.go
}


for contract in $@; do
  gen_binding ${contract}
done
