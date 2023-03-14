package redis_client

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	DefaultTimeout = 30 * time.Second
)

type ConfigSet struct {
	Key   string
	Value string
}

// RedisConfig - Redis client config.
type RedisConfig struct {
	Name         string   `json:"name,omitempty"` // redis name
	Environment  string   `json:"environment,omitempty"`
	Hosts        []string `json:"host,omitempty"`
	Host         string   `json:"addr,omitempty"`
	Active       int      `json:"active,omitempty"`        // pool
	Idle         int      `json:"idle,omitempty"`          // pool
	DialTimeout  uint     `json:"dial_timeout,omitempty"`  // In seconds
	ReadTimeout  uint     `json:"read_timeout,omitempty"`  // In seconds
	WriteTimeout uint     `json:"write_timeout,omitempty"` // In seconds
	IdleTimeout  uint     `json:"idle_timeout,omitempty"`  // In seconds

	DBNum    int    `json:"db_num,omitempty"`   // db num
	Username string `json:"username,omitempty"` //
	Password string `json:"password,omitempty"` // password

	MaxRetries int    `json:"max_retries,omitempty"` //
	MasterName string `json:"master_name,omitempty"` //

	ExpiredEvents bool `json:"expired_events,omitempty"`
}

// ToRedisCacheConfig func
func (c *RedisConfig) ToRedisCacheConfig() string {
	// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
	// rc.key = cf["key"]
	// rc.conninfo = cf["conn"]
	// rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	// rc.password = cf["password"]
	return fmt.Sprintf(`{"conn":"%s", "dbNum":"%d", "password":"%s"}`, c.Hosts, c.DBNum, c.Password)
}

// default value env key is "Redis";
// if configKeys was set, key env will be first value (not empty) of this
func getRedisConfigFromEnv(configKeys ...string) []*RedisConfig {
	configKey := "Redis"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	listConf := make([]*RedisConfig, 0)
	if err := viper.UnmarshalKey(configKey, &listConf); err != nil {
		err = fmt.Errorf("not found config name with env %q for Redis with error: %+v", configKey, err)
		panic(err)
	}

	return listConf
}
