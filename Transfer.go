package b2binpay

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetTransfer(id string) (*Transfer, error) {
	url := fmt.Sprintf("%stransfer/%s", c.APIAddress, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)
	//fmt.Println(url, contentType)

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
		return nil, fmt.Errorf("failed to get transfer information: %s", resp.Status)
	}

	var transfer Transfer

	err = json.Unmarshal(body, &transfer)
	if err != nil {

		return nil, err
	}
	fmt.Println("HERE OK")

	return &transfer, nil
}
