package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/gin-gonic/gin"
)

// ActiveHandler returns the tokens that are currently active
type Active struct {
	Store store.Store
}

func (h Active) Handle(c *gin.Context) {

	tokens, err := h.Store.Active(c.Request.Context())

	if err != nil {
		basicError := models.BasicError{
			Code:    "FAILED_GET_ACTIVE_TOKENS",
			Message: "failed to get the active tokens",
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return
	}

	c.JSON(http.StatusOK, tokens)
}
