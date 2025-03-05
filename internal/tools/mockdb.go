package mockdb

import (
	"time"
)

type mockDB struct{}

var mockCoinDetails = map[string]CoinsDetails{
	"alex": {
		Coins: 100,
		Username: "Alex"
	},
	"marie": {
		Coins: 100,
		Username: "marie"
	},
	"vladimir": {
		Coins: 100,
		Username: "vladimir"
	},
}

var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken: "123ABC"
		Username: "Alex"
	},
	"marie": {
		AuthToken: "456DEF"
		Username: "marie"
	},
	"vladimir": {
		AuthToken: "321CBA"
		Username: "vladimir"
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second*1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinsDetails {
	time.Sleep(time.Second*1)

	var clientData = CoinsDetails{}
	clientData, ok = mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}