package services

// import (
// 	"errors"

// 	userGroupModel "github.com/sherpaurgen/Tilicho/internal/auth/models"
// 	"github.com/sherpaurgen/Tilicho/internal/auth/repositories"
// 	"golang.org/x/crypto/bcrypt"
// )

// type UserServiceImpl struct {
// 	userRepo repositories.UserRepository
// }

// func NewUserService(userRepo repositories.UserRepository) *UserServiceImpl {
// 	return &UserServiceImpl{userRepo: userRepo}
// }

// func (s *UserServiceImpl) RegisterUser(user *userGroupModel.User) error {
// 	// Hash the user's password before storing
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(hashedPassword)

// 	// Create the user in the database
// 	err = s.userRepo.CreateUser(user)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *UserServiceImpl) AuthenticateUser(email, password string) (*userGroupModel.User, error) {
// 	// Fetch the user by email from the database
// 	user, err := s.userRepo.GetUserByEmail(email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Compare the provided password with the stored hash
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		return nil, errors.New("invalid credentials")
// 	}

// 	return user, nil
// }

// // Implementations for other service methods..
