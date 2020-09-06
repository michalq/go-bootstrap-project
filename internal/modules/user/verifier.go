package user

// Verifier service responsible for verificating users
type Verifier struct {
	usersRepository Users
}

// Verify goes through complicated process of verificating user account in system
func (a *Verifier) Verify(userEntity *User) error {
	userEntity.Verified = true
	return a.usersRepository.Update(userEntity)
}
