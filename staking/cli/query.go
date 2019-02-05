package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// get checkpoint present in buffer
func GetValidatorInfo(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-validator-info",
		Short: "show validator information via validator address",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			validatorAddress := viper.GetString(FlagValidatorAddress)

			res, err := cliCtx.QueryStore(common.GetValidatorKey([]byte(validatorAddress)), "staker")
			if err != nil {
				fmt.Printf("Error fetching validator information from store, Error: %v ValidatorAddr: %v", err, validatorAddress)
				return err
			}

			var _validator types.Validator
			err = cdc.UnmarshalBinary(res, &_validator)
			if err != nil {
				fmt.Printf("Error unmarshalling validator , Error: %v", err)
				return err
			}
			return nil
		},
	}

	cmd.MarkFlagRequired(FlagValidatorAddress)
	return cmd
}