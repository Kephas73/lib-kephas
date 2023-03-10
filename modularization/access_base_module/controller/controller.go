package controller

import (
	"github.com/Kephas73/lib-kephas/api"
	"github.com/Kephas73/lib-kephas/modularization/access_base_module/service"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"time"
)

type AccessBaseController struct {
	AccessBase service.IAccessBaseService
}

func NewAuthController(cache *redis_client.RedisPool, sqlx *sqlx.DB, timeout time.Duration) *AccessBaseController {
	return &AccessBaseController{
		AccessBase: service.NewAccessBaseService(cache, sqlx, timeout),
	}
}

func (controller *AccessBaseController) Status(ctx echo.Context) error {
	return api.WriteSuccessEmptyContent(ctx)
}
