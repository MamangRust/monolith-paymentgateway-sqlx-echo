package repository

import (
	"github.com/jmoiron/sqlx"
)

// Repositories is a struct containing role command and query repositories.
type Repositories interface {
	RoleQueryRepository
	RoleCommandRepository
}

type repositories struct {
	RoleQueryRepository
	RoleCommandRepository
}

// NewRepositories creates a new instance of Repositories with the provided database
// connection. This repository is responsible for
// executing command and query operations related to role records in the database.
//
// Parameters:
//   - DB: A pointer to the sqlx.DB object for executing database queries.
//
// Returns:
//   - A pointer to the newly created Repositories instance.
func NewRepositories(DB *sqlx.DB) Repositories {
	return &repositories{
		RoleQueryRepository:   NewRoleQueryRepository(DB),
		RoleCommandRepository: NewRoleCommandRepository(DB),
	}
}
