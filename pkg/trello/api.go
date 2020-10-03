package trello

import (
	"fmt"

	gotrello "github.com/VojtechVitek/go-trello"
)

// GetBoard ...
func (client *Client) GetBoard(name string) (*gotrello.Board, error) {
	boards, err := client.User.Boards()
	if err != nil {
		return nil, err
	}
	for _, board := range boards {
		if board.Name == name {
			return &board, nil
		}
	}
	return nil, fmt.Errorf("could not find board: [%s]", name)
}
