package nameshake


import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func NewHandler(keeper Keeper) sdk.Handler {
  return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

    switch msg:= msg.(type) {
      case MsgSetName:
        return handleMsgSetName(ctx, keeper, msg)
      case MsgBuyName:
        return handleMsgBuyName(ctx, keeper, msg)
      default:
        errMsg := fmt.Sprintf("Unrecognized nameshake Msg type: %v", msg.Type())
        return sdk.ErrUnknownRequest(errMsg).Result()
    }
  }
}

func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
  if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
    return sdk.ErrInfufficientCoins("Bid not high enough").Result()
  }

  if keeper.hasOwner(ctx, msg.Name) {
    _, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
    if err != nil {
      return sdk.ErrInfufficientCoins("Buyer does not have enough coins").Result()
    }

  } else {
    _, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
    if err != nil {
      return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
    }
  }
  keeper.SetOwner(ctx, msg.Name, msg.Buyer)
  keeper.SetPrice(ctx, msg.Name, msg.Bid)
  return sdk.Result{}

}






