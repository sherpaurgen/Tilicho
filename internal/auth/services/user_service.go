package services

import userModel "github.com/sherpaurgen/Tilicho/internal/auth/models"

type UserService interface {
	RegisterUser(user *userModel.User) error
	AuthenticateUser(email, password string) (*userModel.User, error)
	UpdateUserProfile(user *userModel.User) error
	DeleteUser(id string) error
	// Additional methods for other user-related operations
}
