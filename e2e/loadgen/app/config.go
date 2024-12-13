package app

import (
	"fmt"
	"strings"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

type Config struct {
	Count        int
	Network      string
	Role         string
	ChainIDPairs string
}

type ChainIDPair struct {
	From string
	To   string
}

func (c Config) String() string {
	var s strings.Builder
	_, _ = s.WriteString(fmt.Sprintf("count:%d\n", c.Count))
	_, _ = s.WriteString(fmt.Sprintf("metwork:%s\n", c.Network))
	_, _ = s.WriteString(fmt.Sprintf("role:%s\n", c.Role))
	_, _ = s.WriteString("Chain ID XCall Pairs from->to:\n")
	for _, pair := range c.GetChainIDPairs() {
		_, _ = s.WriteString(fmt.Sprintf(" %s -> %s\n", pair.From, pair.To))
	}

	return s.String()
}

func (c Config) GetChainIDPairs() []ChainIDPair {
	pairs := make([]ChainIDPair, 0)
	for _, pair := range strings.Split(c.ChainIDPairs, ",") {
		chainIDs := strings.Split(pair, ":")
		pairs = append(pairs, ChainIDPair{
			From: chainIDs[0],
			To:   chainIDs[1],
		})
	}

	return pairs
}

func DefaultConfig() Config {
	var chainIDPairs strings.Builder
	// all possible combinations of chainids for the default network staging
	chainIDMetadata := evmchain.Testnets()
	for _, xCallFromChain := range chainIDMetadata {
		for _, xCallToChain := range chainIDMetadata {
			if xCallFromChain == xCallToChain {
				continue
			}
			_, _ = chainIDPairs.WriteString(fmt.Sprintf("%s:%s,", xCallFromChain.Name, xCallToChain.Name))
		}
	}

	chainIDPairsStr := chainIDPairs.String()[:chainIDPairs.Len()-1]

	return Config{
		Count:        100,
		Network:      netconf.Staging.String(),
		Role:         string(eoa.RoleXCaller),
		ChainIDPairs: chainIDPairsStr,
	}
}
