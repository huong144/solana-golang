package stringer

import (
	"github.com/spf13/cobra"
	"solana-crawl-service/Solana/services"
	"time"
)

var solanaFetchOldCommand = &cobra.Command{
	Use:   "solana-fetch-old",
	Short: "Fetching new data from solana network",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		for true {
			go services.GetSyncOldBlock()
			time.Sleep(2 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(solanaFetchOldCommand)
}
