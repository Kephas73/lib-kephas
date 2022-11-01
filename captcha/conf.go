package captcha

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Decode struct {
		Website   string `json:"website"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		ClientKey string `json:"client_key"`
		Url       struct {
			Create string `json:"create"`
			Get    string `json:"get"`
		} `json:"url"`
	}
	PathKey string
}

var configCaptcha *Config

func getConfigFromEnv(configKeys ...string) {
	configKey := "Captcha"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	configCaptcha = &Config{}

	if err := viper.UnmarshalKey(configKey, configCaptcha); err != nil {
		panic(err)
	}
}
