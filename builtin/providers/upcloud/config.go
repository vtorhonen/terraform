package upcloud

import (
	"github.com/Jalle19/upcloud-go-sdk/upcloud/client"
	"github.com/Jalle19/upcloud-go-sdk/upcloud/service"
)

type Config struct {
	Username string
	Password string
}

func (c *Config) Service(*service.Service, error) {
	client, err := client.New(c.Username, c.Password)
	if err != nil {
		return nil, err
	}
	svc := service.New(client)
	return svc, nil
}
