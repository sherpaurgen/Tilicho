package services

import (
	"context"
	"fmt"

	userModel "github.com/sherpaurgen/Tilicho/internal/auth/models"
)

// this is repository interface Store
type Store interface {
	GetUserByUsername(context.Context, string) (userModel.User, error)
	// Additional methods for other user-related operations
}

type UserService struct {
	Store Store
}

// returns pointer to Service
func NewUserService(store Store) *UserService {
	return &UserService{Store: store}
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (userModel.User, error) {
	fmt.Println("retrieving user by username")
	usr, err := s.Store.GetUserByUsername(ctx, username)
	if err != nil {
		fmt.Println(err)
		return userModel.User{}, nil
	}
	return *usr, nil
}

func (s *UserService) CreateUser(ctx context.Context, userobj userModel.User) (userModel.User, error) {
	fmt.Println("creating user...")
	//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	usr, err := s.Store.CreateUser(ctx, userobj)
	if err != nil {
		fmt.Println("GetUserByUsername: Problem in dbquery:", err)
	}
	return usr, nil
}
