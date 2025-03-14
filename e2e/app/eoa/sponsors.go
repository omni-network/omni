package eoa

import (
	"fmt"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type Sponsor struct {
	Address        common.Address
	ChainID        uint64
	Name           string
	FundThresholds FundThresholds
}

// genstakeSponsors returns a list of GenesisStake migration sponsors.
// Each sponsor has target 2 ETH, min 1 ETH fund threshold. Total target 20 ETH, min 10 ETH.
func genstakeSponsors(
	chainID uint64,
	addresses ...string,
) []Sponsor {
	if len(addresses) != 10 {
		panic("must have 10 genesis stake sponsors")
	}

	var sponsors []Sponsor
	for i, address := range addresses {
		sponsors = append(sponsors, Sponsor{
			Address: common.HexToAddress(address),
			ChainID: chainID,
			Name:    fmt.Sprintf("genstake-sponsor-%d", i),
			FundThresholds: FundThresholds{
				minEther:    1,
				targetEther: 2,
			},
		})
	}

	return sponsors
}

var (
	sponsors = map[netconf.ID][]Sponsor{
		netconf.Mainnet: genstakeSponsors(evmchain.IDEthereum,
			"0x20350BD9216895f5B2D44655664a446fc5cD6821",
			"0x2908BBC5C02af08DEFA5608B33640353841f094B",
			"0xd03085EbAee6Edc58B5390162FC01d015371e684",
			"0x06582FF52b9F4D8c659aFab3731a6317eC8aa604",
			"0xD5532Ed5E581e05E736086e0208d9484b80b6ceC",
			"0x6d6B535F30Baa50769Cbe48bEe2D9C1699C9c2F9",
			"0xa74C6747F5B4E332B8f2dD865AA15D0429fdbE8F",
			"0x8b01F99569630F32db76ECabcb8472d6cd4051DD",
			"0xac722049F6D87cd9DFd9be35f6D3cF4E0C5c5478",
			"0x63D03B7d9D9D42C3E133D5ed3434A954105e4335",
		),
		netconf.Omega: genstakeSponsors(evmchain.IDHolesky,
			"0xf01E77c3D61D66E2bA2DEe0bD6e6cbA4f07c12DE",
			"0x8c10eE9caab1a2a32581C2f164f1C549c12E7EBC",
			"0x86c20f92E74B0f67d70aE35b3203060b15c024cc",
			"0xEEe09549E474fbb0d06DB98243e11b5633BE6B1a",
			"0x02BD9FD3A8C8cB6BDc7d37Fdd000e0113D5BA69B",
			"0xC1e6ec370fCEA7B04d5730690BB7c95Ca67A5492",
			"0xA43D62250219106B5FA16E02Ae9D82F3f142DbB2",
			"0xB78c1CAF2768580d8C3AF53c9F0c23bD9Fa88096",
			"0x77622Ad2a608fB4b1205aA809d1bE3A13Ea3F692",
			"0xc881A382A197940BEB4c843FD0830B5B2f35390D",
		),
		// same as omega
		netconf.Staging: genstakeSponsors(evmchain.IDHolesky,
			"0xf01E77c3D61D66E2bA2DEe0bD6e6cbA4f07c12DE",
			"0x8c10eE9caab1a2a32581C2f164f1C549c12E7EBC",
			"0x86c20f92E74B0f67d70aE35b3203060b15c024cc",
			"0xEEe09549E474fbb0d06DB98243e11b5633BE6B1a",
			"0x02BD9FD3A8C8cB6BDc7d37Fdd000e0113D5BA69B",
			"0xC1e6ec370fCEA7B04d5730690BB7c95Ca67A5492",
			"0xA43D62250219106B5FA16E02Ae9D82F3f142DbB2",
			"0xB78c1CAF2768580d8C3AF53c9F0c23bD9Fa88096",
			"0x77622Ad2a608fB4b1205aA809d1bE3A13Ea3F692",
			"0xc881A382A197940BEB4c843FD0830B5B2f35390D",
		),
	}
)

// AllSponsors returns all sponsors for the network.
func AllSponsors(network netconf.ID) []Sponsor {
	return sponsors[network]
}
