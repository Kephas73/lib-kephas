package health_check_module

import (
	"github.com/Kephas73/lib-kephas/env"
	"github.com/Kephas73/lib-kephas/modularization/health_check_module/controller"
	"github.com/labstack/echo/v4"
	"path"
)

var healthCheckCtrl *controller.HealthCheckCtrl

func Initialize(e *echo.Echo) {
	healthCheckCtrl = controller.NewHealthCheckCtrl()

	initRouter(e)
}

func initRouter(e *echo.Echo) {

	gr := e.Group(path.Join(env.Environment.SettingAPI.Path, env.Environment.SettingAPI.Version))

	gr.GET("/health-check/status", healthCheckCtrl.Status)
}
