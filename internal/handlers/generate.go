package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/gin-gonic/gin"
)

// Generate Handler returns a token
type Generate struct{}

func (h Generate) Handle(c *gin.Context) {
	token, err := models.NewToken()

	if err != nil {
		logger.Log.Info(err)

		basicError := models.BasicError{
			Code:    "FAILED_GENERATE_TOKEN",
			Message: "service failed to generate code",
		}

		c.JSON(http.StatusInternalServerError, basicError)
	}

	c.JSON(http.StatusOK, token)
}
