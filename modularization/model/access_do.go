package model

import "time"

type Policy struct {
	PermissionID int        `json:"permission_id" db:"permission_id"`
	Method       string     `json:"method" db:"method"`
	Path         string     `json:"path" db:"path"`
	Name         string     `json:"name" db:"name"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type PermissionRole struct {
	RoleID       int        `json:"role_id" db:"role_id"`
	PermissionID int        `json:"permission_id" db:"permission_id"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type Permission struct {
	PermissionID int        `json:"permission_id" db:"permission_id"`
	Name         string     `json:"name" db:"name"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type Role struct {
	RoleID    int        `json:"role_id" db:"role_id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UserRole struct {
	UserUUID  string     `json:"user_uuid" db:"user_uuid"`
	RoleID    int        `json:"role_id" db:"role_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
