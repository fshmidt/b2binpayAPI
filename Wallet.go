package b2binpay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Метод для получения информации о кошельке

func (c *Client) GetWallet(walletID string) (*Wallet, error) {
	url := fmt.Sprintf("%swallet/%s", c.APIAddress, walletID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)

	fmt.Println("Used token: ", "Bearer "+c.RefreshToken)
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
		return nil, fmt.Errorf("failed to get wallet information: %s", resp.Status)
	}

	var wallet Wallet
	err = json.Unmarshal(body, &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
