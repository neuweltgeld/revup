package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"revup/x/revup/types"
)

var _ = strconv.Itoa(0)

func CmdCreateReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-review [product] [rate] [comment]",
		Short: "Broadcast message createReview",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProduct := args[0]
			argRate := args[1]
			argComment := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateReview(
				clientCtx.GetFromAddress().String(),
				argProduct,
				argRate,
				argComment,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
