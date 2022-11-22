package auth_base

import (
	"github.com/Kephas73/lib-kephas/JWToken"
	"github.com/Kephas73/lib-kephas/api"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/labstack/echo/v4"
	"time"
)

// AuthBaseCtrl func:
// prerequisite: init jwt
type AuthBaseCtrl struct {
	Timeout         time.Duration
	CacheRepository *redis_client.RedisPool
}

var AuthBase *AuthBaseCtrl

func NewAuthBaseCtrl(cache *redis_client.RedisPool, timeout time.Duration) *AuthBaseCtrl {
	return &AuthBaseCtrl{
		Timeout:         timeout,
		CacheRepository: cache,
	}
}

func Initialize(cache *redis_client.RedisPool, timeout time.Duration) {
	AuthBase = NewAuthBaseCtrl(cache, timeout)
}

func (ctrl *AuthBaseCtrl) JWTGateway(next echo.HandlerFunc) echo.HandlerFunc {
	return ctrl.BaseGateway(next, false)
}

func (ctrl *AuthBaseCtrl) SecureGateway(next echo.HandlerFunc) echo.HandlerFunc {
	return ctrl.BaseGateway(next, true)
}

func (ctrl *AuthBaseCtrl) BaseGateway(next echo.HandlerFunc, shortLive bool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		accessUUID, accountUUID, deviceKind, deviceIP, _, errCode := JWToken.ExtractToken(ctx, shortLive)
		if errCode != nil {
			return api.WriteError(ctx, errCode)
		}

		// Check token redis

		// Set context
		api.SetContextDataString(ctx, constant.AccessUUID, accessUUID)
		api.SetContextDataString(ctx, constant.DeviceKind, deviceKind)
		api.SetContextDataString(ctx, constant.DeviceIP, deviceIP)
		api.SetContextDataString(ctx, constant.UserID, accountUUID)

		return next(ctx)
	}
}

func (ctrl *AuthBaseCtrl) JWTRefresh(next echo.HandlerFunc) echo.HandlerFunc {
	return ctrl.Refresh(next, false)
}

func (ctrl *AuthBaseCtrl) Refresh(next echo.HandlerFunc, shortLive bool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		accessUUID, accountUUID, errCode := JWToken.ExtractRefreshToken(ctx)
		if errCode != nil {
			return api.WriteError(ctx, errCode)
		}

		// Check token redis

		// Set context
		api.SetContextDataString(ctx, constant.AccessUUID, accessUUID)
		api.SetContextDataString(ctx, constant.UserID, accountUUID)

		return next(ctx)
	}
}
