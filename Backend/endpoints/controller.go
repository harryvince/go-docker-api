package endpoints

import (
	"github.com/gin-gonic/gin"
)

// Controller registers all endpoints
func Routes(route *gin.Engine) {
	// Register endpoints here
	router := route.Group("/api")
	router.POST("/postgres", CreatePostgres)
	router.POST("/mysql", CreateMySQL)
	router.POST("/ubuntu", CreateUbuntuInstance)
	router.GET("/system-stats", SystemStats)
	router.GET("/containers", GetContainersEndpoint)
	router.POST("/delete-container", DeleteContainer)
}
