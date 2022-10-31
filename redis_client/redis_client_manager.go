package redis_client

import (
	"fmt"
)

type redisClientManager struct {
	redisClients map[string]*RedisPool
}

var redisClients = &redisClientManager{make(map[string]*RedisPool)}

// InstallRedisClientManager func;
// default value env key is "Redis";
// if configKeys was set, key env will be first value (not empty) of this;
// For the config, if addr is not empty, create redis node connection,
// if addrs and master name is not empty, create redis sentinel connection,
// otherwise create redis cluster connection
func InstallRedisClientManager(configKeys ...string) {
	configs := getRedisConfigFromEnv(configKeys...)
	if len(configs) == 0 {
		err := fmt.Errorf("not found config for redis cluster manager")

		panic(err)
	}

	for _, config := range configs {
		client := NewRedisUniversalClient(config)

		redisClients.redisClients[config.Name] = client
	}
}

// InstallRedisClientManagerWithConfig func;
// For the config, if addr is not empty, create redis node connection,
// if addrs and master name is not empty, create redis sentinel connection,
// otherwise create redis cluster connection
func InstallRedisClientManagerWithConfig(configs []*RedisConfig) {
	if len(configs) == 0 {
		err := fmt.Errorf("not found config for redis cluster manager")
		panic(err)
	}

	for _, config := range configs {
		client := NewRedisUniversalClient(config)

		redisClients.redisClients[config.Name] = client
	}
}

// GetRedisClient func
func GetRedisClient(redisName string) (client *RedisPool) {
	client, ok := redisClients.redisClients[redisName]
	if !ok {
		listName := make([]string, 0)
		for name := range redisClients.redisClients {
			listName = append(listName, name)
		}
	}
	return
}

// GetRedisClientManager func
func GetRedisClientManager() map[string]*RedisPool {
	return redisClients.redisClients
}
