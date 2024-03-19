package fireblocks

type CreateTransactionRequest struct {
	Operation          string          `json:"operation"`
	Note               string          `json:"note,omitempty"`
	ExternalTxID       string          `json:"externalTxId,omitempty"`
	AssetID            string          `json:"assetId,omitempty"`
	Source             Source          `json:"source"`
	Destination        *Destination    `json:"destination,omitempty"`
	Destinations       []Destinations  `json:"destinations,omitempty"`
	CustomerRefID      string          `json:"customerRefId,omitempty"`
	Amount             string          `json:"amountAll,omitempty"`
	TreatAsGrossAmount bool            `json:"treatAsGrossAmount,omitempty"`
	ForceSweep         bool            `json:"forceSweep,omitempty"`
	FeeLevel           string          `json:"feeLevel,omitempty"`
	Fee                string          `json:"fee,omitempty"`
	PriorityFee        string          `json:"priorityFee,omitempty"`
	MaxFee             string          `json:"maxFee,omitempty"`
	GasLimit           string          `json:"gasLimit,omitempty"`
	GasPrice           string          `json:"gasPrice,omitempty"`
	NetworkFee         string          `json:"networkFee,omitempty"`
	ReplaceTxByHash    string          `json:"replaceTxByHash,omitempty"`
	ExtraParameters    *RawMessageData `json:"extraParameters,omitempty"`
}

type Source struct {
	Type     string `json:"type"`
	SubType  string `json:"subType,omitempty"`
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	WalletID string `json:"walletId,omitempty"`
}

type Destination struct {
	Type           string          `json:"type"`
	SubType        string          `json:"subType,omitempty"`
	ID             string          `json:"id,omitempty"`
	Name           string          `json:"name,omitempty"`
	WalletID       string          `json:"walletId,omitempty"`
	OneTimeAddress *OneTimeAddress `json:"oneTimeAddress,omitempty"`
}

type Destinations struct {
	Amount      string      `json:"amount"`
	Destination Destination `json:"destination"`
}

type OneTimeAddress struct {
	Address string `json:"address,omitempty"`
	Tag     string `json:"tag,omitempty"`
}

type RawMessageData struct {
	Messages  []UnsignedRawMessage `json:"messages"`
	Algorithm string               `json:"algorithm,omitempty"`
}

type UnsignedRawMessage struct {
	Content   string `json:"content"`
	Algorithm string `json:"algorithm,omitempty"`
}

type TransactionResponse struct {
	ID             string         `json:"id"`
	Status         string         `json:"status"`
	SystemMessages SystemMessages `json:"systemMessages"`
}

type SystemMessages struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
