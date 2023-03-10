package model

import "time"

type Policy struct {
	PermissionID int        `json:"permission_id" db:"permission_id"`
	Path         string     `json:"path" db:"path"`
	Method       string     `json:"method" db:"method"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type PermissionRole struct {
	RoleID       int        `json:"role_id" db:"role_id"`
	PermissionID int        `json:"permission_id" db:"permission_id"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type UserRole struct {
	UserUUID  string     `json:"user_uuid" db:"user_uuid"`
	RoleID    int        `json:"role_id" db:"role_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
