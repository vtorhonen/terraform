package upcloud

import (
	"fmt"
	"log"

	"github.com/Jalle19/upcloud-go-sdk/upcloud"
	"github.com/Jalle19/upcloud-go-sdk/upcloud/client"
	"github.com/Jalle19/upcloud-go-sdk/upcloud/service"
)

type Config struct {
	Username string
	Password string
}

func (c *Config) Client(*service.Service, error) {
	client, err := client.New(c.Username, c.Password)
	if err != nil {
		return nil, fmt.Errorf("Error creating UpCloud client: %s", err)
	}
	svc, err := service.New(client)
	if err != nil {
		if serviceError, ok := err.(*upcloud.Error); ok {
			errMsg := fmt.Errorf("Error creating Service object. Error code: %d, Error message: %s",
				serviceError.ErrorCode, serviceError.ErrorMessage)
			return nil, errMsg
		}
	}

	log.Printf("[INFO] UpCloud Client configured for user: %s", c.Username)
	return svc, nil
}
