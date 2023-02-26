package handlers

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constants"
	"github.com/nuttchai/go-rest/internal/services"
	"github.com/nuttchai/go-rest/internal/utils/api"
)

type TUserHandler struct {
	userService services.IUserService
}

var (
	UserHandler IUserHandler
)

func InitUserHandler() IUserHandler {
	UserHandler = &TUserHandler{
		userService: services.InitUserService(),
	}
	return UserHandler
}

func (h *TUserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.userService.GetUser(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(user, constants.GetUserSuccessMsg)
	return c.JSON(res.Status, res)
}