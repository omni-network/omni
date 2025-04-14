package cctp

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

var (
	tokenMessengers = map[uint64]common.Address{
		evmchain.IDEthereum:    addr("0xbd3fa81b58ba92a82136038b25adec7066af3155"),
		evmchain.IDOptimism:    addr("0x2B4069517957735bE00ceE0fadAE88a26365528f"),
		evmchain.IDArbitrumOne: addr("0x19330d10D9Cc8751218eaf51E8885D058642E08A"),
		evmchain.IDBase:        addr("0x1682Ae6375C4E4A97e4B583BC394c861A46D8962"),
		evmchain.IDSepolia:     addr("0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5"),
		evmchain.IDArbSepolia:  addr("0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5"),
		evmchain.IDBaseSepolia: addr("0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5"),
		evmchain.IDOpSepolia:   addr("0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5"),
	}
	messageTransmitters = map[uint64]common.Address{
		evmchain.IDEthereum:    addr("0x0a992d191deec32afe36203ad87d7d289a738f81"),
		evmchain.IDOptimism:    addr("0x4d41f22c5a0e5c74090899e5a8fb597a8842b3e8"),
		evmchain.IDArbitrumOne: addr("0xC30362313FBBA5cf9163F0bb16a0e01f01A896ca"),
		evmchain.IDBase:        addr("0xAD09780d193884d503182aD4588450C416D6F9D4"),
		evmchain.IDSepolia:     addr("0x7865fAfC2db2093669d92c0F33AeEF291086BEFD"),
		evmchain.IDArbSepolia:  addr("0xaCF1ceeF35caAc005e15888dDb8A3515C41B4872"),
		evmchain.IDBaseSepolia: addr("0x7865fAfC2db2093669d92c0F33AeEF291086BEFD"),
		evmchain.IDOpSepolia:   addr("0x7865fAfC2db2093669d92c0F33AeEF291086BEFD"),
	}
)

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
