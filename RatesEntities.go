package b2binpay

type RatesResponse struct {
	Data []Rate `json:"data"`
}
type Rate struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes RateAttributes `json:"attributes"`
}

type RateAttributes struct {
	Left      string `json:"left"`
	Right     string `json:"right"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	Exp       int    `json:"exp"`
	CreatedAt string `json:"created_at"`
	ExpiredAt string `json:"expired_at"`
}
