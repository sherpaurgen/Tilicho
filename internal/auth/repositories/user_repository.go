package repositories

import userModel "github.com/sherpaurgen/Tilicho/internal/auth/models"

type UserRepository interface {
	CreateUser(user *userModel.User) error
	GetUserByID(id string) (*userModel.User, error)
	GetUserByEmail(email string) (*userModel.User, error)
	UpdateUser(user *userModel.User) error
	DeleteUser(id string) error
	// Additional methods for other CRUD operations on users
}
