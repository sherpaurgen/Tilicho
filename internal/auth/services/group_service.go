package services

import groupModel "github.com/sherpaurgen/Tilicho/internal/auth/models"

type GroupService interface {
	RegisterUser(user *groupModel.Group) error
	AuthenticateUser(email, password string) (*groupModel.Group, error)
	UpdateUserProfile(user *groupModel.Group) error
	DeleteUser(id string) error
	// Additional methods for other user-related operations
}
