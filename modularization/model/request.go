package model

import (
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/error_code"
)

type RoleReq struct {
	RoleID        string `json:"role_id" db:"role_id"`
	Name          string `json:"name" db:"name"`
	PermissionIDs []int  `json:"permission_ids" db:"permission_ids"`
	TeamName      string `json:"team_name" db:"team_name"`
	Description   string `json:"description" db:"description"`
}

func (req *RoleReq) Validate() *error_code.ErrorCode {
	if req.Name == constant.StrEmpty {
		errC := error_code.NewError(error_code.ERROR_DATA_INVALID, "name empty", base.GetFunc())
		return errC
	}

	if len(req.PermissionIDs) == constant.ValueEmpty {
		errC := error_code.NewError(error_code.ERROR_PERMISSION_EMPTY, "permission empty", base.GetFunc())
		return errC
	}

	return nil
}
