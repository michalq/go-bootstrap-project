package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/michalq/go-bootstrap-project/internal/api-v1/controller"
	"github.com/michalq/go-bootstrap-project/internal/api-v1/response"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// V1Routing handles requests for api/v1
type V1Routing struct {
	userControllerProvider controller.ContextScoperUserController
}

// NewV1Routing creates new instance of V1Routing
func NewV1Routing(userControllerProvider controller.ContextScoperUserController) *V1Routing {
	return &V1Routing{userControllerProvider}
}

// SetupRoutes is responsible for setting api/v1 routing
func (v *V1Routing) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/users/:userID",
		func(c *gin.Context) {
			resp, err := v.userControllerProvider(c).DetailsAction(c.Param("userID"))
			response.JSONHandler(c, resp, err)
		},
	)
	v1.GET("/users",
		func(c *gin.Context) {
			resp, err := v.userControllerProvider(c).ListAction()
			response.JSONHandler(c, resp, err)
		},
	)
	v1.POST("/users",
		func(c *gin.Context) {
			apiUser := &api.User{}
			if err := c.ShouldBind(apiUser); err != nil {
				response.JSONHandler(c, nil, err)
				return
			}
			resp, err := v.userControllerProvider(c).CreateAction(apiUser)
			response.JSONHandler(c, resp, err)
		},
	)
	v1.POST("/users/:userID",
		func(c *gin.Context) {
			apiUser := &api.User{}
			if err := c.ShouldBind(apiUser); err != nil {
				response.JSONHandler(c, nil, err)
				return
			}
			resp, err := v.userControllerProvider(c).UpdateAction(c.Param("userID"), apiUser)
			response.JSONHandler(c, resp, err)
		},
	)
	v1.POST("/users/:userID/verifications",
		func(c *gin.Context) {
			resp, err := v.userControllerProvider(c).VerifyAction(c.Param("userID"))
			response.JSONHandler(c, resp, err)
		},
	)
}
