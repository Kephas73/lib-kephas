package model

import "time"

type UserRoleRes struct {
	UserUUID  string     `json:"user_uuid" db:"user_uuid"`
	RoleID    string     `json:"role_id" db:"role_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	RoleInfo  Role       `json:"role_info"`
}
