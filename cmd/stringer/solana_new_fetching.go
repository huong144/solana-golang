package stringer

import (
	"github.com/spf13/cobra"
	"solana-crawl-service/Solana/services"
)

var solanaFetchNewCommand = &cobra.Command{
	Use:   "solana-fetch-new",
	Short: "Fetching new data from solana network",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		//var wg sync.WaitGroup
		//wg.Add(10)
		for true {
			services.GetSyncBlock()
			//defer wg.Done()

		}
		//wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(solanaFetchNewCommand)
}
