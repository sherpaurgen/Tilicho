package models

type Group struct {
	Groupid int64
	Name    string
	Roles   []*Role
	Users   []*User
}
