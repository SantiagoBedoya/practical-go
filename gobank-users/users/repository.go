package users

// Repository define interface for repositories
type Repository interface {
	Create(*User) (*User, error)
	FindByID(string) (*User, error)
	FindByEmail(*User) (*User, error)
}
