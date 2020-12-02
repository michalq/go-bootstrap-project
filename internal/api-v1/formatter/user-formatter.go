package formatter

import (
	"github.com/michalq/go-bootstrap-project/internal/modules/user"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// UserFormatter formats between domain and restAPI models
type UserFormatter struct{}

// NewUserFormatter creates instance of UserFormatter
func NewUserFormatter() *UserFormatter {
	return &UserFormatter{}
}

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

// DomainToAPIList formats list of users to api response
func (u *UserFormatter) DomainToAPIList(domainUsers []*user.User) []*api.User {
	var apiUsers []*api.User
	for _, domainUser := range domainUsers {
		apiUsers = append(apiUsers, u.DomainToAPI(domainUser))
	}
	return apiUsers
}
