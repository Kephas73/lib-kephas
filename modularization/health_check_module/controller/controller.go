package controller

import (
	"github.com/Kephas73/lib-kephas/api"
	"github.com/labstack/echo/v4"
)

type HealthCheckCtrl struct{}

func NewHealthCheckCtrl() *HealthCheckCtrl {
	return &HealthCheckCtrl{}
}

func (ctrl *HealthCheckCtrl) Status(ctx echo.Context) error {
	return api.WriteSuccessEmptyContent(ctx)
}
