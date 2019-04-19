package nameshake

import (
  "github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
  cdc.RegisterConcrete(MsgSetName{}, "nameshake/SetName", nil)
  cdc.RegisterConcrete(MsgBuyName{}, "nameshake/BuyName", nil)
}
