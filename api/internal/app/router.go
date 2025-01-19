package app

import (
	"ctf-challenge/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	systemRoutes := r.Group("/system")
	adminRoutes := r.Group("/admin")
	shortlinkRouters := r.Group("/s")

	handlers.RegisterHealthCheckPath(systemRoutes)
	adminRoutes.Use(addAdminHeaders())
	handlers.RegisterAdminPaths(adminRoutes)
	handlers.RegisterShortlinkPaths(shortlinkRouters)
	return r
}

func addAdminHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Admin-Token")

		if token != "test-token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Admin token required",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
