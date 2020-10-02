package trello

import (
	gotrello "github.com/VojtechVitek/go-trello"
)

// NewClient ...
func NewClient(apiKey string, token string) (*gotrello.Client, error) {
	trello, err := gotrello.NewAuthClient(apiKey, &token)
	if err != nil {
		return nil, err
	}
	return trello, nil
}
