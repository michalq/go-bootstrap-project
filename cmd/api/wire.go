//+build wireinject

package main

import (
	"github.com/google/wire"
	apiv1 "github.com/michalq/go-bootstrap-project/internal/api-v1"
	"github.com/michalq/go-bootstrap-project/internal/api-v1/controller"
	"github.com/michalq/go-bootstrap-project/internal/api-v1/formatter"
	"go.mongodb.org/mongo-driver/mongo"

	userDAL "github.com/michalq/go-bootstrap-project/internal/dal/user"

	"github.com/michalq/go-bootstrap-project/internal/modules/user"
)

func InitializeApi(mongoDB *mongo.Database) (*apiv1.V1Routing, error) {
	wire.Build(
		apiv1.NewV1Routing,

		controller.NewContextScopedUserController,
		formatter.NewUserFormatter,

		user.NewVerifier,
		userDAL.NewUsers,

		wire.Bind(new(user.Users), new(*userDAL.Users)),
	)
	return &apiv1.V1Routing{}, nil
}
