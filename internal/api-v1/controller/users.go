package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/michalq/go-bootstrap-project/internal/api-v1/formatter"
	"github.com/michalq/go-bootstrap-project/internal/api-v1/response"
	"github.com/michalq/go-bootstrap-project/internal/modules/user"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// UsersController control flow for users
type UsersController struct {
	ctx             context.Context
	usersRepository user.Users
	userFormatter   *formatter.UserFormatter
	userVerifier    *user.Verifier
}

// ContextScoperUserController context scoped provider for controller
type ContextScoperUserController func(context.Context) *UsersController

// NewContextScopedUserController creates context-scoped provider, that can be used in multiple context
func NewContextScopedUserController(
	usersRepository user.Users,
	userFormatter *formatter.UserFormatter,
	userVerifier *user.Verifier,
) ContextScoperUserController {
	return func(ctx context.Context) *UsersController {
		return &UsersController{ctx, usersRepository, userFormatter, userVerifier}
	}
}

// DetailsAction render details page
func (u *UsersController) DetailsAction(userID string) (*response.Response, error) {
	user, err := u.usersRepository.FindOneByID(userID)
	if err != nil {
		return nil, err
	}
	return response.NewResponse(http.StatusOK, u.userFormatter.DomainToAPI(user)), nil
}

// CreateAction creates user
func (u *UsersController) CreateAction(apiUser *api.User) (*response.Response, error) {
	userModel := u.userFormatter.APIToDomain(apiUser)
	if _, err := u.usersRepository.Save(userModel); err != nil {
		return nil, err
	}
	return response.NewResponse(http.StatusCreated, api.UserCreated{Id: userModel.ID}), nil
}

// UpdateAction updates user data
func (u *UsersController) UpdateAction(userID string, apiUser *api.User) (*response.Response, error) {
	userModel, err := u.usersRepository.FindOneByID(userID)
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, errors.New("user not found")
	}
	userModel = u.userFormatter.APIToDomain(apiUser)
	userModel.ID = userID
	if err := u.usersRepository.Update(userModel); err != nil {
		return nil, err
	}
	return response.NewResponse(http.StatusNoContent, nil), nil
}

// VerifyAction verify user
func (u *UsersController) VerifyAction(userID string) (*response.Response, error) {
	if _, err := u.userVerifier.FullVerification(userID); err != nil {
		return nil, err
	}
	return response.NewResponse(http.StatusNoContent, nil), nil
}

// ListAction returns users
func (u *UsersController) ListAction() (*response.Response, error) {
	users, err := u.usersRepository.FindAll()
	if err != nil {
		return nil, err
	}
	usersFormatted := u.userFormatter.DomainToAPIList(users)
	return response.NewResponse(http.StatusOK, usersFormatted), nil
}
