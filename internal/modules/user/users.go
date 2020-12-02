package user

// Users shares method to manage and retrieve users
type Users interface {
	// FindOneByID Returns user by id
	FindOneByID(userID string) (*User, error)

	// FindAll finds all created users
	FindAll() ([]*User, error)

	// Save saves user in database
	Save(*User) (string, error)

	// Save saves user in database
	Update(*User) error
}
