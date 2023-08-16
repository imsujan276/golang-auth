package models

type Role struct {
	Name        string
	Permissions []string
}

var Roles = map[string]Role{
	"admin": {
		Name:        "admin",
		Permissions: []string{"read", "write", "update", "delete"},
	},
	"user": {
		Name:        "user",
		Permissions: []string{"read", "write", "update"},
	},
}
