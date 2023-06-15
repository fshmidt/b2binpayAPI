package b2binpay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Метод для получения информации о валюте
func (c *Client) GetCurrency(currencyID string) (*Currency, error) {
	url := fmt.Sprintf("%scurrency/%s", c.APIAddress, currencyID)

	req, err := http.NewRequest("GET", url, nil)
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
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get currency information: %s", resp.Status)
	}
	var currency Currency
	err = json.Unmarshal(body, &currency)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}
