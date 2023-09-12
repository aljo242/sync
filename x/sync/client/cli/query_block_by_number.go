package cli

import (
	"strconv"

	"github.com/aljo242/sync/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBlockByNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-by-number [block-number]",
		Short: "Query blockByNumber",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBlockNumber := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBlockByNumberRequest{

				BlockNumber: reqBlockNumber,
			}

			res, err := queryClient.BlockByNumber(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
