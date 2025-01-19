package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminPaths(r *gin.RouterGroup) {
	r.GET("/", adminGetHandler)
}

func adminGetHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "admin endpoint reached",
	})
}
