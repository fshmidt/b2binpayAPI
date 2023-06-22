package b2binpay

type InvoiceResponse struct {
	Data Invoice `json:"data"`
}
type Destination struct {
	AddressType string `json:"address_type"`
	Address     string `json:"address"`
}

type Invoice struct {
	ID                    string      `json:"id"`
	Status                int         `json:"status"`
	Address               string      `json:"address"`
	AddressType           string      `json:"address_type"`
	Label                 string      `json:"label"`
	TrackingID            string      `json:"tracking_id"`
	ConfirmationsNeeded   int         `json:"confirmations_needed"`
	TimeLimit             int         `json:"time_limit"`
	CallbackURL           string      `json:"callback_url"`
	Inaccuracy            string      `json:"inaccuracy"`
	TargetAmountRequested string      `json:"target_amount_requested"`
	RateRequested         string      `json:"rate_requested"`
	RateExpiredAt         string      `json:"rate_expired_at"`
	InvoiceUpdatedAt      string      `json:"invoice_updated_at"`
	PaymentPage           string      `json:"payment_page"`
	TargetPaid            string      `json:"target_paid"`
	SourceAmountRequested string      `json:"source_amount_requested"`
	TargetPaidPending     string      `json:"target_paid_pending"`
	Destination           Destination `json:"destination"`
	Currency              Currency    `json:"currency"`
	Wallet                Wallet      `json:"wallet"`
}

type DepositRequest struct {
	Data DepositData `json:"data"`
}

type DepositData struct {
	Type          string               `json:"type"`
	Attributes    DepositAttributes    `json:"attributes"`
	Relationships DepositRelationships `json:"relationships"`
}

type DepositAttributes struct {
	Label                 string `json:"label"`
	TrackingID            string `json:"tracking_id"`
	ConfirmationsNeeded   int    `json:"confirmations_needed"`
	TimeLimit             int    `json:"time_limit"`
	Inaccuracy            int    `json:"inaccuracy"`
	TargetAmountRequested string `json:"target_amount_requested"`
}

type DepositRelationships struct {
	Wallet   Wallet   `json:"wallet"`
	Currency Currency `json:"currency"`
}

type InvoiceOptions struct {
	Data OptionsData `json:"data"`
}

type OptionsData struct {
	Type       string            `json:"type"`
	Attributes OptionsAttributes `json:"attributes"`
}

type OptionsAttributes struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
