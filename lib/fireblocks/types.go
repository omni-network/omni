package fireblocks

type createTransactionRequest struct {
	Operation          string           `json:"operation"`
	Note               string           `json:"note,omitempty"`
	ExternalTxID       string           `json:"externalTxId,omitempty"`
	AssetID            string           `json:"assetId,omitempty"`
	Source             source           `json:"source"`
	Destination        *destination     `json:"destination,omitempty"`
	Destinations       []destinations   `json:"destinations,omitempty"`
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

type destinations struct {
	Amount      string      `json:"amount"`
	Destination destination `json:"destination"`
}

type oneTimeAddress struct {
	Address string `json:"address,omitempty"`
	Tag     string `json:"tag,omitempty"`
}

type extraParameters struct {
	RawMessageData rawMessageData `json:"rawMessageData"`
}

type rawMessageData struct {
	Messages  []UnsignedRawMessage `json:"messages"`
	Algorithm string               `json:"algorithm,omitempty"`
}

type UnsignedRawMessage struct {
	Content        string `json:"content"`
	DerivationPath []int  `json:"derivationPath,omitempty"`
}

type TransactionResponse struct {
	ID             string          `json:"id"`
	Status         string          `json:"status"`
	SystemMessages *SystemMessages `json:"systemMessages,omitempty"`
}

type SystemMessages struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
