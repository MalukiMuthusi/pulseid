package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/gin-gonic/gin"
)

type Validate struct {
	Store store.Store
}

func (h Validate) Handle(c *gin.Context) {

	var tokenParameter models.TokenParameter

	if err := c.ShouldBindUri(&tokenParameter); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.BasicError{
			Code:    "INVALID_TOKEN_PARAMETER",
			Message: "provide a valid token parameter",
		})

		return
	}

	token, err := h.Store.GetToken(c.Request.Context(), tokenParameter.Token)
	if err != nil {
		basicError := models.BasicError{
			Code:    "TOKEN_NOT_FOUND",
			Message: "token not found",
		}

		c.JSON(http.StatusInternalServerError, basicError)
		return
	}

	validity := token.CheckValidity()

	tokenValidity := models.TokenValidity{
		Validity: validity,
	}

	c.JSON(http.StatusOK, tokenValidity)

}
