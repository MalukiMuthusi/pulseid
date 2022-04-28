package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Inactive struct{}

func (h Inactive) Handle(c *gin.Context) {
	res := map[string]interface{}{"message": "not implemented"}

	c.JSON(http.StatusNotImplemented, res)
}
