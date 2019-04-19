package cli

import (
  "github.com/spf13/cobra"
  "github.com/cosmos/cosmos-sdk/client/context"
  "github.com/cosmos/cosmos-sdk/client/utils"
  "github.com/cosmos/cosmos-sdk/codec"
  "github.com/hschoenburg/nameshake/x/nameshake"

  sdk "github.com/cosmos/cosmos-sdk/types"
  authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func GetCmdBuyName (cdc *codec.Codec) *cobra.Command {
  return &cobra.Command{
    Use: "buy-name [name] [amount]",
    Short: "bid for name",
    Args: cobra.ExactArgs(2),
    RunE: func (cmd *cobra.Command, args []string) error {
      cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

      txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

      if err := cliCtx.EnsureAccountExists(); err != nil {
        return err
      }

      coins, err := sdk.ParseCoins(args[1])
      if err != nil {
        return err
      }

      msg := nameshake.NewMsgBuyName(args[0], coins, cliCtx.GetFromAddress())
      err := msg.ValidateBasic()
      if err != nil {
        return err
      }

      cliCtx.ParseResponse = true

      return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
    },
  }
}

func GetCmdSetName(cdc *codec.Codec) *cobra.Command {
  return &cobra.Command{
    Use: "set-name [name] [value]",
    Short: "set the value assosciated with name",
    Args: cobra.ExactArgs(2),
    RunE: func(cmd *cobra.Command, args []string) error {
      cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
      txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

      if err := cliCtx.EnsureAccountExists(); err != nil {
        return err
      }

      msg := nameshake.MsgSetName(args[0], args[1], cliCtx.GetFromAddress())
      err := msg.ValidateBasic()
      if err != nil {
        return err
      }

      cliCtx.PrintResponse = true

      return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
    },
  }
}




