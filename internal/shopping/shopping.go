package shopping

import (
	"feargusogorman.com/leto/v1/pkg/trello"
)

// GetShoppingList ...
func GetShoppingList(member string, boardname string) (string, error) {
	client, err := trello.NewClient(member)
	if err != nil {
		return "", err
	}
	board, err := client.GetBoard(boardname)
	if err != nil {
		return "", err
	}
	return board.Name, nil
}
