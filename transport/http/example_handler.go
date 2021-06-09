package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/usecase"
)

type HttpExample struct {
	useCaseI usecase.ExampleUseCaseI
}

func NewUserHTTPHandler(r *gin.Engine, useCaseI usecase.ExampleUseCaseI) {
	handler := &HttpExample{useCaseI}

	api := r.Group("/api")
	{
		api.GET("/", handler.Check)
	}
}

func (handler *HttpExample) Check(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Gin Gonic Boilerplate V1",
		},
	)
}
