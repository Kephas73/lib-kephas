package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type LoggerConfig struct {
	MaxSize     int    `json:"max_size,omitempty"`    // MB
	MaxBackups  int    `json:"max_backups,omitempty"` // File backup tối đa được giữ
	MaxAge      int    `json:"max_age,omitempty"`     // Thời gian giữ file backup (day)
	Compress    bool   `json:"compress,omitempty"`    // Nén log cũ
	Path        string `json:"path,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
	SentryDns   string `json:"sentry_dns,omitempty"`
	Description string `json:"description,omitempty"`
}

var loggerConf *LoggerConfig

func createConfigFromEnv(configKeys ...string) {
	configKey := "Logger"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	loggerConf = &LoggerConfig{}

	if err := viper.UnmarshalKey(configKey, loggerConf); err != nil {
		err := fmt.Errorf("not found config name with env %q for logstash with error: %+v", configKey, err)
		panic(err)
	}

	if loggerConf.MaxSize == 0 {
		loggerConf.MaxSize = 10 // MB
	}

	if loggerConf.MaxBackups == 0 {
		loggerConf.MaxBackups = 10
	}

	if loggerConf.MaxAge == 0 {
		loggerConf.MaxAge = 5
	}
}
