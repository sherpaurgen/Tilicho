package models

type Role struct {
	ID          int64
	Name        string
	Description string
	Permissions []*Permission
	Groups      []*Group
}
