package b2binpay

type Wallet struct {
	Data WalletData `json:"data"`
}

type WalletData struct {
	Type       string           `json:"type"`
	ID         string           `json:"id"`
	Attributes WalletAttributes `json:"attributes"`
	Relations  WalletRelations  `json:"relationships"`
}

type WalletAttributes struct {
	Status                int64             `json:"status"`
	CreatedAt             string            `json:"created_at"`
	BalanceConfirmed      string            `json:"balance_confirmed"`
	BalancePending        string            `json:"balance_pending"`
	BalanceUnusable       string            `json:"balance_unusable"`
	MinimalTransferAmount string            `json:"minimal_transfer_amount"`
	Destination           WalletDestination `json:"destination"`
}

type WalletDestination struct {
	AddressType string `json:"address_type"`
	Address     string `json:"address"`
}

type WalletRelations struct {
	Currency WalletCurrency `json:"currency"`
}

type WalletCurrency struct {
	Data WalletCurrencyData `json:"data"`
}

type WalletCurrencyData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
