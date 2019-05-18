package nameshake

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

type GenesisState struct {
	WhoisRecords []Whois `json:"whois_records"`
	Params       Params  `json:"nameshake_params"`
}
