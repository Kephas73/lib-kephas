package health_check

import (
	"github.com/Kephas73/lib-kephas/api"
	"github.com/Kephas73/lib-kephas/env"
	"github.com/labstack/echo/v4"
	"path"
)

type HealthCheckCtrl struct{}

var healthCheckCtrl *HealthCheckCtrl

func NewHealthCheckCtrl() *HealthCheckCtrl {
	return &HealthCheckCtrl{}
}

func Initialize(e *echo.Echo) {
	healthCheckCtrl = NewHealthCheckCtrl()

	initRouter(e)
}

func initRouter(e *echo.Echo) {

	gr := e.Group(path.Join(env.Environment.SettingAPI.Path, env.Environment.SettingAPI.Version))

	gr.GET("/health-check/status", healthCheckCtrl.Status)
}

func (ctrl *HealthCheckCtrl) Status(ctx echo.Context) error {
	return api.WriteSuccessEmptyContent(ctx)
}
