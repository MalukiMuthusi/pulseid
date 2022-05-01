package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/gin-gonic/gin"
)

type Recall struct {
	Store store.Store
}

func (h Recall) Handle(c *gin.Context) {
	var tokenParameter models.TokenParameter

	if err := c.ShouldBindUri(&tokenParameter); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.BasicError{
			Code:    "INVALID_TOKEN_PARAMETER",
			Message: "provide a valid token parameter",
		})

		return
	}

	token, err := h.Store.RecallToken(c.Request.Context(), tokenParameter.Token)
	if err != nil {
		basicError := models.BasicError{
			Code:    "FAILED_RECALL_TOKEN",
			Message: "failed to recall token, internall error occurred",
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return
	}

	c.JSON(http.StatusOK, models.RecallTokenResponse{
		Recall: token.Recalled,
	})

}
