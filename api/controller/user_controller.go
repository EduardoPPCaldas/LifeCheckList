package controller

import (
	"net/http"

	"github.com/EduardoPPCaldas/LifeCheckList/api/request"
	"github.com/EduardoPPCaldas/LifeCheckList/application/dto"
	"github.com/EduardoPPCaldas/LifeCheckList/application/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	uc usecase.UserUCHandler
}

func NewUserController(uc usecase.UserUCHandler) *UserController {
	return &UserController{
		uc: uc,
	}
}

func (controller UserController) SignUp(ctx echo.Context) error {
	var request request.CreateUserRequest

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := controller.uc.SignUp(dto.CreateUserDto{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusCreated, user)
}
