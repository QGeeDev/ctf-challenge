package handlers

import "github.com/gin-gonic/gin"

func RegisterHealthCheckPath(r *gin.RouterGroup) {
	r.GET("/_health", healthcheckHandler)
}

func healthcheckHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "Healthy",
	})
}
