package b2binpay

type Transfer struct {
	Data TransferData `json:"data"`
}

type TransferData struct {
	Type          string                `json:"type"`
	ID            string                `json:"id"`
	Attributes    TransferAttributes    `json:"attributes"`
	Relationships TransferRelationships `json:"relationships"`
}

type TransferAttributes struct {
	Confirmations int64  `json:"confirmations"`
	OpID          int64  `json:"op_id"`
	OpType        int64  `json:"op_type"`
	RiskStatus    int64  `json:"risk_status"`
	Risk          int64  `json:"risk"`
	Status        int64  `json:"status"`
	Amount        string `json:"amount"`
	AmountTarget  string `json:"amount_target"`
	AmountCleared string `json:"amount_cleared"`
	Commission    string `json:"commission"`
	TxID          string `json:"txid"`
	Fee           string `json:"fee"`
	UserMessage   string `json:"user_message"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TransferRelationships struct {
	Currency CurrencyRelation `json:"currency"`
	Wallet   WalletRelation   `json:"wallet"`
	Parent   ParentRelation   `json:"parent"`
}

type CurrencyRelation struct {
	Data CurrencyData `json:"data"`
}

type WalletRelation struct {
	Data WalletData `json:"data"`
}

type ParentRelation struct {
	Data interface{} `json:"data"` // Can be null or a transfer object
}
