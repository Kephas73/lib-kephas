package mail_client

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Address  string
	Host     string
	Port     int
	Username string
	Password string
}

var config *Config

// default value env key is "AWS";
// if configKeys was set, key env will be first value (not empty) of this
func getConfigFromEnv(configKeys ...string) *Config {
	configKey := "MailServer"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	config = &Config{}
	if err := viper.UnmarshalKey(configKey, config); err != nil {
		err := fmt.Errorf("not found config name with env %q for MailServer with error: %+v", configKey, err)
		panic(err)
	}

	if config.Host == "" {
		err := fmt.Errorf("not found any addr as Host for MailServer at %q", fmt.Sprintf("%s.Host", configKey))
		panic(err)
	}

	if config.Address == "" {
		err := fmt.Errorf("not found any addr as Address for MailServer at %q", fmt.Sprintf("%s.Address", configKey))
		panic(err)
	}

	if config.Port == 0 {
		err := fmt.Errorf("not found any addr as Port for MailServer at %q", fmt.Sprintf("%s.Port", configKey))
		panic(err)
	}

	if config.Username == "" {
		err := fmt.Errorf("not found any addr as Username for MailServer at %q", fmt.Sprintf("%s.Username", configKey))
		panic(err)
	}

	if config.Password == "" {
		err := fmt.Errorf("not found any addr as Password for MailServer at %q", fmt.Sprintf("%s.Password", configKey))
		panic(err)
	}
	return config
}
