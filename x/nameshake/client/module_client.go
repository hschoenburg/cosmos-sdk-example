package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	nameshakecmd "github.com/hschoenburg/demoDapp/x/nameshake/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	nameshkQueryCmd := &cobra.Command{
		Use:   "nameshake",
		Short: "Querying commands for nameshake module",
	}

	nameshkQueryCmd.AddCommand(client.GetCommands(
		nameshakecmd.GetCmdResolveName(mc.storeKey, mc.cdc),
		nameshakecmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return nameshakeQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	nameshkTxCmd := &cobra.Command{
		Use:   "nameshake",
		Short: "Nameshake transaction subcommands",
	}

	nameshkTxCmd.AddCommand(clint.PostCommands(
		nameshakecmd.GetCmdBuyName(mc.cdc),
		nameshakecmd.GetCmdSetName(mc.cdc),
	)...)

	return nameshkTxCmd
}
