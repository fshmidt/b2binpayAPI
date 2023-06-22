package b2binpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func (c *Client) GetPayout(payoutID string) (*Payout, error) {
	url := fmt.Sprintf("%spayout/%s", c.APIAddress, payoutID)

	req, err := http.NewRequest("GET", url, nil)
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
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	var payoutResponse PayoutResponse
	err = json.NewDecoder(resp.Body).Decode(&payoutResponse)
	if err != nil {

		return nil, err
	}

	return &payoutResponse.Data, nil
}

func (c *Client) CreatePayout(payout PayoutResponse) (*PayoutResponse, error) {
	// Генерация уникального idempotency ключа
	idempotencyKey := generateUUID()

	// Создание HTTP запроса
	url := fmt.Sprintf("%spayout/", c.APIAddress) // Замените baseURL на фактическую базовую URL

	requestBodyBytes, err := json.Marshal(payout)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Idempotency-Key", idempotencyKey)

	// Отправка HTTP запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Обработка HTTP ответа
	if resp.StatusCode != http.StatusCreated {
		// Возникла ошибка при создании платежа
		return nil, fmt.Errorf("failed to create payout: %s", resp.Status)
	}

	// Чтение и декодирование JSON ответа
	var payoutResponse PayoutResponse
	err = json.NewDecoder(resp.Body).Decode(&payoutResponse)
	if err != nil {
		return nil, err
	}

	return &payoutResponse, nil
}

func generateUUID() string {
	return uuid.New().String()
}

func (c *Client) CalculatePayout(payout PayoutResponse) (*PayoutCalculationResponse, error) {

	// Создание HTTP запроса
	url := fmt.Sprintf("%spayout/calculate/", c.APIAddress) // Замените baseURL на фактическую базовую URL

	requestBodyBytes, err := json.Marshal(payout)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", contentType)

	// Отправка HTTP запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Обработка HTTP ответа
	if resp.StatusCode != http.StatusCreated {
		// Возникла ошибка при создании платежа
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("ERR CALC ", string(body))
		return nil, fmt.Errorf("failed to create payout: %s", resp.Status)
	}

	// Чтение и декодирование JSON ответа
	var payoutCalculation PayoutCalculationResponse
	err = json.NewDecoder(resp.Body).Decode(&payoutCalculation)
	if err != nil {
		return nil, err
	}

	return &payoutCalculation, nil
}
