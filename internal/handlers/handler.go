package handlers

import (
	"github.com/MalukiMuthusi/pulseid/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()

	gin.DebugPrintRouteFunc = DebugPrintRoute

	generate := Generate{}
	r.GET("/generate", generate.Handle)

	return r
}

func DebugPrintRoute(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	logger.Log.WithFields(logrus.Fields{"httpMethod": httpMethod, "path": absolutePath, "handlerName": handlerName, "nuHandlers": nuHandlers}).Info("endpointRequest")
}
