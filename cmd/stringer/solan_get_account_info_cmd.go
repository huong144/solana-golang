package stringer

import (
	"github.com/spf13/cobra"
	"time"
)

var solanaAccountInfoCommand = &cobra.Command{
	Use:   "account-info",
	Short: "Fetching new data from solana network",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		for true {
			testData := make([]string, 1)
			testData[0] = "2q5gTDddf2NHZwYqDkS94Trf8vEX5D5mjEVMFvZTmMMsodxBCqJnRexx5sJEbBTy5SBv1cjGiuPxGNKk8pAd59YV"
			//go services.DecodeTx(testData)
			time.Sleep(2 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(solanaAccountInfoCommand)
}
