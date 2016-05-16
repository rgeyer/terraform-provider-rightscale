package main

import (
	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/rsapi"
)

// Config struct holds params for rightscale client
type Config struct {
	AccountID    int
	RefreshToken string
	APIHost      string
}

// Client connection to RS API(s)
func (c *Config) Client() (*cm15.API, error) {
	auth := rsapi.NewOAuthAuthenticator(c.RefreshToken, c.AccountID)
	client := cm15.New(c.APIHost, auth)

	return client, nil
}
