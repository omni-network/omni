#!/usr/bin/env bash
# Generate bindings for solidity contracts

DIR=${DIR:-./bindings}
PKG=${PKG:-bindings}

# generate bindings for the given contract
# works on contract name of fully qualified path to the contract
# params:
#  $1: contract name (ex. 'OmniPortal', 'src/protocol/OmniPortal.sol:OmniPortal')
gen_binding() {
  contract=$1

  # strip path prefix, if used
  # ex src/protocol/OmniPortal.sol:OmniPortal => OmniPortal
  name=$(echo ${contract} | cut -d ":" -f 2)

  # convert to lower case to respect golang package naming conventions
  name_lower=$(echo ${name} | tr '[:upper:]' '[:lower:]')

  temp=$(mktemp -d)
  forge inspect ${contract} abi > ${temp}/${name}.abi
  forge inspect ${contract} bytecode > ${temp}/${name}.bin

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
