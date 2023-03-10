package service

import (
	"github.com/Kephas73/lib-kephas/modularization/access_base_module/repository"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

type IAccessBaseService interface {
}

type AccessBaseService struct {
	Timout time.Duration
	Cache  *redis_client.RedisPool

	AccessBaseRepository repository.IAccessBaseRepository
}

var (
	accessBaseServiceInstance *AccessBaseService
	muxAccessBaseSvInstance   sync.Mutex
)

func NewAccessBaseService(cache *redis_client.RedisPool, sqlx *sqlx.DB, timeout time.Duration) IAccessBaseService {

	if accessBaseServiceInstance == nil {
		authService := AccessBaseService{
			Timout: timeout,
			Cache:  cache,

			AccessBaseRepository: repository.NewAccessBaseRepository(sqlx),
		}

		muxAccessBaseSvInstance.Lock()
		accessBaseServiceInstance = &authService
		muxAccessBaseSvInstance.Unlock()
	}

	return accessBaseServiceInstance
}
