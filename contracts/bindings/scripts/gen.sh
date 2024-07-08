#!/usr/bin/env bash
# Generate bindings for solidity contracts

# forge project root
ROOT=${ROOT}

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
  forge inspect --root ${ROOT} ${contract} abi > ${temp}/${name}.abi
  forge inspect --root ${ROOT} ${contract} bytecode > ${temp}/${name}.bin

  abigen \
    --abi ${temp}/${name}.abi \
    --bin ${temp}/${name}.bin \
    --type ${name} \
    --pkg bindings \
    --out ./bindings/${name_lower}.go
}


for contract in $@; do
  gen_binding ${contract}
done
