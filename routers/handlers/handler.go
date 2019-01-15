package handlers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/client")
	{
		api.GET("/test")
	}

	return r
}
