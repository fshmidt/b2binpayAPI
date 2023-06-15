package main

import (
	"b2binpay"
	"fmt"
)

// Пример использования

func main() {
	// Получение токена авторизации
	login := "teo@fanated.com"
	password := "7777777"
	c := b2binpay.NewClient(login, password, true)

	if c.GotCredentialsFile() {
		fmt.Println("Found & parsed saved credentials. No need for Auth request.")
	} else {
		err := c.GetAuthToken()
		if err != nil {
			fmt.Println("Error obtaining auth token:", err)
			return
		}
	}
	fmt.Println("Access Token:", c.AccessToken, "\nRefresh:", c.RefreshToken, "\nExpired:", c.RefreshExpiredAt)

	//Обновление токена авторизации
	if c.IsAccessTokenExpired() {
		err := c.RefreshAuthToken()
		if err != nil {
			fmt.Println("Error refreshing auth token:", err)
			fmt.Println("New Authorization")
			err = c.GetAuthToken()
		} else {
			fmt.Println("New Access Token FROM REFRESHING:", c.AccessToken, "\nRefresh:", c.RefreshToken, "\nExpired:", c.RefreshExpiredAt)
		}
	}

	//Получение сведений о кошельке
	walet, errw := c.GetWallet("1")
	if errw != nil {
		fmt.Println("ERR: ", errw)
	} else {
		fmt.Println("WALLET: ", walet.Data.ID, walet.Data.Attributes.Status, walet.Data.Attributes.BalanceConfirmed)
	}

	//Получение сведений о валюте:
	cur, errc := c.GetCurrency("2015")
	if errc != nil {
		fmt.Println("ERR: ", errc)
	} else {
		fmt.Println("CURRENCY: ", cur.Data.ID, cur.Data.Attributes.Alias, cur.Data.Attributes.Alpha, cur.Data.Attributes.MinimalTransferAmount)
	}

	// Получение сведений о трансфере

	transfer, err := c.GetTransfer("80")
	if err != nil {
		fmt.Println("ERR: ", err)
	} else {
		fmt.Println("TRANSFER: ", transfer.Data.Attributes.Amount, transfer.Data.Attributes.Risk)
	}

	// Rates

	rates, err := c.GetRates()
	if err != nil {
		fmt.Println("ERR rates: ", err)
	} else {
		fmt.Println("Rates: ", rates)
	}
}
