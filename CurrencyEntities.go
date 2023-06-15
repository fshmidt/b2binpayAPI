package b2binpay

type Currency struct {
	Data CurrencyData `json:"data"`
}

type CurrencyData struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Attributes    CurrencyAttributes `json:"attributes"`
	Relationships CurrencyRelations  `json:"relationships"`
}

type CurrencyAttributes struct {
	ISO                   int64  `json:"iso"`
	Name                  string `json:"name"`
	Alpha                 string `json:"alpha"`
	Exp                   int64  `json:"exp"`
	ConfirmationBlocks    int64  `json:"confirmation_blocks"`
	MinimalTransferAmount string `json:"minimal_transfer_amount"`
	BlockDelay            int64  `json:"block_delay"`
	Alias                 string `json:"alias"`
}

type CurrencyRelations struct {
	Parent *CurrencyParent `json:"parent"`
}
type CurrencyParent struct {
	Data CurrencyParentData `json:"data"`
}

type CurrencyParentData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
