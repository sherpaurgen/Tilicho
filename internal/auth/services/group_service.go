package services

import userGroupModel "github.com/sherpaurgen/Tilicho/internal/auth/models"

type GroupService interface {
	RegisterUser(user *userGroupModel.Group) error
	AuthenticateUser(email, password string) (*userGroupModel.Group, error)
	UpdateUserProfile(user *userGroupModel.Group) error
	DeleteUser(id string) error
	// Additional methods for other user-related operations
}
