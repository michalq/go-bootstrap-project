package user

import (
	"errors"
)

var (
	// ErrUserAlreadyVerified is returned when user is trying to verify but is already verified
	ErrUserAlreadyVerified = errors.New("user already verified")
	// ErrUserWithoutName is returned when user doesn't provide name and tries to verify
	ErrUserWithoutName = errors.New("user did not provide name")
)

// Verifier service responsible for verificating users
type Verifier struct {
	usersRepository Users
}

// NewVerifier creates new instance of Verifier
func NewVerifier(usersRepository Users) *Verifier {
	return &Verifier{usersRepository}
}

// FullVerification process full verification
func (a *Verifier) FullVerification(userID string) (*User, error) {
	userModel, err := a.usersRepository.FindOneByID(userID)
	if err != nil {
		return nil, err
	}
	err = a.Verify(userModel)
	if err != nil {
		return nil, err
	}
	userModel.Verified = true
	if err := a.usersRepository.Update(userModel); err != nil {
		return nil, err
	}
	return userModel, nil
}

// Verify simply verify conditions
func (a *Verifier) Verify(userEntity *User) error {
	if userEntity.Verified == true {
		return ErrUserAlreadyVerified
	}
	if userEntity.Name == "" {
		return ErrUserWithoutName
	}
	return nil
}
