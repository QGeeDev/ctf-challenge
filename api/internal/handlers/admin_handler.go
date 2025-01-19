package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) RegisterAdminPaths(r *gin.RouterGroup) {
	r.GET("/", h.adminGetHandler)
}

func (h *AdminHandler) adminGetHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "admin endpoint reached",
	})
}
