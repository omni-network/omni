package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type APIKeys struct {
	EtherscanAPIKey string
	ArbscanAPIKey   string
}

func bindFlags(flags *pflag.FlagSet, apikeys *APIKeys) {
	flags.StringVar(&apikeys.EtherscanAPIKey, "etherscan-api-key", apikeys.EtherscanAPIKey, "Etherscan API key")
	flags.StringVar(&apikeys.ArbscanAPIKey, "arbscan-api-key", apikeys.ArbscanAPIKey, "Arbscan API key")
}

func newGetCreationTxHashCmd() *cobra.Command {
	apikeys := &APIKeys{
		EtherscanAPIKey: "",
		ArbscanAPIKey:   "",
	}

	cmd := &cobra.Command{
		Use:   "get-creation-tx-hash",
		Short: "Get contract creation tx hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			chain := ChainName(args[0])
			if err := chain.Validate(); err != nil {
				return err
			}

			addr := common.HexToAddress(args[1])

			hash, err := creationTxHash(ctx, chain, addr, apikeys)
			if err != nil {
				return err
			}

			log.Info(ctx, "✅", "tx_hash", hash.Hex())

			return nil
		},
	}

	bindFlags(cmd.Flags(), apikeys)

	return cmd
}

type GetContractCreationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		ContractAddress string `json:"contractAddress"`
		ContractCreator string `json:"contractCreator"`
		TxHash          string `json:"txHash"`
	} `json:"result"`
}

func creationTxHash(ctx context.Context, chain ChainName, addr common.Address, apikeys *APIKeys) (common.Hash, error) {
	apikey := apikeys.EtherscanAPIKey
	if chain == ArbSepolia {
		apikey = apikeys.ArbscanAPIKey
	}

	url := fmt.Sprintf(
		chain.VerifierURL()+"?module=contract&action=getcontractcreation&contractaddresses=%s&apikey=%s",
		addr.Hex(),
		apikey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "create request", "url", url)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "get contract creation", "url", url)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "read response", "url", url)
	}

	var data GetContractCreationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "unmarshal response", "url", url)
	}

	if len(data.Result) != 1 {
		return common.Hash{}, errors.New("unexpected response", "url", url, "data", data)
	}

	hash := common.HexToHash(data.Result[0].TxHash)

	return hash, nil
}
