package uniswap

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// routers maps chainID to Uniswap SwapRouter02 addreses
	// Reference: https://docs.uniswap.org/contracts/v3/reference/deployments/
	routers = map[uint64]common.Address{
		evmchain.IDEthereum:    addr("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"),
		evmchain.IDArbitrumOne: addr("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"),
		evmchain.IDBase:        addr("0x2626664c2603336E57B271c5C0b26F421741e481"),
		evmchain.IDOptimism:    addr("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"),
		evmchain.IDSepolia:     addr("0x3bFA4769FB09eefC5a80d6E87c3B9C650f7Ae48E"),
		evmchain.IDArbSepolia:  addr("0x101F443B4d1b059569D643917553c771E1b9663E"),
		evmchain.IDBaseSepolia: addr("0x94cC0AaC535CCDB3C01d6787D6413C739ae12bc4"),
		evmchain.IDOpSepolia:   addr("0x94cC0AaC535CCDB3C01d6787D6413C739ae12bc4"),
	}

	// quoter maps chainID to Uniswap QuoterV2 addreses
	// Reference: https://docs.uniswap.org/contracts/v3/reference/deployments/
	quoters = map[uint64]common.Address{
		evmchain.IDEthereum:    addr("0x61fFE014bA17989E743c5F6cB21bF9697530B21e"),
		evmchain.IDArbitrumOne: addr("0x61fFE014bA17989E743c5F6cB21bF9697530B21e"),
		evmchain.IDOptimism:    addr("0x61fFE014bA17989E743c5F6cB21bF9697530B21e"),
		evmchain.IDBase:        addr("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a"),
		evmchain.IDSepolia:     addr("0xEd1f6473345F45b75F8179591dd5bA1888cf2FB3"),
		evmchain.IDArbSepolia:  addr("0x2779a0CC1c3e0E44D2542EC3e79e3864Ae93Ef0B"),
		evmchain.IDBaseSepolia: addr("0xC5290058841028F1614F3A6F0F5816cAd0df5E27"),
		evmchain.IDOpSepolia:   addr("0x0FBEa6cf957d95ee9313490050F6A0DA68039404"),
	}
)

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
