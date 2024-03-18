package models

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Groups   []*Group
}
