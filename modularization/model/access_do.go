package model

import (
	"encoding/json"
	"time"
)

type Policy struct {
	PermissionID int        `json:"permission_id" db:"permission_id"`
	Method       string     `json:"method" db:"method"`
	Path         string     `json:"path" db:"path"`
	Name         string     `json:"name" db:"name"`
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
	RoleID        string     `json:"role_id" db:"role_id"`
	Name          string     `json:"name" db:"name"`
	PermissionIDs []int      `json:"permission_ids" db:"permission_ids"`
	TeamName      string     `json:"team_name" db:"team_name"`
	Description   string     `json:"description" db:"description"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
}

type RoleDB struct {
	RoleID        string          `json:"role_id" db:"role_id"`
	Name          string          `json:"name" db:"name"`
	PermissionIDs json.RawMessage `json:"permission_ids" db:"permission_ids"`
	TeamName      string          `json:"team_name" db:"team_name"`
	Description   string          `json:"description" db:"description"`
	CreatedAt     *time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at" db:"updated_at"`
}

type UserRole struct {
	UserUUID  string     `json:"user_uuid" db:"user_uuid"`
	RoleID    string     `json:"role_id" db:"role_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
