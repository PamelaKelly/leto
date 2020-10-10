package cmd

import (
	"fmt"
	"os"

	"feargusogorman.com/leto/v1/pkg/trello"
	"github.com/spf13/cobra"
)

const (
	trelloAppEnvKey   = "TRELLO_APPKEY"
	trelloTokenEnvKey = "TRELLO_TOKEN"
)

var (
	boardName string
	user      string
	listName  string
	appKey    string
	token     string
)

func init() {
	// flags
	trelloCmd.PersistentFlags().StringVar(&user, "user", "", "Specifcy a user in Trello")
	trelloCmd.MarkPersistentFlagRequired("user")
	getBoardCmd.Flags().StringVar(&boardName, "name", "", "Specifcy a board in Trello")
	getListCmd.Flags().StringVar(&listName, "list", "", "Set a specific list on the board")
	getBoardCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(trelloCmd)
	trelloCmd.AddCommand(getBoardCmd)
	trelloCmd.AddCommand(getListCmd)
}

var trelloCmd = &cobra.Command{
	Use:   "trello",
	Short: "Interact with Trello via CLI",
	Long:  "Use this command to interact with Trello",
}

var getBoardCmd = &cobra.Command{
	Use:   "board",
	Short: "Gets a Trello board",
	RunE: func(cmd *cobra.Command, args []string) error {
		// validate requirements for instantiating a connection with the Trello Api
		appKey = os.Getenv(trelloAppEnvKey)
		token = os.Getenv(trelloTokenEnvKey)
		// Temporary simple validation
		if appKey == "" {
			return fmt.Errorf("required enviornment variable empty: %s", trelloAppEnvKey)
		}
		if token == "" {
			return fmt.Errorf("required enviornment variable empty: %s", trelloTokenEnvKey)
		}
		client, err := trello.NewClient(appKey, token, user)
		if err != nil {
			return fmt.Errorf("failed to instantiate Trello client with error: %s", err.Error())
		}
		board, err := client.GetBoard(boardName)
		if err != nil {
			return fmt.Errorf("error getting board from Trello: %s", err.Error())
		}
		fmt.Printf("Retrieved board %s for user %s\n", board.Name, user)
		return nil
	},
}

var getListCmd = &cobra.Command{
	Use:   "list",
	Short: "Gets a list from a Trello board",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := trello.NewClient(appKey, token, user)
		if err != nil {
			return fmt.Errorf("failed to instantiate Trello client with error: %s", err.Error())
		}
		list, err := client.GetList(boardName, listName)
		if err != nil {
			return fmt.Errorf("error gettings list %s from board %s", listName, boardName)
		}
		fmt.Printf("Retrieved list %s on board %s for user %s\n", list.Name, boardName, user)
		return nil
	},
}
