package trello

import (
	"fmt"
	"os"

	"github.com/adlio/trello"
	"github.com/sirupsen/logrus"
)

const (
	trelloAppEnvKey   = "TRELLO_APPKEY"
	trelloTokenEnvKey = "TRELLO_TOKEN"
)

// Client wrapper for Trello api
type Client struct {
	API  *trello.Client
	User *trello.Member
}

// NewClient ...
func NewClient(member string) (*Client, error) {
	// validate requirements for instantiating a connection with the Trello Api
	appKey := os.Getenv(trelloAppEnvKey)
	token := os.Getenv(trelloTokenEnvKey)
	if appKey == "" {
		return nil, fmt.Errorf("required enviornment variable empty: %s", trelloAppEnvKey)
	}
	if token == "" {
		return nil, fmt.Errorf("required enviornment variable empty: %s", trelloTokenEnvKey)
	}

	// Instantiate logger to allow debug level logging
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	// Instantiate client and attach logger
	client := trello.NewClient(appKey, token)
	client.Logger = logger

	// set member
	user, err := client.GetMember(member, trello.Defaults())
	if err != nil {
		return nil, err
	}
	return &Client{
		API:  client,
		User: user,
	}, nil
}
