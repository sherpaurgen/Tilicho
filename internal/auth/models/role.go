package models

type Role struct {
	Roleid      int64
	Name        string
	Description string
	Permissions []*Permission
	Groups      []*Group
}
