package database

//repo layer

import (
	"context"
	"database/sql"
	"fmt"

	userModel "github.com/sherpaurgen/Tilicho/internal/auth/models"
)

type UserRow struct {
	Userid   string
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
	Groups   sql.NullString
	IsActive sql.NullString
}

type UserRepository interface {
	GetUserByUsername(context.Context, string) (*userModel.User, error)
}

func convertUserRowToUserStruct(u UserRow) userModel.User {
	return userModel.User{
		Userid:   u.Userid,
		Username: u.Username.String,
		Email:    u.Email.String,
		Groups:   u.Groups.String,
		IsActive: u.IsActive.String,
	}
}

func (db *Database) GetUserByUsername(ctx context.Context, username string) (userModel.User, error) {
	var userRow UserRow
	//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)

	err := db.Pool.QueryRow(context.Background(), "SELECT userid,username,email,groups,isactive FROM users_table WHERE username=$1", username).Scan(
		&userRow.Userid, &userRow.Username, &userRow.Email, &userRow.Groups, &userRow.IsActive)
	if err != nil {
		fmt.Println("GetUserByUsername: Problem in dbquery:", err)
	}
	return convertUserRowToUserStruct(userRow), nil
}

func (db *Database) CreateUser(ctx context.Context, userobj userModel.User) (userModel.User, error) {

	//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	query := `
	INSERT INTO users_table (username, email, password, groups, isactive)
	VALUES ($1, $2, $3, $4, $5)`
	err := db.Pool.QueryRow(context.Background(), query, userobj.Username, userobj.Email, userobj.Password, userobj.Groups, userobj.IsActive)
	if err != nil {
		fmt.Println("GetUserByUsername: Problem in dbquery:", err)
	}
	return userobj, nil
}
