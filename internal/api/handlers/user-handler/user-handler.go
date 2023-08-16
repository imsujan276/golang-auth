package userhandler

import (
	"fmt"
	"net/http"
	usercontroller "pomo/internal/api/controllers/user-controller"
	"pomo/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	userService usercontroller.Service
}

func NewUserHandler(userService usercontroller.Service) *handler {
	return &handler{
		userService: userService,
	}
}

// get user by UUID and return user
func (h *handler) GetUserByUUID(ctx *gin.Context) {
	uuidString := ctx.Param("uuid")
	fmt.Println(uuidString)
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		utils.APIResponse(ctx, "Invalid UUID", http.StatusBadRequest, nil)
		return
	}
	user, status := h.userService.GetUserByUUID(uuid)
	if status != http.StatusOK {
		utils.APIResponse(ctx, "User not found", status, nil)
		return
	}
	utils.APIResponse(ctx, "User Found", http.StatusOK, user)
}

func (h *handler) GetMe(ctx *gin.Context) {
	user, status := h.userService.GetMe(ctx)
	if status != http.StatusOK {
		utils.APIResponse(ctx, "User not found", status, nil)
		return
	}
	utils.APIResponse(ctx, "User Found", http.StatusOK, user)
}
