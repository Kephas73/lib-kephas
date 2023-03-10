package controller

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/api"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/error_code"
	"github.com/Kephas73/lib-kephas/modularization/access_base_module/service"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"time"
)

type AccessBaseController struct {
	AccessBaseService service.IAccessBaseService
}

func NewAccessBaseController(cache *redis_client.RedisPool, sqlx *sqlx.DB, rbacModelPath string, timeout time.Duration) *AccessBaseController {
	return &AccessBaseController{
		AccessBaseService: service.NewAccessBaseService(cache, sqlx, rbacModelPath, timeout),
	}
}

func (controller *AccessBaseController) AccessGateway(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		nextAPI := false
		ctx := api.GetRequestContext(e)
		userUUID := api.GetContextDataString(e, constant.UserID)

		role, errC := controller.AccessBaseService.GetRoleByUser(ctx, userUUID)
		if errC == nil && role != nil && role.RoleID != constant.ValueEmpty {
			lP, errC := controller.AccessBaseService.GetsPermissionByRole(ctx, role.RoleID)
			if errC == nil && len(lP) != constant.ValueEmpty {

				for _, v := range lP {
					if access, errC := controller.AccessBaseService.Access(fmt.Sprintf("%d",v.PermissionID), e.Path(), e.Request().Method); access && errC == nil {
						nextAPI = true
						break
					}
				}

			}
		}

		if !nextAPI {

			errC = error_code.NewError(error_code.ERROR_ACCESS_DENIED, "denied", base.GetFunc())
			return api.WriteError(e, errC)
		}

		return next(e)
	}
}

func (controller *AccessBaseController) Status(e echo.Context) error {
	return api.WriteSuccessEmptyContent(e)
}
