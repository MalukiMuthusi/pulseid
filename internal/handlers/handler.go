package handlers

import (
	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUpRouter(store store.Store) *gin.Engine {

	r := gin.New()

	gin.DebugPrintRouteFunc = DebugPrintRoute

	auth := Auth{}

	generate := Generate{Store: store}
	generateAPI := r.Group("/generate")
	generateAPI.Use(auth.Middleware())
	generateAPI.GET("", generate.Handle)

	validate := Validate{
		Store: store,
	}
	r.GET("/validate/:token", validate.Handle)

	recall := Recall{}
	r.GET("/recall/:token", recall.Handle)

	active := Active{}
	r.GET("/active", active.Handle)

	inactive := Inactive{}
	r.GET("/inactive", inactive.Handle)

	return r
}

func DebugPrintRoute(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	logger.Log.WithFields(logrus.Fields{"httpMethod": httpMethod, "path": absolutePath, "handlerName": handlerName, "nuHandlers": nuHandlers}).Info("endpointRequest")
}
