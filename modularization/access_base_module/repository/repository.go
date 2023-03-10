package repository

import (
	"context"
	"database/sql"
	"github.com/Kephas73/lib-kephas/logger"
	"github.com/Kephas73/lib-kephas/modularization/model"
	"github.com/jmoiron/sqlx"
	"sync"
)

type IAccessBaseRepository interface {
	SelectsPolicy(ctx context.Context) ([]*model.Policy, error)

	SelectPermissionByRole(ctx context.Context, roleID int) ([]*model.PermissionRole, error)

	SelectUserRoleByUserUUID(ctx context.Context, userUUID string) (*model.UserRole, error)
}

type AccessBaseRepository struct {
	*sqlx.DB
}

var (
	accessBaseRepositoryInstance *AccessBaseRepository
	muxAccessBaseRpInstance      sync.Mutex
)

func NewAccessBaseRepository(sqlx *sqlx.DB) IAccessBaseRepository {
	if accessBaseRepositoryInstance == nil {
		accessBaseRepository := AccessBaseRepository{
			sqlx,
		}

		muxAccessBaseRpInstance.Lock()
		accessBaseRepositoryInstance = &accessBaseRepository
		muxAccessBaseRpInstance.Unlock()
	}

	return accessBaseRepositoryInstance
}

func (repository *AccessBaseRepository) SelectsPolicy(ctx context.Context) ([]*model.Policy, error) {
	var query = `SELECT permission_id, method, path, name, created_at, updated_at FROM rbac_policy`

	rows, err := repository.QueryxContext(ctx, query)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsPolicy: -Query: %s, Error: %v", query, err)
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Policy, 0)
	for rows.Next() {
		do := &model.Policy{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsPolicy: -Query: %s, Error: %v", query, err)
			continue
		}

		res = append(res, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsPolicy: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	return res, nil
}

func (repository *AccessBaseRepository) SelectPermissionByRole(ctx context.Context, roleID int) ([]*model.PermissionRole, error) {
	var query = `SELECT role_id, permission_id, created_at, updated_at FROM rbac_permission_role WHERE role_id = ?`

	rows, err := repository.QueryxContext(ctx, query, roleID)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectPermissionByRole: -Query: %s, Error: %v", query, err)
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.PermissionRole, 0)
	for rows.Next() {
		do := &model.PermissionRole{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectPermissionByRole: -Query: %s, Error: %v", query, err)
			continue
		}

		res = append(res, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectPermissionByRole: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	return res, nil
}

func (repository *AccessBaseRepository) SelectUserRoleByUserUUID(ctx context.Context, userUUID string) (*model.UserRole, error) {
	var query = `SELECT user_uuid, role_id, created_at, updated_at FROM rbac_user WHERE user_uuid = ? LIMIT 1`
	userRole := model.UserRole{}
	err := repository.GetContext(ctx, &userRole, query, userUUID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("AccessBaseRepository::SelectUserRoleByUserUUID: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	return &userRole, nil
}
