package trello

import (
	gotrello "github.com/VojtechVitek/go-trello"
)

// Client wrapper for Trello api
type Client struct {
	API  *gotrello.Client
	User *gotrello.Member
}

// NewClient ...
func NewClient(apiKey string, token string, member string) (Client, error) {
	trello, err := gotrello.NewAuthClient(apiKey, &token)
	if err != nil {
		return Client{}, err
	}
	// set member
	user, err := trello.Member(member)
	if err != nil {
		return Client{}, err
	}
	return Client{
		API:  trello,
		User: user,
	}, nil
}
