package handlers

import (
	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/MalukiMuthusi/pulseid/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUpRouter(store store.Store, debugPrintRoute DebugPrintRouteFunc) *gin.Engine {

	r := gin.New()

	gin.DebugPrintRouteFunc = debugPrintRoute

	auth := Auth{}

	generate := Generate{Store: store}
	generateAPI := r.Group("/generate")
	generateAPI.Use(auth.Middleware())
	generateAPI.GET("", generate.Handle)

	validate := Validate{
		Store: store,
	}
	r.GET("/validate/:token", validate.Handle)

	recall := Recall{Store: store}
	recallAPI := r.Group("/recall/:token")
	recallAPI.Use(auth.Middleware())
	recallAPI.GET("", recall.Handle)

	active := Active{Store: store}
	r.GET("/active", active.Handle)

	inactive := Inactive{Store: store}
	inactiveAPI := r.Group("/inactive")
	inactiveAPI.Use(auth.Middleware())
	inactiveAPI.GET("", inactive.Handle)

	return r
}

type DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

func DebugPrintRoute(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	logger.Log.WithFields(logrus.Fields{"httpMethod": httpMethod, "path": absolutePath, "handlerName": handlerName, "nuHandlers": nuHandlers}).Info("endpointRequest")
}
