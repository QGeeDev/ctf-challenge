package app

import (
	"ctf-challenge/internal/handlers"
	"ctf-challenge/internal/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(dbService *services.DbService) *gin.Engine {
	r := gin.Default()

	shortlinkService := services.NewShortlinkService(dbService.DbConn)
	adminService := services.NewAdminService(dbService.DbConn)

	adminHandler := handlers.NewAdminHandler(adminService)
	shortlinkHandler := handlers.NewShortlinkHandler(shortlinkService)

	systemRoutes := r.Group("/system")
	adminRoutes := r.Group("/admin")
	shortlinkRouters := r.Group("/s")

	handlers.RegisterHealthCheckPath(systemRoutes)

	adminRoutes.Use(addAdminHeaders())
	adminHandler.RegisterAdminPaths(adminRoutes)

	shortlinkHandler.RegisterShortlinkPaths(shortlinkRouters)
	return r
}

func addAdminHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("ADMIN_TOKEN")
		userToken := ctx.GetHeader("X-Admin-Token")

		if userToken != token {
			if userToken == "" {
				userToken = "---"
			}
			ctx.Header("X-Admin-Token", userToken)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Admin token required",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
