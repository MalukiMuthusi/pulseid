package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/gin-gonic/gin"
)

type Inactive struct {
	Store store.Store
}

func (h Inactive) Handle(c *gin.Context) {

	tokens, err := h.Store.Inactive(c.Request.Context())
	if err != nil {
		basicError := models.BasicError{
			Code:    "FAILED_GET_INACTIVE_TOKENS",
			Message: "internal error, failed to get inactive tokens",
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return
	}

	c.JSON(http.StatusOK, tokens)
}
