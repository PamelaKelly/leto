package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Leto",
	Long:  `All software has versions. This is Leto's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Let Household Helper v0.1 -- HEAD")
	},
}
