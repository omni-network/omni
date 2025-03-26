package targets

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

const NameSymbiotic = "Symbiotic"

var (
	// Symbiotic testnet (mainnet vaults fetched from api).
	SymbioticSepoliaWSTETHVault1 = addr("0x77F170Dcd0439c0057055a6D7e5A1Eb9c48cCD2a")
	SymbioticSepoliaWSTETHVault2 = addr("0x1BAe55e4774372F6181DaAaB4Ca197A8D9CC06Dd")
	SymbioticSepoliaWSTETHVault3 = addr("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8")
	SymbioticHoleskyWSTETHVault1 = addr("0xd88dDf98fE4d161a66FB836bee4Ca469eb0E4a75")
	SymbioticHoleskyWSTETHVault2 = addr("0xa4c81649c79f8378a4409178E758B839F1d57a54")
)

// getSymbiotic returns the Symbiotic target.
func getSymbiotic(ctx context.Context) (Target, error) {
	mainnetVaults, err := getMainnetVaults(ctx)
	if err != nil {
		return Target{}, errors.Wrap(err, "get mainnet vaults")
	}

	return Target{
		Name: NameSymbiotic,
		Addresses: networkChainAddrs(map[uint64]map[common.Address]bool{
			evmchain.IDSepolia: set(
				SymbioticSepoliaWSTETHVault1,
				SymbioticSepoliaWSTETHVault2,
				SymbioticSepoliaWSTETHVault3,
			),
			evmchain.IDHolesky: set(
				SymbioticHoleskyWSTETHVault1,
				SymbioticHoleskyWSTETHVault2,
			),
			evmchain.IDEthereum: set(mainnetVaults...),
		}),
	}, nil
}

func getMainnetVaults(ctx context.Context) ([]common.Address, error) {
	const api = "https://app.symbiotic.fi/api/v2/vaults"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}

	defer resp.Body.Close()

	var vaults []struct {
		Address common.Address `json:"address"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&vaults); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	// Include all vaults, do not discriminate by token,
	// solver will reject unsupported tokens.
	var addresses []common.Address
	for _, v := range vaults {
		addresses = append(addresses, v.Address)
	}

	return addresses, nil
}
