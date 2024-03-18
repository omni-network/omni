package fireblocks

type CreateTransactionRequest struct {
	Operation       string         `json:"operation"`
	Note            string         `json:"note"`
	Source          Source         `json:"source"`
	Destination     Destination    `json:"destination"`
	Destinations    []Destinations `json:"destinations"`
	AssetID         string         `json:"assetId"`
	Amount          string         `json:"amount"`
	CustomerRefID   string         `json:"customerRefId"`
	ExtraParameters RawMessageData `json:"extraParameters"`
}

type Source struct {
	Type     string `json:"type"`
	SubType  string `json:"subType"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	WalletID string `json:"walletId"`
}

type Destination struct {
	Type           string         `json:"type"`
	SubType        string         `json:"subType"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	WalletID       string         `json:"walletId"`
	OneTimeAddress OneTimeAddress `json:"oneTimeAddress"`
}

type Destinations struct {
	Amount      string      `json:"amount"`
	Destination Destination `json:"destination"`
}

type OneTimeAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

type RawMessageData struct {
	Messages []UnsignedRawMessage `json:"messages"`
}

type UnsignedRawMessage struct {
	Content string `json:"content"`
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
