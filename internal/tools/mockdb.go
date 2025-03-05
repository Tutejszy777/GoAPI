package mockdb

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]CoinsDetails{
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

func (d *mockDB)