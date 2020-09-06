package formatter

import (
	"github.com/michalq/go-bootstrap-project/internal/modules/user"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// UserFormatter formats between domain and restAPI models
type UserFormatter struct{}

// DomainToAPI formats user from domain model to api model
func (u *UserFormatter) DomainToAPI(domainUser *user.User) *api.User {
	return &api.User{
		Id:   domainUser.ID,
		Name: domainUser.Name,
	}
}

// APIToDomain formats user from api to domain model
func (u *UserFormatter) APIToDomain(apiUser *api.User) *user.User {
	return &user.User{
		ID:   apiUser.Id,
		Name: apiUser.Name,
	}
}
