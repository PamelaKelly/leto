package cmd

import (
	"fmt"

	"feargusogorman.com/leto/v1/pkg/trello"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(trelloCmd)
	trelloCmd.AddCommand(getBoardCmd)
}

var trelloCmd = &cobra.Command{
	Use:   "trello",
	Short: "Interact with Trello via CLI",
	Long:  "Use this command to send requests to Trello to get, or post data",
}

var getBoardCmd = &cobra.Command{
	Use:   "board",
	Short: "Gets a Trello board",
	Run: func(cmd *cobra.Command, args []string) {
		board, err := trello.GetBoard()
		if err != nil {
			fmt.Printf("Error getting board from Trello: %s", err.Error())
		}
		fmt.Println(board)
	},
}
