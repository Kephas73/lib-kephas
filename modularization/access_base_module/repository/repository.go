package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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
	SelectsUserRoleByUserUUID(ctx context.Context, userUUID ...string) ([]*model.UserRole, error)
	InsertOrUpdateUserRole(ctx context.Context, userRole *model.UserRole) error
	DeleteUserRole(ctx context.Context, userUUID string) error

	SelectsRoleByID(ctx context.Context, id ...string) ([]*model.Role, error)
	SelectsRole(ctx context.Context, name string) ([]*model.Role, error)
	SelectRoleByID(ctx context.Context, id string) (*model.Role, error)
	InsertOrUpdateRole(ctx context.Context, role *model.Role) error
	DeleteRole(ctx context.Context, role *model.Role) error

	SelectsPermission(ctx context.Context, name string) ([]*model.Permission, error)
	SelectsPermissionByID(ctx context.Context, id ...int) ([]*model.Permission, error)
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

func (repository *AccessBaseRepository) SelectsUserRoleByUserUUID(ctx context.Context, userUUID ...string) ([]*model.UserRole, error) {
	var query = `SELECT user_uuid, role_id, created_at, updated_at FROM rbac_user WHERE user_uuid (?)`

	q, a, err := sqlx.In(query, userUUID)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsUserRoleByUserUUID: -Query: %s, Error: %v", q, err)
		return nil, err
	}

	rows, err := repository.QueryxContext(ctx, q, a...)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsUserRoleByUserUUID: -Query: %s, Error: %v", q, err)
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.UserRole, 0)
	for rows.Next() {
		do := &model.UserRole{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsUserRoleByUserUUID: -Query: %s, Error: %v", q, err)
			continue
		}

		res = append(res, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsUserRoleByUserUUID: -Query: %s, Error: %v", q, err)
		return nil, err
	}

	return res, nil
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

func (repository *AccessBaseRepository) DeleteUserRole(ctx context.Context, userUUID string) error {
	var query = `DELETE FROM rbac_user WHERE user_uuid = ?`

	_, err := repository.ExecContext(ctx, query, userUUID)
	if err != nil {
		logger.Error("AccessBaseRepository::DeleteUserRole: -Query: %s, Error: %v", query, err)
		return err
	}

	return nil
}

func (repository *AccessBaseRepository) SelectsRole(ctx context.Context, name string) ([]*model.Role, error) {
	var query = `SELECT role_id, name, permission_ids, team_name, description, created_at, updated_at FROM rbac_role `

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

	resDB := make([]*model.RoleDB, 0)
	for rows.Next() {
		do := &model.RoleDB{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsRole: -Query: %s, Error: %v", query, err)
			continue
		}

		resDB = append(resDB, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsRole: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	res := make([]*model.Role, 0)
	b, _ := json.Marshal(resDB)
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repository *AccessBaseRepository) SelectsRoleByID(ctx context.Context, id ...string) ([]*model.Role, error) {

	var query = `SELECT role_id, name, permission_ids, team_name, description, created_at, updated_at FROM rbac_role WHERE role_id IN (?)`

	q, a, err := sqlx.In(query, id)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsRoleByID: -Query: %s, Error: %v", q, err)
		return nil, err
	}

	rows, err := repository.QueryxContext(ctx, q, a...)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsRoleByID: -Query: %s, Error: %v", q, err)
		return nil, err
	}
	defer rows.Close()

	resDB := make([]*model.RoleDB, 0)
	for rows.Next() {
		do := &model.RoleDB{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsRoleByID: -Query: %s, Error: %v", q, err)
			continue
		}

		resDB = append(resDB, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsRoleByID: -Query: %s, Error: %v", q, err)
		return nil, err
	}

	res := make([]*model.Role, 0)
	b, _ := json.Marshal(resDB)
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repository *AccessBaseRepository) SelectRoleByID(ctx context.Context, id string) (*model.Role, error) {
	var query = `SELECT role_id, name, permission_ids, team_name, description, created_at, updated_at FROM rbac_role WHERE role_id = ?`

	roleDB := model.RoleDB{}
	err := repository.GetContext(ctx, &roleDB, query, id)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("AccessBaseRepository::SelectRoleByID: -Query: %s, Error: %v", query, err)
		return nil, err
	}

	role := model.Role{}
	b, _ := json.Marshal(roleDB)
	err = json.Unmarshal(b, &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (repository *AccessBaseRepository) InsertOrUpdateRole(ctx context.Context, role *model.Role) error {

	var roleDB model.RoleDB
	b, _ := json.Marshal(role)
	err := json.Unmarshal(b, &roleDB)
	if err != nil {
		return err
	}

	var query = `INSERT INTO rbac_role (role_id, name, permission_ids, team_name, description) VALUES (:role_id, :name, :permission_ids, :team_name, :description) 
				ON DUPLICATE KEY UPDATE name = :name, permission_ids = :permission_ids, team_name = :team_name, description = :description`

	_, err = repository.NamedExecContext(ctx, query, roleDB)
	if err != nil {
		logger.Error("AccessBaseRepository::InsertOrUpdate: -Query: %s, Error: %v", query, err)
		return err
	}

	return nil
}

func (repository *AccessBaseRepository) DeleteRole(ctx context.Context, role *model.Role) error {
	var queryRole = `DELETE FROM rbac_role WHERE role_id = :role_id`

	// TODO: có nên để lại role đấy cho user không. Vẫn có thể để lại và hiển thị user ko có gì

	_, err := repository.NamedExecContext(ctx, queryRole, role)
	if err != nil {
		logger.Error("AccessBaseRepository::DeleteRole: -QueryRole: %s, Error: %v", queryRole, err)
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

func (repository *AccessBaseRepository) SelectsPermissionByID(ctx context.Context, id ...int) ([]*model.Permission, error) {
	var query = `SELECT permission_id, name, created_at, updated_at FROM rbac_permission WHERE permission_id IN (?)`

	q, a, err := sqlx.In(query, id)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsPermissionByID: -Query: %s, Error: %v", q, err)
		return nil, err
	}

	rows, err := repository.QueryxContext(ctx, q, a...)
	if err != nil {
		logger.Error("AccessBaseRepository::SelectsPermissionByID: -Query: %s, Error: %v", q, err)
		return nil, err
	}
	defer rows.Close()

	res := make([]*model.Permission, 0)
	for rows.Next() {
		do := &model.Permission{}
		if err = rows.StructScan(do); err != nil {
			logger.Error("AccessBaseRepository::SelectsPermissionByID: -Query: %s, Error: %v", q, err)
			continue
		}

		res = append(res, do)
	}

	if err = rows.Err(); err != nil {
		logger.Error("AccessBaseRepository::SelectsPermissionByID: -Query: %s, Error: %v", q, err)
		return nil, err
	}

	return res, nil
}
