package cmd

import (
	"fmt"

	"feargusogorman.com/leto/v1/internal/shopping"
	"github.com/spf13/cobra"
)

var (
	boardname string
	user      string
	listName  string
)

func init() {
	// flags
	shoppingCmd.PersistentFlags().StringVar(&user, "user", "", "Specifcy a user in Trello")
	shoppingCmd.MarkPersistentFlagRequired("user")
	getShoppingListCmd.Flags().StringVar(&boardname, "name", "", "Specifcy a board in Trello")
	getShoppingListCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(shoppingCmd)
	shoppingCmd.AddCommand(getShoppingListCmd)
}

var shoppingCmd = &cobra.Command{
	Use:   "shopping",
	Short: "Interact with Shopping List stored on Trello",
	Long:  "Use this command to interact with Trello",
}

var getShoppingListCmd = &cobra.Command{
	Use:   "list",
	Short: "Gets a shopping list from Trello",
	Long:  "Assumes list is stored as a Checklist on a Trello Card",
	RunE: func(cmd *cobra.Command, args []string) error {
		shoplist, err := shopping.GetShoppingList(user, boardname)
		if err != nil {
			return fmt.Errorf("error getting board from Trello: %s", err.Error())
		}
		fmt.Printf("Retrieved board %s for user %s\n", shoplist, user)
		return nil
	},
}
