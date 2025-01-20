package handlers

import (
	"ctf-challenge/internal/dto"
	"ctf-challenge/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(service *services.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: service,
	}
}

func (h *AdminHandler) RegisterAdminPaths(r *gin.RouterGroup) {
	r.GET("/", h.adminGetHandler)
	r.POST("/customlink", h.CreateCustomShortlink)
}

func (h *AdminHandler) adminGetHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "admin endpoint reached",
	})
}

func (h *AdminHandler) CreateCustomShortlink(ctx *gin.Context) {
	var createRequest dto.CreateCustomShortlinkRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse request body",
		})
		return
	}

	result, err := h.adminService.AddCustomShortlink(createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":         "Internal server error",
			"error_details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
