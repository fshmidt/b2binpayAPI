# B2BINPAY API Library

<!-- Table of Contents -->
# Table of Contents

1. [Description](#description)
2. [Implementation](#implementation)
3. [Installation](#installation)
4. [Usage](#usage)

[//]: # (1. [Examples]&#40;#examples&#41;)
<!-- End of Table of Contents -->

# Description

This is a Golang library for working with the B2BINPAY API. It allows you to easily create payments, retrieve account information, perform status check requests, and other payment-related operations.

# Implementation

- Bearer Auth

# Installation

To install the library, use the following command:

```go get github.com/fshmidt/b2binpay```

# Usage

To use the library, import it into your Go project:

```go import "github.com/fshmidt/b2binpay"```

Next, create a new client using your B2BINPAY account credentials:

```client := b2binpay.NewClient(apiKey, apiSecret, testMode)```

The testMode parameter specifies whether to use the B2BINPAY test API or the live API.

Credentials are saved in file to dercrease the number of API request. You can check if you already have one:

```
if c.GotCredentialsFile() {
    fmt.Println("Found & parsed saved credentials. No need for Auth request.")
} else {
    err := c.GetAuthToken()
    if err != nil {
        fmt.Println("Error obtaining auth token:", err)
        return
    }
}
```
In case you have expired access token you can check it and renew one:
```	
if c.IsAccessTokenExpired() {
    err := c.RefreshAuthToken()
    if err != nil {
        fmt.Println("Error refreshing auth token:", err)
        fmt.Println("New Authorization")
        err = c.GetAuthToken()
    }
}
```
You can now use the client to get wallet info:
```
walet, err := c.GetWallet(walletId)
	if err != nil {
		// Handle error
    }
```
You can also use the client to retrieve transfer information:
```
transfer, err := c.GetTransfer(transferId)
if err != nil {
    // Handle error
    } 
```
To check the rates, you can use the GetRates method:
```
rates, err := c.GetRates()
	if err != nil {
		fmt.Println("RATES ERR rates: ", err)
    }
```
To get invoice details use GetInvoice method:
```
invoice, err := c.GetInvoice(invoiceId)
	if err != nil {
		// handle error
    }  
```
You also can create invoice with CreateInvoice:
```
invoice, err := c.CreateInvoice(walletID, label, confirmationsNeeded, trackingID, timeLimit, inaccuracy, targetAmountRequested, currencyID)
	if err != nil {
		// handle error
	}
```
For getting payment info use GetPayout method:
```
payout, err := c.GetPayout(payoutId)
	if err != nil {
		// handle error
	}
```
For creating payment you'll need to pass  PayoutResponse object to CreatePayout method:
```
payout := &b2binpay.PayoutResponse{
    b2binpay.Payout{
        Type: "payout-calculation",
        Attributes: b2binpay.PayoutAttributes{
            Amount:    "0.05",
            ToAddress: "2N3Ac2cZzRVoqfJGu1bFaAebq3izTgr1WLv",
        },
        Relationships: b2binpay.PayoutRelationships{
            Wallet: b2binpay.Wallet{
                Data: b2binpay.WalletData{
                    Type: "wallet",
                    ID:   "13",
                },
            },
            Currency: b2binpay.Currency{
                Data: b2binpay.CurrencyData{
                    Type: "currency",
                    ID:   "1000",
                },
            },
        },
    },
}
payoutCreated, err := c.CalculatePayout(*payout)
if err != nil {
    // handle error
}
```