package targets

import (
	"context"
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

const namePrefix = "Symbiotic:"

var (
	// Symbiotic testnet (mainnet vaults fetched from api).
	SymbioticSepoliaWSTETHVault1 = addr("0x77F170Dcd0439c0057055a6D7e5A1Eb9c48cCD2a")
	SymbioticSepoliaWSTETHVault2 = addr("0x1BAe55e4774372F6181DaAaB4Ca197A8D9CC06Dd")
	SymbioticSepoliaWSTETHVault3 = addr("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8")
	SymbioticHoleskyWSTETHVault1 = addr("0xd88dDf98fE4d161a66FB836bee4Ca469eb0E4a75")
	SymbioticHoleskyWSTETHVault2 = addr("0xa4c81649c79f8378a4409178E758B839F1d57a54")
)

// getSymbiotic returns the Symbiotic targets.
func getSymbiotic(ctx context.Context, network netconf.ID) ([]Target, error) {
	if network == netconf.Devnet {
		return nil, nil // Symbiotic not on devnet
	} else if network == netconf.Omega || network == netconf.Staging {
		// Staging and omega on has few hardcoded wstETH vaults
		return []Target{{
			Name: namePrefix + "wstETH",
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
			}),
		}}, nil
	}

	vaults, err := getMainnetVaults(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get mainnet vaults")
	}

	// Convert map[vault address]symbol to map[symbol]map[vault address]bool
	var vaultsBySymbol = make(map[string]map[common.Address]bool)
	for vault, symbol := range vaults {
		if _, ok := vaultsBySymbol[symbol]; !ok {
			vaultsBySymbol[symbol] = make(map[common.Address]bool)
		}
		vaultsBySymbol[symbol][vault] = true
	}

	// Convert into list of targets
	var targets []Target
	for symbol, vaults := range vaultsBySymbol {
		targets = append(targets, Target{
			Name: namePrefix + symbol,
			Addresses: networkChainAddrs(map[uint64]map[common.Address]bool{
				evmchain.IDEthereum: vaults,
			}),
		})
	}

	return targets, nil
}

func getMainnetVaults(ctx context.Context) (map[common.Address]string, error) {
	all := make(map[common.Address]string)
	for i := 1; ; i++ {
		next, err := getMainnetVaultsPage(ctx, i)
		if err != nil {
			return nil, errors.Wrap(err, "get mainnet vaults page")
		} else if len(next) == 0 {
			break
		}

		maps.Copy(all, next)

		time.Sleep(time.Millisecond * 500) // Rate limit
	}

	return all, nil
}

func getMainnetVaultsPage(ctx context.Context, page int) (map[common.Address]string, error) {
	const api = "https://app.symbiotic.fi/api/v2/vaults"
	uri, err := url.Parse(api)
	if err != nil {
		return nil, errors.Wrap(err, "parse api")
	}

	q := uri.Query()
	q.Set("p", strconv.Itoa(page))
	uri.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("nok response", "status", resp.Status, "body", string(body))
	}

	var vaults []struct {
		Address common.Address `json:"address"`
		Token   struct {
			Symbol string `json:"symbol"`
		} `json:"token"`
	}
	if err := json.Unmarshal(body, &vaults); err != nil {
		return nil, errors.Wrap(err, "decode response", "body", string(body))
	}

	// Include all vaults, do not discriminate by token,
	// solver will reject unsupported tokens.
	vaultsBySymbol := make(map[common.Address]string)
	for _, v := range vaults {
		if v.Token.Symbol == "" {
			return nil, errors.New("vault missing token")
		}

		vaultsBySymbol[v.Address] = v.Token.Symbol
	}

	return vaultsBySymbol, nil
}
