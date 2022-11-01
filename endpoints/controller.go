package endpoints

import (
	"github.com/gin-gonic/gin"
)

// Controller registers all endpoints
func Routes(route *gin.Engine) {
	// Register endpoints here
	router := route.Group("/")
	router.POST("/postgres", CreatePostgres)
	router.POST("/mysql", CreateMySQL)
	router.GET("/system-stats", SystemStats)
	router.GET("/containers", GetContainersEndpoint)
}
