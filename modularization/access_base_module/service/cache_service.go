package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kephas73/lib-kephas/logger"
	"github.com/Kephas73/lib-kephas/modularization/model"
	"time"
)

const (
	KeyPermissionRole string = "rbac:role:%d:permission"
	KeyRoleUser       string = "rbac:user:%s:role"
)

func (service *AccessBaseService) CacheSetPermissionRole(permissionRole []*model.PermissionRole, roleID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheSetPermissionRole: -Get connection error: %v", err)

		return err
	}

	b, err := json.Marshal(permissionRole)
	if err != nil {
		logger.Error("AccessBaseService:CacheSetPermissionRole: -Marshal error: %v", err)
		return err
	}

	return conn.Set(ctx, fmt.Sprintf(KeyPermissionRole, roleID), string(b), time.Hour*24).Err()
}

func (service *AccessBaseService) CacheGetPermissionRole(roleID int) ([]*model.PermissionRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheGetPermissionRole: -Get connection error: %v", err)

		return nil, err
	}

	rs, err := conn.Get(ctx, fmt.Sprintf(KeyPermissionRole, roleID)).Result()
	if err != nil {
		logger.Error("AccessBaseService:CacheGetPermissionRole: -Get redis error: %v", err)
		return nil, err
	}

	permissionRole := make([]*model.PermissionRole, 0)
	if err = json.Unmarshal([]byte(rs), &permissionRole); err != nil {
		logger.Error("AccessBaseService:CacheGetPermissionRole: -Get redis error: %v", err)
		return nil, err
	}

	return permissionRole, nil
}

func (service *AccessBaseService) CacheDelPermissionRole(roleID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheDelPermissionRole: -Get connection error: %v", err)

		return err
	}

	return conn.Del(ctx, fmt.Sprintf(KeyPermissionRole, roleID)).Err()
}

//============================================================================================

func (service *AccessBaseService) CacheSetRoleUser(permissionRole []*model.PermissionRole, roleID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheSetPermissionRole: -Get connection error: %v", err)

		return err
	}

	b, err := json.Marshal(permissionRole)
	if err != nil {
		logger.Error("AccessBaseService:CacheSetPermissionRole: -Marshal error: %v", err)
		return err
	}

	return conn.Set(ctx, fmt.Sprintf(KeyPermissionRole, roleID), string(b), time.Hour*24).Err()
}

func (service *AccessBaseService) CacheGetRoleUser(roleID int) ([]*model.UserRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheGetPermissionRole: -Get connection error: %v", err)

		return nil, err
	}

	rs, err := conn.Get(ctx, fmt.Sprintf(KeyPermissionRole, roleID)).Result()
	if err != nil {
		logger.Error("AccessBaseService:CacheGetPermissionRole: -Get redis error: %v", err)
		return nil, err
	}

	roleUser := make([]*model.UserRole, 0)
	if err = json.Unmarshal([]byte(rs), &roleUser); err != nil {
		logger.Error("AccessBaseService:CacheGetPermissionRole: -Get redis error: %v", err)
		return nil, err
	}

	return roleUser, nil
}

func (service *AccessBaseService) CacheDelRoleUser(userUUID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheDelRoleUser: -Get connection error: %v", err)

		return err
	}

	return conn.Del(ctx, fmt.Sprintf(KeyRoleUser, userUUID)).Err()
}
