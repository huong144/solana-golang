package stringer

import (
	"github.com/spf13/cobra"
	"solana-crawl-service/Solana/schema"
)

var pgDBMigration = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate model database",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		schema.AutoMigration()
	},
}

func init() {
	rootCmd.AddCommand(pgDBMigration)
}
