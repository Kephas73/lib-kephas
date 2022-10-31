package redis_client

import (
	"context"
	"fmt"
	"github.com/Kephas73/lib-kephas/env"
	"testing"
)

func init() {
	env.SetupConfigEnv("../config.json")
}

func TestGetRedisClient(t *testing.T) {
	InstallRedisClientManager()
	fmt.Println(GetRedisClient("broker_cache").Get().Ping(context.Background()).String())
}
