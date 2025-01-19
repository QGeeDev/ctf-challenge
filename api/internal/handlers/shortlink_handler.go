package handlers

import (
	"ctf-challenge/internal/dto"
	"ctf-challenge/internal/services"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const QR_IMAGE_DIRECTORY = "/app/data/qr_codes/"

type ShortlinkHandler struct {
	service *services.ShortlinkService
}

func NewShortlinkHandler(service *services.ShortlinkService) *ShortlinkHandler {
	return &ShortlinkHandler{
		service: service,
	}
}

func (h *ShortlinkHandler) RegisterShortlinkPaths(r *gin.RouterGroup) {
	r.POST("/", h.createShortLink)
	r.GET("/:slug", h.getShortlinkBySlug)
	r.GET("/qr", h.GetQRImageFile)
}

func (h *ShortlinkHandler) getShortlinkBySlug(ctx *gin.Context) {
	response, err := h.service.GetFullUrlByShortlinkSlug(ctx.Param("slug"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if response == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Shortlink not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}

func (h *ShortlinkHandler) createShortLink(ctx *gin.Context) {
	var createRequest dto.CreateShortlinkRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body",
		})
		return
	}

	slug, err := h.service.CreateShortlink(createRequest.TargetUrl)
	if err != nil {
		zap.L().Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"shortlink": slug,
	})
}

func (h *ShortlinkHandler) GetQRImageFile(ctx *gin.Context) {
	fileName := ctx.Query("file")
	filePath := QR_IMAGE_DIRECTORY + fileName
	imageData, err := os.ReadFile(filePath)
	if err != nil || len(imageData) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Internal server error",
			"filepath": filePath,
			"err":      err.Error(),
			"len":      len(imageData),
		})
		return
	}

	ctx.Header("Content-Type", "image/png")
	ctx.Header("Content-Length", fmt.Sprintf("%d", len(imageData)))
	ctx.Data(http.StatusOK, "image/png", imageData)
}
