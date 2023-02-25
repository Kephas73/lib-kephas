package mqtt_client

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/env"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	KDefaultContentType string = "application/json"
)

// MessageQueue func;
type MessageQueue struct {
	Topic  string      `json:"Topic,omitempty"` // Must have when process message on consumer
	Action string      `json:"Action"`
	Data   interface{} `json:"Data"`
}

type Config struct {
	Host     string `json:"Host,omitempty"`
	Port     int32  `json:"Port,omitempty"`
	Username string `json:"Username,omitempty"`
	Password string `json:"Password,omitempty"`
	ClientID string `json:"ClientID,omitempty"`
}

var (
	config *Config
)

// default value env key is "MQTT";
// if configKeys was set, key env will be first value (not empty) of this;
func getMQTTConfigFromEnv(configKeys ...string) {
	configKey := "Mqtt"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	config = &Config{}

	if err := viper.UnmarshalKey(configKey, config); err != nil {
		err := fmt.Errorf("not found config with env %q for mqtt with error: %+v", configKey, err)
		panic(err)
	}

	if config.Host == "" {
		err := fmt.Errorf("not found any addr as host for mq at %q", fmt.Sprintf("%s.Host", configKey))
		panic(err)
	}

	if config.Port == 0 {
		err := fmt.Errorf("not found any addr as port for mq at %q", fmt.Sprintf("%s.Port", configKey))
		panic(err)
	}

	if config.ClientID == "" {
		serviceIdentify := fmt.Sprintf("%s_%d", env.Environment.Server.ServiceName, time.Now().Unix())
		if serviceIdentify == "" {
			serviceIdentify = "gtv-backend"
		}

		config.ClientID = serviceIdentify
	}
}
