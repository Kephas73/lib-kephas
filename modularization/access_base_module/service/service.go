package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/error_code"
	"github.com/Kephas73/lib-kephas/logger"
	"github.com/Kephas73/lib-kephas/modularization/access_base_module/repository"
	"github.com/Kephas73/lib-kephas/modularization/model"
	"github.com/Kephas73/lib-kephas/redis_client"
	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

type IAccessBaseService interface {
	Access(permissionID string, pathAPI string, method string) (bool, *error_code.ErrorCode)
	GetsPolicy(ctx context.Context) ([]*model.Policy, *error_code.ErrorCode)
	GetsPermissionByRole(ctx context.Context, roleID int) ([]*model.PermissionRole, *error_code.ErrorCode)
	GetRoleByUser(ctx context.Context, userUUID string) (*model.UserRole, *error_code.ErrorCode)
}

type AccessBaseService struct {
	Timout        time.Duration
	Cache         *redis_client.RedisPool
	RbacModelPath string
	Enforcer      *casbin.Enforcer
	muAccess      sync.Mutex

	AccessBaseRepository repository.IAccessBaseRepository
}

var (
	accessBaseServiceInstance *AccessBaseService
	muxAccessBaseSvInstance   sync.Mutex
)

func NewAccessBaseService(cache *redis_client.RedisPool, sqlx *sqlx.DB, rbacModelPath string, timeout time.Duration) IAccessBaseService {

	if accessBaseServiceInstance == nil {
		authService := AccessBaseService{
			Timout:        timeout,
			Cache:         cache,
			RbacModelPath: rbacModelPath,

			AccessBaseRepository: repository.NewAccessBaseRepository(sqlx),
		}

		muxAccessBaseSvInstance.Lock()
		accessBaseServiceInstance = &authService
		muxAccessBaseSvInstance.Unlock()

		// init policy
		accessBaseServiceInstance.InitPolicy()
	}

	return accessBaseServiceInstance
}

func (service *AccessBaseService) Access(permissionID string, pathAPI string, method string) (bool, *error_code.ErrorCode) {
	access, err := service.Enforcer.Enforce(permissionID, pathAPI, method)
	if err != nil {
		errC := error_code.NewError(error_code.ERROR_ACCESS_DENIED, "access denied", base.GetFunc())
		return false, errC
	}

	return access, nil
}

func (service *AccessBaseService) GetsPolicy(ctx context.Context) ([]*model.Policy, *error_code.ErrorCode) {
	ctx, cancel := context.WithTimeout(ctx, service.Timout)
	defer cancel()

	policy, err := service.AccessBaseRepository.SelectsPolicy(ctx)
	if err != nil {
		logger.Error("AccessBaseService::GetsPolicy: -Error: %v", err)
		errC := error_code.NewError(error_code.ERROR_RETRIEVE_DATA, err.Error(), base.GetFunc())
		return nil, errC
	}

	return policy, nil
}

func (service *AccessBaseService) GetsPermissionByRole(ctx context.Context, roleID int) ([]*model.PermissionRole, *error_code.ErrorCode) {
	permission, err := service.CacheGetPermissionRole(roleID)
	if err != nil || permission == nil {
		permission, err = service.AccessBaseRepository.SelectPermissionByRole(ctx, roleID)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("AccessBaseService::GetsPermissionByRole: -Get permission by roleID: %d, error: %v", roleID, err)
			errC := error_code.NewError(error_code.ERROR_RETRIEVE_DATA, err.Error(), base.GetFunc())
			return nil, errC
		}

		if permission == nil {
			errC := error_code.NewError(error_code.ERROR_ACCESS_DENIED, "access denied", base.GetFunc())
			return nil, errC
		}

		service.CacheSetPermissionRole(permission, roleID)
	}

	return permission, nil
}

func (service *AccessBaseService) GetRoleByUser(ctx context.Context, userUUID string) (*model.UserRole, *error_code.ErrorCode) {
	userRole, err := service.CacheGetRoleUser(userUUID)
	if err != nil || userRole == nil {
		userRole, err = service.AccessBaseRepository.SelectUserRoleByUserUUID(ctx, userUUID)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("AccessBaseService::GetRoleByUser: -Get role by userUUID: %s, error: %v", userUUID, err)
			errC := error_code.NewError(error_code.ERROR_RETRIEVE_DATA, err.Error(), base.GetFunc())
			return nil, errC
		}

		if userRole == nil {
			errC := error_code.NewError(error_code.ERROR_ACCESS_DENIED, "access denied", base.GetFunc())
			return nil, errC
		}

		service.CacheSetRoleUser(userRole)
	}

	return userRole, nil
}

func (service *AccessBaseService) LoadPolicy() *error_code.ErrorCode {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	policy, errC := service.GetsPolicy(ctx)
	if errC != nil {
		logger.Error("AccessBaseService::LoadPolicy: -Error: %v", errC.ErrorMessage)
		return errC
	}

	e, err := casbin.NewEnforcer(service.RbacModelPath)
	if err != nil {
		logger.Error("AccessBaseService::LoadPolicy: -New enforcer error: %v", err)
		errC = error_code.NewError(error_code.ERROR_RELOAD_POLICY, err.Error(), base.GetFunc())
		return errC
	}

	e.ClearPolicy()

	for _, v := range policy {
		e.AddPolicy(fmt.Sprintf("%d", v.PermissionID), v.Path, v.Method)
	}

	e.SavePolicy()

	//if len(e.GetPolicy()) == constant.ValueEmpty {
	//	logger.Error("AccessBaseService::LoadPolicy: -Load policy error by save policy error")
	//	errC = error_code.NewError(error_code.ERROR_RELOAD_POLICY, "load policy error by save policy error", base.GetFunc())
	//	return errC
	//}

	service.muAccess.Lock()
	service.Enforcer = e
	service.muAccess.Unlock()

	logger.Info("AccessController::LoadPolicy: - Policy: %v", e.GetPolicy())

	return nil
}

func (service *AccessBaseService) RunLoopLoadPolicy() {
	for {
		time.Sleep(time.Hour)
		logger.Info("AccessController::RunLoopLoadPolicy")
		if err := service.LoadPolicy(); err != nil {
			continue
		}
	}
}

func (service *AccessBaseService) InitPolicy() {
	if service.Enforcer == nil {
		if err := service.LoadPolicy(); err != nil {
			panic(err)
		}
	}

	go service.RunLoopLoadPolicy()
}
