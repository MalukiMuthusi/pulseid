package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ActiveHandler returns the tokens that are currently active
type Active struct{}

func (h Active) Handle(c *gin.Context) {
	res := map[string]interface{}{"message": "not implemented"}

	c.JSON(http.StatusNotImplemented, res)
}
