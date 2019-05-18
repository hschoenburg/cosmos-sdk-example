package nameshake

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ValidateGenesis validates the provided module genesis state to ensure the
// expected invariants holds. (i.e. params in correct bounds, no duplicate validators)
func ValidateGenesis(data GenesisState) error {

	for _, record := range data.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("Invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Price == nil {
			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: nil,
		Params:       DefaultParams(),
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {

	for _, record := range data.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	panic("not implemented")
}
