package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/michalq/go-bootstrap-project/internal/api/controller"
	"github.com/michalq/go-bootstrap-project/internal/api/response"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// V1 handles requests for api/v1
type V1 struct {
	userControllerProvider *controller.UserControllerProvider
}

// SetupRouting is responsible for setting api/v1 routing
func (v *V1) SetupRouting(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET(
		"/users/:userID",
		func(c *gin.Context) {
			resp, err := v.userControllerProvider(c).Details(c.Param("userID"))
			response.JSONHandler(c, resp, err)
		},
	)
	v1.GET(
		"/users",
		func(c *gin.Context) {
			resp, err := v.userControllerProvider(c).List()
			response.JSONHandler(c, resp, err)
		},
	)
	v1.POST(
		"/users",
		func(c *gin.Context) {
			resp, err := v.userControllerProvider(c).CreateOne(c.ShouldBind(api.User))
			response.JSONHandler(c, resp, err)
		},
	)
}
