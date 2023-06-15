package b2binpay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetRates() ([]Rate, error) {
	url := fmt.Sprintf("%srates/", c.APIAddress)

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
		return nil, fmt.Errorf("failed to get rates: %s", resp.Status)
	}

	var ratesResponse RatesResponse
	err = json.Unmarshal(body, &ratesResponse)
	if err != nil {
		return nil, err
	}

	return ratesResponse.Data, nil
}
