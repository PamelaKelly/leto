package cmd

import (
	"fmt"
	"os"

	"feargusogorman.com/leto/v1/pkg/trello"
	"github.com/spf13/cobra"
)

const (
	trelloAPIEnvKey   = "TRELLO_APIKEY"
	trelloTokenEnvKey = "TRELLO_TOKEN"
)

var board string
var user string

func init() {
	// flags
	trelloCmd.PersistentFlags().StringVar(&user, "user", "", "Specifcy a user in Trello")
	getBoardCmd.PersistentFlags().StringVar(&board, "name", "", "Specifcy a board in Trello")
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Driving as env vars will allow to securely provide these through a K8s deployment
		// as well as locally for ease
		apiKey := os.Getenv(trelloAPIEnvKey)
		token := os.Getenv(trelloTokenEnvKey)

		// Temporary simple validation
		if apiKey == "" {
			return fmt.Errorf("required enviornment variable empty: %s", trelloAPIEnvKey)
		}
		if token == "" {
			return fmt.Errorf("required enviornment variable empty: %s", trelloTokenEnvKey)
		}
		client, err := trello.NewClient(apiKey, token, user)
		if err != nil {
			return fmt.Errorf("failed in instantiate trello client with error: %s", err.Error())
		}
		// For the moment assuming the first arg is the board name
		board, err := client.GetBoard(board)
		if err != nil {
			return fmt.Errorf("error getting board from Trello: %s", err.Error())
		}
		fmt.Printf("Retrieved board: %s for user %s\n", board.Name, user)
		return nil
	},
}
