#!/usr/bin/env bash
# Generate bindings for the given solidity contracts, to be

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
    --pkg bindings \
    --type ${name} \
    --out ./bindings/${name_lower}.go
}


for contract in $@; do
  gen_binding ${contract} $(mktemp -d)
done
