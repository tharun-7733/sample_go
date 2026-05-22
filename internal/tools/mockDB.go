package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Tharun": {
		AuthToken: "123ABC",
		Username:  "Tharun",
	},
	"Tej": {
		AuthToken: "456DEF",
		Username:  "Tej",
	},
	"Aksh": {
		AuthToken: "789GHI",
		Username:  "Aksh",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Tharun": {
		Coins:   100,
		Username: "Tharun",
	},
	"Tej": {
		Coins:   200,
		Username: "Tej",
	},
	"Aksh": {
		Coins:   300,
		Username: "Aksh",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}