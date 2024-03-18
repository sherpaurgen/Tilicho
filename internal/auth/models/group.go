package models

type Group struct {
	ID    int64
	Name  string
	Roles []*Role
	Users []*User
}
