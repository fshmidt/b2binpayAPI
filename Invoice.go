package b2binpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetInvoice(id string) (Invoice, error) {
	url := fmt.Sprintf("%sdeposit/%s", c.APIAddress, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Invoice{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return Invoice{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Body)
		return Invoice{}, fmt.Errorf("failed to get invoice: %s", resp.Status)
	}

	var invoice Invoice
	err = json.NewDecoder(resp.Body).Decode(&invoice)
	if err != nil {
		return Invoice{}, err
	}

	return invoice, nil
}

func (c *Client) OptionsInvoice() (*http.Response, error) {
	options := &InvoiceOptions{
		Data: OptionsData{
			Type: "auth-token",
			Attributes: OptionsAttributes{
				Login:    c.Login,
				Password: c.Password,
			},
		},
	}
	jsonPayload, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	url := c.APIAddress + "deposit/"

	req, err := http.NewRequest("OPTIONS", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateInvoice(walletID string, label string, confirmationsNeeded int, trackingID string, timeLimit int, inaccuracy int, targetAmountRequested string, currencyID string) (*Invoice, error) {
	invoice := &DepositRequest{
		Data: DepositData{
			Type: "deposit",
			Attributes: DepositAttributes{
				Label:                 label,
				TrackingID:            trackingID,
				ConfirmationsNeeded:   confirmationsNeeded,
				TimeLimit:             timeLimit,
				Inaccuracy:            inaccuracy,
				TargetAmountRequested: targetAmountRequested,
			},
			Relationships: DepositRelationships{
				Wallet: Wallet{
					Data: WalletData{
						Type: "wallet",
						ID:   walletID,
					},
				},
				Currency: Currency{
					Data: CurrencyData{
						Type: "currency",
						ID:   currencyID,
					},
				},
			},
		},
	}

	jsonPayload, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}

	url := c.APIAddress + "deposit/"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		fmt.Println("RESP INV:", string(body))
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	var invoiceResponse InvoiceResponse
	err = json.NewDecoder(resp.Body).Decode(&invoiceResponse)
	if err != nil {
		return nil, err
	}

	return &invoiceResponse.Data, nil
}
