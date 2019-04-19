package main

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
  //"github.com/hschoenburg/nameshake"
	"github.com/tendermint/tendermint/libs/log"
  "github.com/cosmos/cosmos-sdk/codec"
	//sdk "github.com/cosmos/cosmos-sdk/types"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "nameshake"
)

type nameShakeApp struct {
  BaseApp *bam.BaseApp
  Cdc codec.Codec
}

func newNameShakeApp(logger log.Logger, db dbm.DB) *nameShakeApp {

  // what is this can we explain it better?
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameShakeApp{
		BaseApp: bApp,
		Cdc:     cdc,
	}

	return app

}
