package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/logger"
	"github.com/Kephas73/lib-kephas/modularization/model"
	"github.com/jmoiron/sqlx"
	"strings"
	"sync"
)

type IAccessBaseRepository interface {
	SelectsPolicy(ctx context.Context) ([]*model.Policy, error)

	SelectUserRoleByUserUUID(ctx context.Context, userUUID string) (*model.UserRole, error)
	InsertOrUpdateUserRole(ctx context.Context, userRole *model.UserRole) error

	// Role
	SelectsRole(ctx context.Context, name string) ([]*model.Role, error)
	InsertOrUpdateRole(ctx context.Context, role *model.Role) error
	DeleteRole(ctx context.Context, role *model.Role) error

	// Detail permission role
	SelectPermissionByRole(ctx context.Context, roleID int) ([]*model.PermissionRole, error)
	UpdatePermissionRole(ctx context.Context, roleID int, permissionID []int) error

	// Permission
	SelectsPermission(ctx context.Context, name string) ([]*model.Permission, error)
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

func (repository *AccessBaseRepository) InsertOrUpdateUserRole(ctx context.Context, userRole *model.UserRole) error {
	var query = `INSERT INTO rbac_user (user_uuid, role_id) VALUES (:user_uuid, :role_id) ON DUPLICATE KEY UPDATE role_id = :role_id`

	_, err := repository.NamedExecContext(ctx, query, userRole)
	if err != nil {
		logger.Error("AccessBaseRepository::InsertOrUpdateUserRole: -Query: %s, Error: %v", query, err)
		return err
	}

	return nil
}

func (repository *AccessBaseRepository) SelectsRole(ctx context.Context, name string) ([]*model.Role, error) {
	var query = `SELECT role_id, name, created_at, updated_at FROM rbac_role `

	var (
		where []string
	)

	if name != constant.StrEmpty {
		where = append(where, fmt.Sprintf("name LIKE '%s'", "%"+name+"%"))
	}

	if len(where) != constant.ValueEmpty {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	query += " ORDER BY role_id ASC, updated_at DESC "

	rows, err := repository.QueryxContext(ctx, query)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsRole: -Query: %s, Error: %v", query, err)
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Role, 0)
	for rows.Next() {
		do := &model.Role{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsRole: -Query: %s, Error: %v", query, err)
			continue
		}

		res = append(res, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsRole: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	return res, nil
}

func (repository *AccessBaseRepository) InsertOrUpdateRole(ctx context.Context, role *model.Role) error {
	var query = `INSERT INTO rbac_role (role_id, name) VALUES (:role_id, :name) ON DUPLICATE KEY UPDATE name = :name`

	_, err := repository.NamedExecContext(ctx, query, role)
	if err != nil {
		logger.Error("AccessBaseRepository::InsertOrUpdate: -Query: %s, Error: %v", query, err)
		return err
	}

	return nil
}

func (repository *AccessBaseRepository) DeleteRole(ctx context.Context, role *model.Role) error {
	var queryRole = `DELETE FROM rbac_role WHERE role_id = :role_id`

	var queryPermissionRole = `DELETE FROM rbac_permission_role WHERE role_id = :role_id`

	tx, err := repository.Beginx()
	if err != nil {
		logger.Error("AccessBaseRepository::DeleteRole: -Error: %v", err)
		return err
	}

	if _, err = tx.NamedExecContext(ctx, queryRole, role); err != nil {
		logger.Error("AccessBaseRepository::DeleteRole: -QueryRole: %s, Error: %v", queryRole, err)
		tx.Rollback()
		return err
	}

	if _, err = tx.NamedExecContext(ctx, queryPermissionRole, role); err != nil {
		logger.Error("AccessBaseRepository::DeleteRole: -QueryPermission: %s, Error: %v", queryPermissionRole, err)
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		logger.Error("AccessBaseRepository::DeleteRole: -Error: %v", err)
		tx.Rollback()
		return err
	}

	return nil
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

func (repository *AccessBaseRepository) UpdatePermissionRole(ctx context.Context, roleID int, permissionID []int) error {
	var queryClear = `DELETE FROM rbac_permission_role WHERE role_id = ?`

	var queryPermissionRole = `INSERT INTO rbac_permission_role(role_id, permission_id) VALUES `

	var values []string

	for _, v := range permissionID {
		values = append(values, fmt.Sprintf("(%d, %d)", roleID, v))
	}

	queryPermissionRole += strings.Join(values, ", ")

	tx, err := repository.Beginx()
	if err != nil {
		logger.Error("AccessBaseRepository::UpdatePermissionRole: -Error: %v", err)
		return err
	}

	if _, err = tx.ExecContext(ctx, queryClear, roleID); err != nil {
		logger.Error("AccessBaseRepository::UpdatePermissionRole: -Query clear: %s, Error: %v", queryClear, err)
		tx.Rollback()
		return err
	}

	if len(permissionID) != constant.ValueEmpty {
		if _, err = tx.ExecContext(ctx, queryPermissionRole); err != nil {
			logger.Error("AccessBaseRepository::UpdatePermissionRole: -Query permission role: %s, Error: %v", queryPermissionRole, err)
			tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		logger.Error("AccessBaseRepository::UpdatePermissionRole: -Error: %v", err)
		tx.Rollback()
		return err
	}

	return nil

}

func (repository *AccessBaseRepository) SelectsPermission(ctx context.Context, name string) ([]*model.Permission, error) {
	var query = `SELECT permission_id, name, created_at, updated_at FROM rbac_permission `

	var (
		where []string
	)

	if name != constant.StrEmpty {
		where = append(where, fmt.Sprintf("name LIKE '%s'", "%"+name+"%"))
	}

	if len(where) != constant.ValueEmpty {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	query += " ORDER BY permission_id ASC, updated_at DESC "

	rows, err := repository.QueryxContext(ctx, query)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsPermission: -Query: %s, Error: %v", query, err)
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Permission, 0)
	for rows.Next() {
		do := &model.Permission{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsPermission: -Query: %s, Error: %v", query, err)
			continue
		}

		res = append(res, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsPermission: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	return res, nil
}
