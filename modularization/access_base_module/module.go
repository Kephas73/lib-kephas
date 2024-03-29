package access_base_module

import (
	"github.com/Kephas73/lib-kephas/env"
	"github.com/Kephas73/lib-kephas/modularization/access_base_module/controller"
	auth_base "github.com/Kephas73/lib-kephas/modularization/auth_base_module"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"path"
	"time"
)

var AccessBase *controller.AccessBaseController

// Initialize access_base_module:
//
// - Use logger lib-kephas
//
// - Create required tables (mysql):
//
// CREATE TABLE `rbac_policy` (
//	`permission_id` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
//	`method` VARCHAR(10) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
//	`path` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
//	`name` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
//	`created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`permission_id`, `path`, `method`) USING BTREE
//)
//COLLATE='utf8_unicode_ci'
//ENGINE=InnoDB;
//================================================================================================
//CREATE TABLE `rbac_role` (
//	`role_id` INT(11) NOT NULL,
//	`name` VARCHAR(255) NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
//	`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`role_id`) USING BTREE
//)
//COLLATE='utf8_unicode_ci'
//ENGINE=InnoDB;
//================================================================================================
//CREATE TABLE `rbac_permission` (
//	`permission_id` INT(11) NOT NULL,
//	`name` VARCHAR(255) NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
//	`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`permission_id`) USING BTREE
//)
//COLLATE='utf8_unicode_ci'
//ENGINE=InnoDB;
//================================================================================================
//CREATE TABLE `rbac_permission_role` (
//	`role_id` INT(11) NOT NULL,
//	`permission_id` INT(11) NOT NULL,
//	`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
//)
//COLLATE='utf8_unicode_ci'
//ENGINE=InnoDB;
//================================================================================================
func Initialize(e *echo.Echo, cache *redis_client.RedisPool, sqlx *sqlx.DB, rbacModelPath string, timeout time.Duration) {

	AccessBase = controller.NewAccessBaseController(cache, sqlx, rbacModelPath, timeout)

	initRouter(e)
}

func initRouter(e *echo.Echo) {
	gr := e.Group(path.Join(env.Environment.SettingAPI.Path, env.Environment.SettingAPI.Version))

	gr.GET("/access-base/status", AccessBase.Status)

	// get role and permission
	gr.GET("/access-base/user/role", AccessBase.GetUserRole, auth_base.AuthBase.JWTGateway)

	// chìa api get list user và role
	// chìa api thay đổi quyền

	gr.GET("/access-base/role", AccessBase.ListRole, auth_base.AuthBase.JWTGateway, AccessBase.AccessGateway)
	gr.POST("/access-base/role", AccessBase.CreateRole, auth_base.AuthBase.JWTGateway, AccessBase.AccessGateway)
	gr.PUT("/access-base/role/:id", AccessBase.UpdateRole, auth_base.AuthBase.JWTGateway, AccessBase.AccessGateway)
	gr.DELETE("/access-base/role/:id", AccessBase.DeleteRole, auth_base.AuthBase.JWTGateway, AccessBase.AccessGateway)

	gr.GET("/access-base/permission", AccessBase.ListPermission, auth_base.AuthBase.JWTGateway, AccessBase.AccessGateway)
}
