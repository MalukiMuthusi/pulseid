package handlers

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Auth struct{}

func (m Auth) Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		var authHeader models.AuthHeader

		if err := c.ShouldBindHeader(&authHeader); err != nil {

			resp := models.BasicError{
				Code:    "INVALID_AUTH",
				Message: "provided authentication credentials",
			}

			c.JSON(http.StatusUnauthorized, resp)
			return
		}

		logger.Log.Info(authHeader.Authorization)

		authorization := strings.Split(authHeader.Authorization, " ")

		if len(authorization) != 2 {
			resp := models.BasicError{
				Code:    "INVALID_AUTH",
				Message: "provide authentication credentials",
			}

			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		credentials := authorization[1]

		b, err := base64.StdEncoding.DecodeString(credentials)
		if err != nil {
			resp := models.BasicError{
				Code:    "INVALID_AUTH",
				Message: "provide authentication credentials",
			}

			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		usernamePassword := string(b)

		s := strings.Split(usernamePassword, ":")
		if len(s) != 2 {
			resp := models.BasicError{
				Code:    "INVALID_AUTH",
				Message: "provide authentication credentials",
			}

			c.JSON(http.StatusUnprocessableEntity, resp)
			return
		}

		username := s[0]
		password := s[1]

		logger.Log.WithFields(logrus.Fields{"username": username, "password": password})

		c.Next()
	}
}
