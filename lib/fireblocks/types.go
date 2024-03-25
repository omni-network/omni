package fireblocks

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/crypto"
)

// Status of a transaction. See https://developers.fireblocks.com/reference/primary-transaction-statuses.
type Status string

const (
	StatusCompleted            Status = "COMPLETED"
	StatusFailed               Status = "FAILED"
	StatusRejected             Status = "REJECTED"
	StatusBlocked              Status = "BLOCKED"
	StatusCancelled            Status = "CANCELED"
	StatusCancelling           Status = "CANCELING"
	StatusConfirming           Status = "CONFIRMING"
	StatusBroadcasting         Status = "BROADCASTING"
	StatusPending3rdParty      Status = "PENDING_3RD_PARTY"
	StatusPendingSignature     Status = "PENDING_SIGNATURE"
	StatusQueued               Status = "QUEUED"
	StatusPendingAuthorization Status = "PENDING_AUTHORIZATION"
	StatusPendingAmlScreening  Status = "PENDING_AML_SCREENING"
	StatusSubmitted            Status = "SUBMITTED"
)

func (s Status) Completed() bool {
	return s == StatusCompleted
}

func (s Status) Failed() bool {
	return map[Status]bool{
		StatusFailed:     true,
		StatusRejected:   true,
		StatusBlocked:    true,
		StatusCancelled:  true,
		StatusCancelling: true,
	}[s]
}

type createTransactionRequest struct {
	Operation          string           `json:"operation"`
	Note               string           `json:"note,omitempty"`
	ExternalTxID       string           `json:"externalTxId,omitempty"`
	AssetID            string           `json:"assetId,omitempty"`
	Source             source           `json:"source"`
	Destination        *destination     `json:"destination,omitempty"`
	Destinations       []any            `json:"destinations,omitempty"`
	CustomerRefID      string           `json:"customerRefId,omitempty"`
	Amount             string           `json:"amountAll,omitempty"`
	TreatAsGrossAmount bool             `json:"treatAsGrossAmount,omitempty"`
	ForceSweep         bool             `json:"forceSweep,omitempty"`
	FeeLevel           string           `json:"feeLevel,omitempty"`
	Fee                string           `json:"fee,omitempty"`
	PriorityFee        string           `json:"priorityFee,omitempty"`
	MaxFee             string           `json:"maxFee,omitempty"`
	GasLimit           string           `json:"gasLimit,omitempty"`
	GasPrice           string           `json:"gasPrice,omitempty"`
	NetworkFee         string           `json:"networkFee,omitempty"`
	ReplaceTxByHash    string           `json:"replaceTxByHash,omitempty"`
	ExtraParameters    *extraParameters `json:"extraParameters,omitempty"`
}

type source struct {
	Type     string `json:"type"`
	SubType  string `json:"subType,omitempty"`
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	WalletID string `json:"walletId,omitempty"`
}

type destination struct {
	Type           string          `json:"type"`
	SubType        string          `json:"subType,omitempty"`
	ID             string          `json:"id,omitempty"`
	Name           string          `json:"name,omitempty"`
	WalletID       string          `json:"walletId,omitempty"`
	OneTimeAddress *oneTimeAddress `json:"oneTimeAddress,omitempty"`
}

type oneTimeAddress struct {
	Address string `json:"address,omitempty"`
	Tag     string `json:"tag,omitempty"`
}

type extraParameters struct {
	RawMessageData rawMessageData `json:"rawMessageData"`
}

type rawMessageData struct {
	Messages  []unsignedRawMessage `json:"messages"`
	Algorithm string               `json:"algorithm,omitempty"`
}

type unsignedRawMessage struct {
	Content string `json:"content"`
}

type transaction struct {
	ID                            string              `json:"id"`
	ExternalTxID                  string              `json:"externalTxId,omitempty"`
	Status                        Status              `json:"status"`
	SubStatus                     string              `json:"subStatus,omitempty"`
	TxHash                        string              `json:"txHash"`
	Operation                     string              `json:"operation"`
	Note                          string              `json:"note,omitempty"`
	AssetID                       string              `json:"assetId,omitempty"`
	Source                        source              `json:"source"`
	SourceAddress                 string              `json:"sourceAddress,omitempty"`
	Tag                           string              `json:"tag,omitempty"`
	Destination                   *destination        `json:"destination"`
	Destinations                  []any               `json:"destinations,omitempty"`
	DestinationAddress            string              `json:"destinationAddress,omitempty"`
	DestinationAddressDescription string              `json:"destinationAddressDescription,omitempty"`
	DestinationTag                string              `json:"destinationTag,omitempty"`
	ContractCallDecodedData       any                 `json:"contractCallDecodedData,omitempty"`
	AmountInfo                    *amountInfo         `json:"amountInfo,omitempty"`
	TreatAsGrossAmount            bool                `json:"treatAsGrossAmount"`
	FeeInfo                       *feeInfo            `json:"feeInfo,omitempty"`
	FeeCurrency                   string              `json:"feeCurrency,omitempty"`
	NetworkRecords                []networkRecords    `json:"networkRecords,omitempty"`
	CreatedAt                     int                 `json:"createdAt"`
	LastUpdated                   int                 `json:"lastUpdated"`
	CreatedBy                     string              `json:"createdBy"`
	SignedBy                      []string            `json:"signedBy"`
	RejectedBy                    string              `json:"rejectedBy"`
	AuthorizationInfo             *authorizationInfo  `json:"authorizationInfo"`
	ExchangeTxID                  string              `json:"exchangeTxId,omitempty"`
	CustomerRefID                 string              `json:"customerRefId,omitempty"`
	AmlScreeningResult            *amlScreeningResult `json:"amlScreeningResult,omitempty"`
	ExtraParameters               map[string]any      `json:"extraParameters,omitempty"`
	SignedMessages                []signedMessage     `json:"signedMessages"`
	NumOfConfirmations            int                 `json:"numOfConfirmations"`
	BlockInfo                     *blockInfo          `json:"blockInfo"`
	Index                         int                 `json:"index"`
	RewardInfo                    *rewardInfo         `json:"rewardInfo,omitempty"`
	SystemMessages                *systemMessages     `json:"systemMessages,omitempty"`
	AddressType                   string              `json:"addressType"`
}

// Pubkey0 returns the public key of the first signed message in the transaction.
func (t transaction) Pubkey0() (*ecdsa.PublicKey, error) {
	if len(t.SignedMessages) != 1 {
		return nil, errors.New("unexpected number of signed messages", "count", len(t.SignedMessages))
	}

	msg := t.SignedMessages[0]

	pk, err := hex.DecodeString(msg.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "decode public key")
	}

	pubkey, err := crypto.DecompressPubkey(pk)
	if err != nil {
		return nil, errors.Wrap(err, "decompress public key")
	}

	return pubkey, nil
}

// Sig0 returns the signature (Ethereum RSV format) of the first signed message in the transaction.
func (t transaction) Sig0() ([65]byte, error) {
	if len(t.SignedMessages) != 1 {
		return [65]byte{}, errors.New("unexpected number of signed messages", "count", len(t.SignedMessages))
	}

	msg := t.SignedMessages[0]

	// FullSig field is [R || S] in hex format.
	sig, err := hex.DecodeString(msg.Signature.FullSig)
	if err != nil {
		return [65]byte{}, errors.Wrap(err, "decode signature")
	} else if len(sig) != 64 {
		return [65]byte{}, errors.New("unexpected signature length", "length", len(sig))
	}

	// V is either 0 or 1
	sig = append(sig, byte(msg.Signature.V))

	return [65]byte(sig), nil
}

type amountInfo struct {
	Amount          string `json:"amount,omitempty"`
	RequestedAmount string `json:"requestedAmount,omitempty"`
	NetAmount       string `json:"netAmount,omitempty"`
	AmountUSD       string `json:"amountUSD,omitempty"`
}

type feeInfo struct {
	NetworkFee string `json:"networkFee,omitempty"`
	ServiceFee string `json:"serviceFee,omitempty"`
	GasPrice   string `json:"gasPrice,omitempty"`
}

type amlScreeningResult struct {
	Provider string `json:"provider,omitempty"`
	Payload  any    `json:"payload,omitempty"`
}

type networkRecords struct {
	Source             source       `json:"source"`
	Destination        *destination `json:"destination,omitempty"`
	TxHash             string       `json:"txHash,omitempty"`
	NetworkFee         string       `json:"networkFee,omitempty"`
	AssetID            string       `json:"assetId,omitempty"`
	NetAmount          string       `json:"netAmount,omitempty"`
	IsDropped          bool         `json:"isDropped,omitempty"`
	Type               string       `json:"type,omitempty"`
	DestinationAddress string       `json:"destinationAddress,omitempty"`
	SourceAddress      string       `json:"sourceAddress,omitempty"`
	AmountUSD          string       `json:"amountUSD,omitempty"`
	Index              int          `json:"index"`
	RewardInfo         *rewardInfo  `json:"rewardInfo,omitempty"`
}

type rewardInfo struct {
	SrcRewards  string `json:"srcRewards"`
	DestRewards string `json:"destRewards"`
}

type signature struct {
	FullSig string `json:"fullSig"`
	R       string `json:"r"`
	S       string `json:"s"`
	V       int    `json:"v"`
}
type signedMessage struct {
	Content        string    `json:"content"`
	Algorithm      string    `json:"algorithm"`
	DerivationPath []int     `json:"derivationPath"`
	Signature      signature `json:"signature"`
	PublicKey      string    `json:"publicKey"`
}
type blockInfo struct {
	BlockHeight string `json:"blockHeight"`
	BlockHash   string `json:"blockHash"`
}

type users struct {
	AdditionalProp string `json:"additionalProp"`
}
type groups struct {
	Th    int   `json:"th"`
	Users users `json:"users"`
}
type authorizationInfo struct {
	AllowOperatorAsAuthorizer bool     `json:"allowOperatorAsAuthorizer"`
	Logic                     string   `json:"logic"`
	Groups                    []groups `json:"groups"`
}

type createTransactionResponse struct {
	ID             string          `json:"id"`
	Status         string          `json:"status"`
	SystemMessages *systemMessages `json:"systemMessages,omitempty"`
}

type systemMessages struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Asset struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	ContractAddress string `json:"contractAddress,omitempty"`
	NativeAsset     string `json:"nativeAsset,omitempty"`
	Decimals        int    `json:"decimals,omitempty"`
}

type pubkeyResponse struct {
	Algorithm      string `json:"algorithm"`
	DerivationPath []int  `json:"derivationPath"`
	PublicKey      string `json:"publicKey"`
}

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type vaultsResponse struct {
	Accounts []account
	Paging   paging
}

type account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type paging struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
