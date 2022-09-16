package models

type BaseUserManagersInterface interface {
	Create(name string) (err error)
	// UpdateUser(user User)
}
