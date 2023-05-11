package controller

import (
	"context"
	"fmt"
	"github.com/Kephas73/lib-kephas/api"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/error_code"
	"github.com/Kephas73/lib-kephas/modularization/access_base_module/service"
	"github.com/Kephas73/lib-kephas/modularization/model"
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
		if errC == nil && role != nil && role.RoleID != constant.StrEmpty {
			if len(role.RoleInfo.PermissionIDs) != constant.ValueEmpty {

				for _, v := range role.RoleInfo.PermissionIDs {
					if access, errC := controller.AccessBaseService.Access(fmt.Sprintf("%d", v), e.Path(), e.Request().Method); access && errC == nil {
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

func (controller *AccessBaseController) GetRoleByUser(ctx context.Context, userUUID string) (*model.UserRoleRes, *error_code.ErrorCode) {
	return controller.AccessBaseService.GetRoleByUser(ctx, userUUID)
}

func (controller *AccessBaseController) GetMapRoleUser(ctx context.Context, userUUID ...string) map[string]*model.UserRoleRes {
	return controller.AccessBaseService.GetMapRoleUser(ctx, userUUID...)
}

func (controller *AccessBaseController) ChangeRoleUser(ctx context.Context, userUUID, roleID string) *error_code.ErrorCode {
	return controller.AccessBaseService.ChangeRoleUser(ctx, &model.UserRole{UserUUID: userUUID, RoleID: roleID})
}

func (controller *AccessBaseController) DeleteRoleUser(ctx context.Context, userUUID string) *error_code.ErrorCode {
	return controller.AccessBaseService.DeleteRoleUser(ctx, userUUID)
}

func (controller *AccessBaseController) Status(e echo.Context) error {
	return api.WriteSuccessEmptyContent(e)
}

func (controller *AccessBaseController) ListRole(e echo.Context) error {

	ctx := api.GetRequestContext(e)
	name := e.QueryParam("name")
	userUUID := api.GetContextDataString(e, constant.UserID)

	if userUUID == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_NULL_ID, "not found user uuid", base.GetFunc())
		return api.WriteError(e, errC)
	}

	role, errC := controller.AccessBaseService.GetsRole(ctx, name)
	if errC != nil {
		return api.WriteError(e, errC)
	}

	return api.WriteSuccess(e, role)
}

func (controller *AccessBaseController) CreateRole(e echo.Context) error {

	var roleReq model.RoleReq
	ctx := api.GetRequestContext(e)
	userUUID := api.GetContextDataString(e, constant.UserID)

	if userUUID == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_NULL_ID, "not found user uuid", base.GetFunc())
		return api.WriteError(e, errC)
	}

	if err := e.Bind(&roleReq); err != nil {
		errC := error_code.NewError(error_code.ERROR_BIND_DATA, err.Error(), base.GetFunc())
		return api.WriteError(e, errC)
	}

	if errC := roleReq.Validate(); errC != nil {
		return api.WriteError(e, errC)
	}

	role, errC := controller.AccessBaseService.CreateRole(ctx, &model.Role{Name: roleReq.Name, PermissionIDs: roleReq.PermissionIDs, TeamName: roleReq.TeamName, Description: roleReq.Description})
	if errC != nil {
		return api.WriteError(e, errC)
	}

	return api.WriteSuccess(e, role)
}

func (controller *AccessBaseController) UpdateRole(e echo.Context) error {

	var roleReq model.RoleReq
	ctx := api.GetRequestContext(e)
	userUUID := api.GetContextDataString(e, constant.UserID)
	id := e.Param("id")

	if userUUID == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_NULL_ID, "not found user uuid", base.GetFunc())
		return api.WriteError(e, errC)
	}

	if id == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_DATA_INVALID, "not found role id", base.GetFunc())
		return api.WriteError(e, errC)
	}

	if err := e.Bind(&roleReq); err != nil {
		errC := error_code.NewError(error_code.ERROR_BIND_DATA, err.Error(), base.GetFunc())
		return api.WriteError(e, errC)
	}

	if errC := roleReq.Validate(); errC != nil {
		return api.WriteError(e, errC)
	}

	role, errC := controller.AccessBaseService.UpdateRole(ctx, &model.Role{Name: roleReq.Name, RoleID: id, PermissionIDs: roleReq.PermissionIDs, TeamName: roleReq.TeamName, Description: roleReq.Description})
	if errC != nil {
		return api.WriteError(e, errC)
	}

	return api.WriteSuccess(e, role)
}

func (controller *AccessBaseController) DeleteRole(e echo.Context) error {

	ctx := api.GetRequestContext(e)
	userUUID := api.GetContextDataString(e, constant.UserID)
	id := e.Param("id")

	if userUUID == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_NULL_ID, "not found user uuid", base.GetFunc())
		return api.WriteError(e, errC)
	}

	if id == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_DATA_INVALID, "not found role id", base.GetFunc())
		return api.WriteError(e, errC)
	}

	errC := controller.AccessBaseService.DeleteRole(ctx, &model.Role{RoleID: id})
	if errC != nil {
		return api.WriteError(e, errC)
	}

	return api.WriteSuccessEmptyContent(e)
}

func (controller *AccessBaseController) ListPermission(e echo.Context) error {

	ctx := api.GetRequestContext(e)
	name := e.QueryParam("name")
	userUUID := api.GetContextDataString(e, constant.UserID)

	if userUUID == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_NULL_ID, "not found user uuid", base.GetFunc())
		return api.WriteError(e, errC)
	}

	permission, errC := controller.AccessBaseService.GetsPermission(ctx, name)
	if errC != nil {
		return api.WriteError(e, errC)
	}

	return api.WriteSuccess(e, permission)
}

func (controller *AccessBaseController) GetUserRole(e echo.Context) error {

	ctx := api.GetRequestContext(e)
	userUUID := api.GetContextDataString(e, constant.UserID)

	if userUUID == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_NULL_ID, "not found user uuid", base.GetFunc())
		return api.WriteError(e, errC)
	}

	role, errC := controller.AccessBaseService.GetRoleByUser(ctx, userUUID)
	if errC != nil {
		return api.WriteError(e, errC)
	}

	return api.WriteSuccess(e, role)
}
