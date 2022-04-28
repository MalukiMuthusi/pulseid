package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Generate Handler returns a token
type Generate struct{}

func (h Generate) Handle(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, map[string]interface{}{"message": "not implemented"})
}
