package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/goofynugtz/kafka-producer-consumer/pkg/controllers"
)

// all routes start with "/"
func PublicRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/recieve", controllers.RecieveProduct())
}
