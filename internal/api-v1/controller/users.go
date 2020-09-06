package controller

import (
	"context"

	"github.com/michalq/go-bootstrap-project/internal/api/formatter"
	"github.com/michalq/go-bootstrap-project/internal/modules/user"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// UsersController control flow for users
type UsersController struct {
	ctx             *context.Context
	usersRepository user.Users
	userFormatter   formatter.UserFormatter
	userVerifier *user.Verifier
}

// NewUserController basic contrustor
func NewUserController(
	ctx *context.Context,
	usersRepository user.Users,
	userFormatter formatter.UserFormatter,
) *UsersController {
	return &UsersController(ctx, usersRepository, userFormatter)
}

type UserControllerProvider interface {
	(*context.Context) *UsersController
}

// Details render details page
func (u *UsersController) Details(userID string) (interface{}, error) {
	user, err := u.usersRepository.FindOneByID(userID)
	if err != nil {
		return nil, err
	}
	return api.NewResponse(200, u.userFormatter.DomainToAPI(user)), nil
}

// CreateOne creates user
func (u *UsersController) CreateOne(apiUser *api.User) (interface{}, error) {
	user := u.userFormatter.APIToDomain(apiUser)
	if err := u.usersRepository.Save(user); err != nil {
		return nil, err
	}
	return
}

// Update updates user data
func (u *UsersController) Update(userID string, apiUser *api.User) (interface{}, error) {
	user := u.userFormatter.APIToDomain(apiUser)
	user.ID = userID
	if err := u.usersRepository.Update(user); err != nil {
		return nil, err
	}
	return
}

// Verify verify user
func (u *UsersController) Verify(userID string) (interface{}, error) {
	user := u.usersRepository.FindOne(userID)
	if err := u.userVerfier.Verify(apiUser); err != nil {
		return err
	}
	return
}

// List returns users
func (u *UsersController) List() {
	users, err := u.usersRepository.FindAll()
	if err != nil {
		return err
	}

}
