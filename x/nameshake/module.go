package nameshake

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	ModuleName = "nameshake"
	RouterKey  = "nameshake"
)

func NewAppModule(k Keeper) AppModule {
	return AppModule{keeper: k}
}

type AppModuleBasic struct{}

func (am AppModuleBasic) Name() string {
	return ModuleName
}

func (am AppModuleBasic) RegisterCodec(*codec.Codec) {
	panic("not implemented")
}

// get raw genesis raw message for testing
func (am AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

//safetyguard function. testing is this JSON even Marshall-able
func (am AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

func (am AppModule) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// staking is a great reference for this
// some modules just store parameters. others store accounts and more data
func (am AppModule) InitGenesis(ctx types.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return InitGenesis(ctx, am.keeper, genesisState)
}

// takes keepers, constructs a JSON object.
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

// staking invariants
//its really slow. Doesnt run every block
//invariants check
//staking has multiple invariant routers that it registers with the router
func (am AppModule) RegisterInvariants(types.InvariantRouter) {
	panic("not implemented")
}

func (am AppModule) Route() string {
	panic("not implemented")
}

func (am AppModule) NewHandler() types.Handler {
	panic("not implemented")
}

func (am AppModule) QuerierRoute() string {
	panic("not implemented")
}

func (am AppModule) NewQuerierHandler() types.Querier {
	panic("not implemented")
}

func (am AppModule) BeginBlock(types.Context, abci.RequestBeginBlock) types.Tags {
	panic("not implemented")
}

func (am AppModule) EndBlock(types.Context, abci.RequestEndBlock) ([]abci.ValidatorUpdate, types.Tags) {
	panic("not implemented")
}

// type check AppModule and AppModule basic
var (
	_ sdk.AppModule      = AppModule{}
	_ sdk.AppModuleBasic = AppModuleBasic{}
)
