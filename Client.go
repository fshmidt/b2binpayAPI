package b2binpay

import (
	"time"
)

const (
	contentType = "application/vnd.api+json"
	mockAPI     = "https://api-sandbox.b2binpay.com/"
	realAPI     = "https://api.b2binpay.com/\n\n"
)

type Client struct {
	APIAddress       string
	Login            string
	Password         string
	AccessToken      string
	RefreshToken     string
	RefreshExpiredAt time.Time
	AccessExpiredAt  time.Time
}

func NewClient(login, password string, testmode bool) Client {
	client := Client{Login: login, Password: password}
	if testmode == true {
		client.APIAddress = mockAPI
	} else {
		client.APIAddress = realAPI
	}
	return client
}
