package trello

import (
	"fmt"

	gotrello "github.com/VojtechVitek/go-trello"
)

// GetBoard get a trello board
func (c *Client) GetBoard(name string) (*gotrello.Board, error) {
	boards, err := c.User.Boards()
	fmt.Println(boards)
	if err != nil {
		return nil, err
	}
	for _, board := range boards {
		fmt.Printf("Board name: %s\n", board.Name)
		if board.Name == name {
			return &board, nil
		}
	}
	return nil, fmt.Errorf("could not find board: [%s]", name)
}

// GetList gets a list from a specific trello board
func (c *Client) GetList(boardname string, name string) (*gotrello.List, error) {
	board, err := c.GetBoard(boardname)
	if err != nil {
		return nil, err
	}
	lists, err := board.Lists()
	for _, list := range lists {
		if list.Name == name {
			return &list, nil
		}
	}
	return nil, fmt.Errorf("could not find list: [%s]", name)
}
