package handlers

import (
	"ctf-challenge/internal/dto"
	"ctf-challenge/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterShortlinkPaths(r *gin.RouterGroup) {
	r.POST("/", createShortLink)
	r.GET("/:slug", getShortlinkBySlug)
}

func getShortlinkBySlug(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": services.GetFullUrlByShortlinkSlug(ctx.Param("slug")),
	})
}

func createShortLink(ctx *gin.Context) {
	var createRequest dto.CreateShortlinkRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body",
		})
		return
	}

	slug, err := services.CreateShortlink(createRequest.TargetUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"shortlink": slug,
	})
}
