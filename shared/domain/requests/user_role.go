package requests

// CreateUserRoleRequest represents the payload for assigning a role to a user.
// Used when granting specific permissions or access levels to users.
type CreateUserRoleRequest struct {
	UserId int `json:"user_id" validate:"required"` // ID of the user to receive the role
	RoleId int `json:"role_id" validate:"required"` // ID of the role to be assigned
}

// RemoveUserRoleRequest represents the payload for revoking a role from a user.
// Used when removing specific permissions or access levels from users.
type RemoveUserRoleRequest struct {
	UserId int `json:"user_id" validate:"required"` // ID of the user losing the role
	RoleId int `json:"role_id" validate:"required"` // ID of the role to be removed
}
