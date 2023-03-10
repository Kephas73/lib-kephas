package auth_base_module

import (
	"github.com/Kephas73/lib-kephas/modularization/auth_base_module/controller"
	"github.com/Kephas73/lib-kephas/redis_client"
	"time"
)

var AuthBase *controller.AuthBaseCtrl

func Initialize(cache *redis_client.RedisPool, timeout time.Duration) {
	AuthBase = controller.NewAuthBaseCtrl(cache, timeout)
}
