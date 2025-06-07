package solvernet

import (
	"context"
	"crypto/ecdsa"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// Permit2 contract address.
	permit2Address = common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3")
)

// GeneratePermit2Signature generates the EIP-712 signature for Permit2 witness transfers.
// This replicates the exact logic from HashLib.gaslessOrderDigest.
func GeneratePermit2Signature(
	ctx context.Context,
	backend *ethbackend.Backend,
	userKey *ecdsa.PrivateKey,
	order bindings.IERC7683GaslessCrossChainOrder,
	orderData bindings.SolverNetOrderData,
	spender common.Address,
) ([]byte, error) {
	return generatePermit2Signature(ctx, backend, userKey, order, orderData, spender)
}

func generatePermit2Signature(
	ctx context.Context,
	backend *ethbackend.Backend,
	userKey *ecdsa.PrivateKey,
	order bindings.IERC7683GaslessCrossChainOrder,
	orderData bindings.SolverNetOrderData,
	spender common.Address,
) ([]byte, error) {
	// Get the Permit2 contract to access its domain separator
	permit2Contract, err := bindings.NewIPermit2(permit2Address, backend)
	if err != nil {
		return nil, errors.Wrap(err, "bind permit2 contract")
	}

	// Get the domain separator from Permit2
	domainSeparator, err := permit2Contract.DOMAINSEPARATOR(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "get permit2 domain separator")
	}

	// Build the digest using the exact same logic as HashLib.gaslessOrderDigest
	digest, err := buildGaslessOrderDigest(order, orderData, spender, domainSeparator)
	if err != nil {
		return nil, errors.Wrap(err, "build gasless order digest")
	}

	// Sign the digest with the user's private key
	signature, err := crypto.Sign(digest[:], userKey)
	if err != nil {
		return nil, errors.Wrap(err, "sign digest")
	}

	// Adjust the recovery ID for Ethereum compatibility (add 27)
	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}

// buildGaslessOrderDigest replicates HashLib.gaslessOrderDigest exactly.
func buildGaslessOrderDigest(
	order bindings.IERC7683GaslessCrossChainOrder,
	orderData bindings.SolverNetOrderData,
	inbox common.Address,
	domainSeparator [32]byte,
) ([32]byte, error) {
	// Type hashes from HashLib.sol - these must match exactly
	permit2TokenPermissionsTypeHash := crypto.Keccak256Hash([]byte("TokenPermissions(address token,uint256 amount)"))
	permit2WitnessTypeHash := crypto.Keccak256Hash([]byte("PermitWitnessTransferFrom(TokenPermissions permitted,address spender,uint256 nonce,uint256 deadline,GaslessCrossChainOrder witness)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,OmniOrderData orderData)OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)TokenExpense(address spender,address token,uint96 amount)TokenPermissions(address token,uint256 amount)"))

	// Hash TokenPermissions - matches HashLib.sol exactly
	tokenPermissionsArgs := abi.Arguments{
		{Type: mustType("bytes32")},
		{Type: mustType("address")},
		{Type: mustType("uint256")},
	}
	tokenPermissionsData, err := tokenPermissionsArgs.Pack(
		permit2TokenPermissionsTypeHash,
		orderData.Deposit.Token,
		orderData.Deposit.Amount,
	)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "pack token permissions")
	}
	tokenPermissionsHash := crypto.Keccak256Hash(tokenPermissionsData)

	// Generate witness hash - matches HashLib.witnessHash exactly
	witnessHash, err := buildWitnessHash(order, orderData)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "build witness hash")
	}

	// Create final struct hash - matches HashLib.sol exactly
	structHashArgs := abi.Arguments{
		{Type: mustType("bytes32")},
		{Type: mustType("bytes32")},
		{Type: mustType("address")},
		{Type: mustType("uint256")},
		{Type: mustType("uint32")},
		{Type: mustType("bytes32")},
	}
	structData, err := structHashArgs.Pack(
		permit2WitnessTypeHash,
		tokenPermissionsHash,
		inbox,
		order.Nonce,
		order.OpenDeadline,
		witnessHash,
	)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "pack struct hash")
	}
	structHash := crypto.Keccak256Hash(structData)

	// Create EIP-712 hash - matches HashLib.sol exactly
	digest := crypto.Keccak256Hash(
		append(
			append([]byte("\x19\x01"), domainSeparator[:]...),
			structHash[:]...,
		),
	)

	return digest, nil
}

// buildWitnessHash replicates HashLib.witnessHash exactly.
func buildWitnessHash(
	order bindings.IERC7683GaslessCrossChainOrder,
	orderData bindings.SolverNetOrderData,
) ([32]byte, error) {
	// Type hash from HashLib.sol - must match exactly
	gaslessOrderTypeHash := crypto.Keccak256Hash([]byte("GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,OmniOrderData orderData)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)TokenExpense(address spender,address token,uint96 amount)"))

	// Hash the OmniOrderData - matches HashLib.hashOmniOrderData exactly
	omniOrderDataHash, err := hashOmniOrderData(orderData)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hash omni order data")
	}

	// Create witness hash - matches HashLib.sol exactly
	witnessArgs := abi.Arguments{
		{Type: mustType("bytes32")},
		{Type: mustType("address")},
		{Type: mustType("address")},
		{Type: mustType("uint256")},
		{Type: mustType("uint256")},
		{Type: mustType("uint32")},
		{Type: mustType("uint32")},
		{Type: mustType("bytes32")},
		{Type: mustType("bytes32")},
	}
	witnessData, err := witnessArgs.Pack(
		gaslessOrderTypeHash,
		order.OriginSettler,
		order.User,
		order.Nonce,
		order.OriginChainId,
		order.OpenDeadline,
		order.FillDeadline,
		order.OrderDataType,
		omniOrderDataHash,
	)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "pack witness hash")
	}
	witnessHash := crypto.Keccak256Hash(witnessData)

	return witnessHash, nil
}

// hashOmniOrderData replicates HashLib.hashOmniOrderData exactly.
func hashOmniOrderData(orderData bindings.SolverNetOrderData) ([32]byte, error) {
	// Type hash from HashLib.sol - must match exactly
	omniOrderDataTypeHash := crypto.Keccak256Hash([]byte("OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)TokenExpense(address spender,address token,uint96 amount)"))

	// Hash the deposit
	depositHash, err := hashDeposit(orderData.Deposit)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hash deposit")
	}

	// Hash the calls
	callsHash, err := hashCalls(orderData.Calls)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hash calls")
	}

	// Hash the expenses
	expensesHash, err := hashTokenExpenses(orderData.Expenses)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "hash expenses")
	}

	// Create OmniOrderData hash - matches HashLib.sol exactly
	omniOrderDataArgs := abi.Arguments{
		{Type: mustType("bytes32")},
		{Type: mustType("address")},
		{Type: mustType("uint64")},
		{Type: mustType("bytes32")},
		{Type: mustType("bytes32")},
		{Type: mustType("bytes32")},
	}
	omniOrderData, err := omniOrderDataArgs.Pack(
		omniOrderDataTypeHash,
		orderData.Owner,
		orderData.DestChainId,
		depositHash,
		callsHash,
		expensesHash,
	)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "pack omni order data")
	}
	omniOrderDataHash := crypto.Keccak256Hash(omniOrderData)

	return omniOrderDataHash, nil
}

// hashDeposit replicates HashLib.hashDeposit exactly.
func hashDeposit(deposit bindings.SolverNetDeposit) ([32]byte, error) {
	depositTypeHash := crypto.Keccak256Hash([]byte("Deposit(address token,uint96 amount)"))

	depositArgs := abi.Arguments{
		{Type: mustType("bytes32")},
		{Type: mustType("address")},
		{Type: mustType("uint96")},
	}
	depositData, err := depositArgs.Pack(
		depositTypeHash,
		deposit.Token,
		deposit.Amount,
	)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "pack deposit")
	}
	depositHash := crypto.Keccak256Hash(depositData)

	return depositHash, nil
}

// hashCalls replicates HashLib.hashCalls exactly.
func hashCalls(calls []bindings.SolverNetCall) ([32]byte, error) {
	// Deterministic hash for empty array - matches HashLib.sol exactly
	if len(calls) == 0 {
		return crypto.Keccak256Hash([]byte("")), nil
	}

	callTypeHash := crypto.Keccak256Hash([]byte("Call(address target,bytes4 selector,uint256 value,bytes params)"))

	callHashes := make([][32]byte, len(calls))
	for i, call := range calls {
		// Extract selector from call params (first 4 bytes)
		var selector [4]byte
		if len(call.Params) >= 4 {
			copy(selector[:], call.Params[:4])
		}

		callArgs := abi.Arguments{
			{Type: mustType("bytes32")},
			{Type: mustType("address")},
			{Type: mustType("bytes4")},
			{Type: mustType("uint256")},
			{Type: mustType("bytes32")},
		}
		callData, err := callArgs.Pack(
			callTypeHash,
			call.Target,
			selector,
			call.Value,
			crypto.Keccak256Hash(call.Params),
		)
		if err != nil {
			return [32]byte{}, errors.Wrap(err, "pack call")
		}
		callHash := crypto.Keccak256Hash(callData)
		callHashes[i] = callHash
	}

	// Pack all call hashes together - matches HashLib.sol exactly
	var packed []byte
	for _, hash := range callHashes {
		packed = append(packed, hash[:]...)
	}

	return crypto.Keccak256Hash(packed), nil
}

// hashTokenExpenses replicates HashLib.hashTokenExpenses exactly.
func hashTokenExpenses(expenses []bindings.SolverNetTokenExpense) ([32]byte, error) {
	// Deterministic hash for empty array - matches HashLib.sol exactly
	if len(expenses) == 0 {
		return crypto.Keccak256Hash([]byte("")), nil
	}

	tokenExpenseTypeHash := crypto.Keccak256Hash([]byte("TokenExpense(address spender,address token,uint96 amount)"))

	expenseHashes := make([][32]byte, len(expenses))
	for i, expense := range expenses {
		expenseArgs := abi.Arguments{
			{Type: mustType("bytes32")},
			{Type: mustType("address")},
			{Type: mustType("address")},
			{Type: mustType("uint96")},
		}
		expenseData, err := expenseArgs.Pack(
			tokenExpenseTypeHash,
			expense.Spender,
			expense.Token,
			expense.Amount,
		)
		if err != nil {
			return [32]byte{}, errors.Wrap(err, "pack token expense")
		}
		expenseHash := crypto.Keccak256Hash(expenseData)
		expenseHashes[i] = expenseHash
	}

	// Pack all expense hashes together - matches HashLib.sol exactly
	var packed []byte
	for _, hash := range expenseHashes {
		packed = append(packed, hash[:]...)
	}

	return crypto.Keccak256Hash(packed), nil
}

// mustType returns an ABI type, panicking on error.
func mustType(typ string) abi.Type {
	t, err := abi.NewType(typ, "", nil)
	if err != nil {
		panic(err)
	}

	return t
}
