package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	generate_rg := rg.Group("prescripts/generate")

	generate_rg.GET("", HandleGeneratePrescript)
}
