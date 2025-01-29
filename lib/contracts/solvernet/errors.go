package solvernet

import (
	"strings"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func DetectCustomError(custom error) string {
	contracts := map[string]*bind.MetaData{
		"inbox":      bindings.SolverNetInboxMetaData,
		"outbox":     bindings.SolverNetOutboxMetaData,
		"mock_erc20": bindings.MockERC20MetaData,
	}

	errMsg := custom.Error()

	// errors from SafeTransferLib, not present in abis
	safeTranserLibErrs := map[string]string{
		"0x90b8ec18": "TransferFailed()",
		"0x3e3f8f73": "ApproveFailed()",
		"0x7939f424": "TransferFromFailed()",
		"0xb12d13eb": "ETHTransferFailed()",
		"0x54cd9435": "TotalSupplyQueryFailed()",
		"0x6b836e6b": "Permit2Failed()",
		"0x8757f0fd": "Permit2AmountOverflow()",
	}

	for id, msg := range safeTranserLibErrs {
		if strings.Contains(errMsg, id) {
			return "SafeTransferLib::" + msg
		}
	}

	for name, contract := range contracts {
		abi, err := contract.GetAbi()
		if err != nil {
			return "BUG"
		}
		for n, e := range abi.Errors {
			if strings.Contains(errMsg, e.ID.Hex()[:10]) {
				return name + "::" + n
			}
		}
	}

	return "unknown"
}
