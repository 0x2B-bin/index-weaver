package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGeneratePrescript(ctx *gin.Context) {
	seed := ctx.Query("seed")
	prescript := GeneratePrescript(seed);
	ctx.JSON(http.StatusOK, prescript)
}
