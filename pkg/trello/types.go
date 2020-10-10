package trello

import (
	"github.com/adlio/trello"
	"github.com/sirupsen/logrus"
)

// Client wrapper for Trello api
type Client struct {
	API  *trello.Client
	User *trello.Member
}

// NewClient ...
func NewClient(appKey string, token string, member string) (Client, error) {
	// Instantiate logger to allow debug level logging
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	// Instantiate client and attach logger
	client := trello.NewClient(appKey, token)
	client.Logger = logger

	// set member
	user, err := client.GetMember(member, trello.Defaults())
	if err != nil {
		return Client{}, err
	}
	return Client{
		API:  client,
		User: user,
	}, nil
}
