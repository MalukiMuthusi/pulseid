package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/gin-gonic/gin"
)

type Recall struct{}

func (h Recall) Handle(c *gin.Context) {
	var tokenParameter models.TokenParameter

	if err := c.ShouldBindUri(&tokenParameter); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.BasicError{
			Code:    "INVALID_TOKEN_PARAMETER",
			Message: "provide a valid token parameter",
		})

		return
	}

	res := map[string]interface{}{"message": "not implemented"}

	c.JSON(http.StatusNotImplemented, res)
}
