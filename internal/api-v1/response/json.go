package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/michalq/go-bootstrap-project/pkg/api"
)

// Response is general response returned by controllers
type Response struct {
	HTTPCode int
	Body     interface{}
}

// NewResponse creates new Response object
func NewResponse(httpCode int, body interface{}) *Response {
	return &Response{httpCode, body}
}

// JSONHandler produce response
func JSONHandler(ctx *gin.Context, resp *Response, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
		})
		fmt.Println(err)
		return
	}

	if resp.HTTPCode == http.StatusNoContent {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(resp.HTTPCode, resp.Body)
}
