package b2binpay

type PayoutResponse struct {
	Data Payout `json:"data"`
}

type Payout struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Attributes    PayoutAttributes    `json:"attributes"`
	Relationships PayoutRelationships `json:"relationships"`
}

type PayoutAttributes struct {
	Amount              string      `json:"amount"`
	Exp                 int         `json:"exp"`
	Address             string      `json:"address"`
	ToAddress           string      `json:"to_address"`
	TagType             string      `json:"tag_type"`
	Tag                 string      `json:"tag"`
	Destination         Destination `json:"destination"`
	TrackingID          interface{} `json:"tracking_id"`
	ConfirmationsNeeded interface{} `json:"confirmations_needed"`
	FeeAmount           string      `json:"fee_amount"`
	IsFeeIncluded       bool        `json:"is_fee_included"`
	Status              int         `json:"status"`
	CallbackURL         interface{} `json:"callback_url"`
}

type PayoutRelationships struct {
	Currency Currency `json:"currency"`
	Wallet   Wallet   `json:"wallet"`
}

type PayoutCalculationResponse struct {
	Data PayoutCalculationData `json:"data"`
}

type PayoutCalculationData struct {
	Type       string                      `json:"type"`
	ID         string                      `json:"id"`
	Attributes PayoutCalculationAttributes `json:"attributes"`
}

type PayoutCalculationAttributes struct {
	IsInternal bool                        `json:"is_internal"`
	Fee        PayoutCalculationFee        `json:"fee"`
	Commission PayoutCalculationCommission `json:"commission"`
}

type PayoutCalculationFee struct {
	Low        string `json:"low"`
	Medium     string `json:"medium"`
	High       string `json:"high"`
	DustAmount string `json:"dust_amount"`
	Currency   int    `json:"currency"`
}

type PayoutCalculationCommission struct {
	Amount   string `json:"amount"`
	Currency int    `json:"currency"`
}
