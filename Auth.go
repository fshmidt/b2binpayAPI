package b2binpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	authEndpoint    = "token/"
	refreshEndpoint = "token/refresh/"
)

// Метод для получения токена авторизации
func (c *Client) GetAuthToken() error {

	authReq := AuthRequest{
		Data: AuthRequestData{
			Type: "auth-token",
			Attributes: AuthRequestAttributes{
				Login:    c.Login,
				Password: c.Password,
			},
		},
	}

	authReqJSON, err := json.Marshal(authReq)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.APIAddress+authEndpoint, contentType, bytes.NewBuffer(authReqJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to obtain auth token: %s", resp.Status)
	}

	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return err
	}

	refreshExpiredAt, err := time.Parse(time.RFC3339, authResp.Data.Attributes.RefreshExpiredAt)
	if err != nil {
		return err
	}
	accessExpiredAt, err := time.Parse(time.RFC3339, authResp.Data.Attributes.AccessExpiredAt)
	if err != nil {
		return err
	}

	c.AccessToken, c.RefreshToken, c.RefreshExpiredAt, c.AccessExpiredAt = authResp.Data.Attributes.Access, authResp.Data.Attributes.Refresh, refreshExpiredAt, accessExpiredAt
	err = c.SaveCredentials()
	if err != nil {
		return err
	}

	return nil
}
func (c *Client) SaveCredentials() error {
	fileName := fmt.Sprintf("%s.json", c.Login)

	_, err := os.Stat(fileName)

	if os.IsExist(err) {
		fmt.Println("Removing previous credentials")
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal("Can't delete previous credentials.", err)
		}
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	clientData := struct {
		AccessToken      string    `json:"accessToken"`
		RefreshToken     string    `json:"refreshToken"`
		RefreshExpiredAt time.Time `json:"refreshExpiredAt"`
		AccessExpiredAt  time.Time `json:"accessExpiredAt"`
	}{
		AccessToken:      c.AccessToken,
		RefreshToken:     c.RefreshToken,
		RefreshExpiredAt: c.RefreshExpiredAt,
		AccessExpiredAt:  c.AccessExpiredAt,
	}
	err = json.NewEncoder(file).Encode(&clientData)
	if err != nil {
		return err
	}
	return nil
}

// Метод для сохранения токена авторизации

// Метод для обновления токена авторизации
func (c *Client) RefreshAuthToken() error {
	refreshReq := RefreshRequest{
		Data: RefreshRequestData{
			Type: "auth-token",
			Attributes: RefreshRequestAttributes{
				Refresh: c.RefreshToken,
			},
		},
	}

	refreshReqJSON, err := json.Marshal(refreshReq)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.APIAddress+refreshEndpoint, contentType, bytes.NewBuffer(refreshReqJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to refresh auth token: %s", resp.Status)
	}

	var refreshResp RefreshResponse
	err = json.Unmarshal(body, &refreshResp)
	if err != nil {
		return err
	}

	refreshExpiredAt, err := time.Parse(time.RFC3339, refreshResp.Data.Attributes.RefreshExpiredAt)
	if err != nil {
		return err
	}
	accessExpiredAt, err := time.Parse(time.RFC3339, refreshResp.Data.Attributes.AccessExpiredAt)
	if err != nil {
		return err
	}
	c.AccessToken, c.RefreshToken, c.RefreshExpiredAt, c.AccessExpiredAt = refreshResp.Data.Attributes.Access, refreshResp.Data.Attributes.Refresh, refreshExpiredAt, accessExpiredAt

	return nil
}

// Метод для проверки срока годности RefreshExpiredAt
func (c *Client) IsRefreshTokenExpired() bool {
	return time.Now().After(c.RefreshExpiredAt)
}

// Метод для проверки срока годности AccessExpiredAt
func (c *Client) IsAccessTokenExpired() bool {
	return time.Now().After(c.AccessExpiredAt)
}

// Метод для проверки наличия файла с данными в рабочей папке
func (c *Client) GotCredentialsFile() bool {
	fileName := c.Login + ".json"
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return false
	}

	file, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return false
	}

	if c.IsRefreshTokenExpired() {
		file.Close()
		fmt.Println("Removing expired credentials")
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal("Can't delete expired client.", err)
		}

		return false
	}
	return true
}
