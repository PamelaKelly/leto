package shopping

import (
	"fmt"

	"feargusogorman.com/leto/v1/pkg/trello"
)

// GetShoppingList ...
func GetShoppingList(member string, boardname string, listname string, cardname string) (string, error) {
	client, err := trello.NewClient(member)
	if err != nil {
		return "", err
	}
	board, err := client.GetBoard(boardname)
	if err != nil {
		return "", err
	}
	fmt.Println(board.Name)
	list, err := client.GetList(board, listname)
	if err != nil {
		return "", err
	}
	fmt.Println(list.Name)
	card, err := client.GetCard(list, cardname)
	if err != nil {
		return "", err
	}

	for _, id := range card.IDCheckLists {
		fmt.Println(id)
		// Get the checklist corresponding to that id
	}
	return card.Name, nil
}
