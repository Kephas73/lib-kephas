package logstash

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var TimeoutDefault = 10

type LogStashConfig struct {
	Hosts   []string
	Timeout int
}

var logStashConf *LogStashConfig

func createConfigFromEnv(configKeys ...string) {
	configKey := "LogStash"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	logStashConf = &LogStashConfig{}

	if err := viper.UnmarshalKey(configKey, logStashConf); err != nil {
		err := fmt.Errorf("not found config name with env %q for logstash with error: %+v", configKey, err)
		panic(err)
	}

	if len(logStashConf.Hosts) == 0 {
		err := fmt.Errorf("not found hosts for logstash with env %q", fmt.Sprintf("%s.Hosts", configKey))
		panic(err)
	}

	if logStashConf.Timeout <= 0 {
		logStashConf.Timeout = TimeoutDefault
	}
}
