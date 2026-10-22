package main

import (
	//"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()	
	router.Use(gin.Recovery())
	api_group := router.Group("/api/v1")

	RegisterRoutes(api_group)

	combos := (len(timeConditions) * len(actions)) + (len(actions) * len(markers) * (len(actions) - 1)) + (len(actions) * len(markers) * (len(actions)-1) * len(postScripts)) + (len(timeConditions) * len(actions) * len(postScripts)) + len(singles)

	fmt.Printf("Prescript combos: %d\n", combos)

	router.Run()
}
