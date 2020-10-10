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
func (c *Client) GetList(board *trello.Board, name string) (*trello.List, error) {
	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, err
	}
	for _, list := range lists {
		if list.Name == name {
			return list, nil
		}
	}
	return nil, fmt.Errorf("could not find list %s on board %s", name, board.Name)
}

// GetCard ...
func (c *Client) GetCard(list *trello.List, name string) (*trello.Card, error) {
	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}
	for _, card := range cards {
		if card.Name == name {
			return card, nil
		}
	}
	return nil, fmt.Errorf("could not find card %s on list %s", name, list.Name)
}
