package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

// A module for managing interactions with Trello

// GetBoard gets a trello board
// Params:
// + name string - the name of the trello board
// Return:
// + a struct representing a trello board
// + an error if the board is not found
func (c *Client) GetBoard(name string) (*trello.Board, error) {
	boards, err := c.User.GetBoards(trello.Defaults())
	if err != nil {
		return nil, err
	}
	for _, board := range boards {
		if board.Name == name {
			board, err := c.API.GetBoard(board.ID, trello.Defaults())
			if err != nil {
				return nil, err
			}
			return board, nil
		}
	}
	return nil, fmt.Errorf("could not find board: [%s]", name)
}

// GetList gets a list from a specific trello board
// Params:
// + boardname string - the name of the board where the list is located
// + name string - the name of the list
// Return:
// + A struct representing a trello list
// + an error if the list is not found
func (c *Client) GetList(boardname string, name string) (*trello.List, error) {
	board, err := c.GetBoard(boardname)
	if err != nil {
		return nil, err
	}
	lists, err := board.GetLists(trello.Defaults())
	for _, list := range lists {
		if list.Name == name {
			return list, nil
		}
	}
	return nil, fmt.Errorf("could not find list: [%s]", name)
}
