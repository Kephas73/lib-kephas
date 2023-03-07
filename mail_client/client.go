package mail_client

import (
	"fmt"
)

type MailClient struct {
	config *Config
}

var mailClientInstance *MailClient

func InstanceMailClientManager(configKeys ...string) *MailClient {
	if mailClientInstance != nil {
		return mailClientInstance
	}

	if config == nil {
		getConfigFromEnv(configKeys...)
	}

	if config == nil {
		err := fmt.Errorf("need config for mail client first")
		panic(err)
	}

	mailClientInstance = &MailClient{
		config: config,
	}

	return mailClientInstance
}

func GetMailClientInstance() *MailClient {
	if mailClientInstance == nil {
		return InstanceMailClientManager()
	}

	return mailClientInstance
}
