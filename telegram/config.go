package telegram

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Bot struct {
		Token     string `json:"token"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		ParseMode string `json:"parse_mode"`
		Endpoint  string `json:"endpoint"`
		Timeout   int    `json:"timeout"`
	}
}

var teleConfig *Config

func getTeleConfigFromEnv(configKeys ...string) {
	configKey := "Telegram"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	teleConfig = &Config{}
	if err := viper.UnmarshalKey(configKey, &teleConfig); err != nil {
		err = fmt.Errorf("not found config name with env %q for telegram app authorization with error: %+v", configKey, err)
		panic(err)
	}

	if teleConfig.Bot.Token == "" {
		err := fmt.Errorf("not found token property of config for %q", configKey)
		panic(err)
	}

	if teleConfig.Bot.Username == "" {
		err := fmt.Errorf("not found username property of config for %q", configKey)
		panic(err)
	}

	if teleConfig.Bot.Name == "" {
		err := fmt.Errorf("not found name property of config for %q", configKey)
		panic(err)
	}

	if teleConfig.Bot.ParseMode == "" {
		err := fmt.Errorf("not found parsemode property of config for %q", configKey)
		panic(err)
	}

	if teleConfig.Bot.Endpoint == "" {
		err := fmt.Errorf("not found endpoint property of config for %q", configKey)
		panic(err)
	}

	if teleConfig.Bot.Timeout == 0 {
		teleConfig.Bot.Timeout = 10
	}
}
