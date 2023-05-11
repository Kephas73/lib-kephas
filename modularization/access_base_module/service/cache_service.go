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
	KeyRole     string = "rbac:role:%s"
	KeyRoleUser string = "rbac:user:%s:role"
)

func (service *AccessBaseService) CacheSetRole(role *model.Role) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheSetRole: -Get connection error: %v", err)

		return err
	}

	b, err := json.Marshal(role)
	if err != nil {
		logger.Error("AccessBaseService:CacheSetRole: -Marshal error: %v", err)
		return err
	}

	return conn.Set(ctx, fmt.Sprintf(KeyRole, role.RoleID), string(b), time.Hour*24).Err()
}

func (service *AccessBaseService) CacheGetRole(roleID string) (*model.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheGetRole: -Get connection error: %v", err)

		return nil, err
	}

	rs, err := conn.Get(ctx, fmt.Sprintf(KeyRole, roleID)).Result()
	if err != nil {
		logger.Error("AccessBaseService:CacheGetRole: -Get redis error: %v", err)
		return nil, err
	}

	role := model.Role{}
	if err = json.Unmarshal([]byte(rs), &role); err != nil {
		logger.Error("AccessBaseService:CacheGetRole: -Get redis error: %v", err)
		return nil, err
	}

	return &role, nil
}

func (service *AccessBaseService) CacheDelRole(roleID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheDelRole: -Get connection error: %v", err)

		return err
	}

	return conn.Del(ctx, fmt.Sprintf(KeyRole, roleID)).Err()
}

//============================================================================================

func (service *AccessBaseService) CacheSetRoleUser(roleUser *model.UserRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheSetRoleUser: -Get connection error: %v", err)

		return err
	}

	b, err := json.Marshal(roleUser)
	if err != nil {
		logger.Error("AccessBaseService:CacheSetRoleUser: -Marshal error: %v", err)
		return err
	}

	return conn.Set(ctx, fmt.Sprintf(KeyRoleUser, roleUser.UserUUID), string(b), time.Hour*24).Err()
}

func (service *AccessBaseService) CacheGetRoleUser(userUUID string) (*model.UserRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), service.Timout)
	defer cancel()

	conn := service.Cache.Get()
	if conn == nil {
		err := fmt.Errorf("can not get connection")
		logger.Error("AccessBaseService:CacheGetRoleUser: -Get connection error: %v", err)

		return nil, err
	}

	rs, err := conn.Get(ctx, fmt.Sprintf(KeyRoleUser, userUUID)).Result()
	if err != nil {
		logger.Error("AccessBaseService:CacheGetRoleUser: -Get redis error: %v", err)
		return nil, err
	}

	roleUser := model.UserRole{}
	if err = json.Unmarshal([]byte(rs), &roleUser); err != nil {
		logger.Error("AccessBaseService:CacheGetRoleUser: -Get redis error: %v", err)
		return nil, err
	}

	return &roleUser, nil
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
