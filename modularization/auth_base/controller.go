package auth_base

import (
	"context"
	"fmt"
	"github.com/Kephas73/lib-kephas/JWToken"
	"github.com/Kephas73/lib-kephas/api"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/error_code"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/labstack/echo/v4"
	"time"
)

const (
	KeySessionUser          string = "auction:site:%d:session:%s"
	FieldTokenWeb           string = "token-web"
	FieldRefreshTokenWeb    string = "refresh-token-web"
	FieldTokenMobile        string = "token-mobile"
	FieldRefreshTokenMobile string = "refresh-token-mobile"
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
		accessUUID, accountUUID, deviceKind, deviceIP, clientSite, _, errCode := JWToken.ExtractToken(ctx, shortLive)
		if errCode != nil {
			return api.WriteError(ctx, errCode)
		}

		// Check token redis
		if !ctrl.CheckAccessToken(accessUUID, accountUUID, deviceKind, clientSite) {
			errCode = error_code.NewError(error_code.ERROR_TOKEN_INVALID, "token does not match current session", base.GetFunc())
			return api.WriteError(ctx, errCode)
		}

		// Set context
		api.SetContextDataString(ctx, constant.AccessUUID, accessUUID)
		api.SetContextDataString(ctx, constant.DeviceKind, deviceKind)
		api.SetContextDataString(ctx, constant.DeviceIP, deviceIP)
		api.SetContextDataString(ctx, constant.UserID, accountUUID)
		api.SetContextDataString(ctx, constant.ClientSite, clientSite)

		return next(ctx)
	}
}

func (ctrl *AuthBaseCtrl) JWTRefresh(next echo.HandlerFunc) echo.HandlerFunc {
	return ctrl.Refresh(next, false)
}

func (ctrl *AuthBaseCtrl) Refresh(next echo.HandlerFunc, shortLive bool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		refreshUUID, accountUUID, deviceKind, clientSite, errCode := JWToken.ExtractRefreshToken(ctx)
		if errCode != nil {
			return api.WriteError(ctx, errCode)
		}

		// Check token redis
		if !ctrl.CheckRefreshToken(refreshUUID, accountUUID, deviceKind, clientSite) {
			errCode = error_code.NewError(error_code.ERROR_REFRESH_TOKEN_INVALID, "token refresh does not match current session", base.GetFunc())
			return api.WriteError(ctx, errCode)
		}

		// Set context
		api.SetContextDataString(ctx, constant.DeviceKind, deviceKind)
		api.SetContextDataString(ctx, constant.UserID, accountUUID)
		api.SetContextDataString(ctx, constant.ClientSite, clientSite)

		return next(ctx)
	}
}

func (ctrl *AuthBaseCtrl) CheckAccessToken(accessUUID, userID string, deviceKind, clientSite int) bool {

	ctx, cancel := context.WithTimeout(context.Background(), ctrl.Timeout)
	defer cancel()

	conn := ctrl.CacheRepository.Get()
	if conn == nil {
		return false
	}

	mResult := conn.HGetAll(ctx, fmt.Sprintf(KeySessionUser, clientSite, userID)).Val()
	if deviceKind == constant.KDeviceKindWeb {
		if access, ok := mResult[FieldTokenWeb]; ok {
			return access == accessUUID
		}
	} else if deviceKind == constant.KDeviceKindMobile {
		if access, ok := mResult[FieldTokenMobile]; ok {
			return access == accessUUID
		}
	}

	return false
}

func (ctrl *AuthBaseCtrl) CheckRefreshToken(refreshUUID, userID string, deviceKind, clientSite int) bool {

	ctx, cancel := context.WithTimeout(context.Background(), ctrl.Timeout)
	defer cancel()

	conn := ctrl.CacheRepository.Get()
	if conn == nil {
		return false
	}

	mResult := conn.HGetAll(ctx, fmt.Sprintf(KeySessionUser, clientSite, userID)).Val()
	if deviceKind == constant.KDeviceKindWeb {
		if refresh, ok := mResult[FieldRefreshTokenWeb]; ok {
			return refresh == refreshUUID
		}
	} else if deviceKind == constant.KDeviceKindMobile {
		if refresh, ok := mResult[FieldRefreshTokenMobile]; ok {
			return refresh == refreshUUID
		}
	}

	return false
}
