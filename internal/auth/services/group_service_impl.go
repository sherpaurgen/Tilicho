package services

// import (
// 	userGroupModel "github.com/sherpaurgen/Tilicho/internal/auth/models"
// 	"github.com/sherpaurgen/Tilicho/internal/auth/repositories"
// 	"golang.org/x/crypto/bcrypt"
// )

// type GroupServiceImpl struct {
// 	groupRepo repositories.UserRepository
// }

// func NewGroupService(groupRepo repositories.UserRepository) *GroupServiceImpl {
// 	return &GroupServiceImpl{groupRepo: groupRepo}
// }

// func (s *GroupServiceImpl) RegisterUser(user *userGroupModel.User) error {
// 	// Hash the user's password before storing
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(hashedPassword)

// 	// Create the user in the database
// 	err = s.groupRepo.CreateUser(user)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Implementations for other service methods...
