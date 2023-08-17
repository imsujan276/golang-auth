package apiHandlers

import (
	"fmt"
	"net/http"
	"pomo/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// get user by UUID and return user
func (h *Handler) GetUserByUUID(ctx *gin.Context) {
	uuidString := ctx.Param("uuid")
	fmt.Println(uuidString)
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		utils.APIResponse(ctx, "Invalid UUID", http.StatusBadRequest, nil)
		return
	}
	user, status := h.service.GetUserByUUID(uuid)
	if status != http.StatusOK {
		utils.APIResponse(ctx, "User not found", status, nil)
		return
	}
	utils.APIResponse(ctx, "User Found", http.StatusOK, user)
}

func (h *Handler) GetMe(ctx *gin.Context) {
	user, status := h.service.GetMe(ctx)
	if status != http.StatusOK {
		utils.APIResponse(ctx, "User not found", status, nil)
		return
	}
	utils.APIResponse(ctx, "User Found", http.StatusOK, user)
}
