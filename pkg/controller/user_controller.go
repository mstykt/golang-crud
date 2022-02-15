package controller

import (
	"github.com/labstack/echo/v4"
	. "golang-crud/pkg/model/request"
	"golang-crud/pkg/service"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userController UserController) CreateUser(ctx echo.Context) error {
	var req UserRequest
	db := new(echo.DefaultBinder)
	err := db.Bind(&req, ctx)
	if err != nil {
		return err
	}
	user, createUserErr := userController.userService.CreateUser(req)
	if createUserErr != nil {
		return ctx.JSON(http.StatusBadRequest, createUserErr.Error())
	}
	return ctx.JSON(http.StatusCreated, user)
}

func (userController UserController) GetUser(ctx echo.Context) error {
	userId := ctx.Param("id")
	id, _ := strconv.ParseUint(userId, 0, 64)
	user, err := userController.userService.GetUser(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (userController UserController) GetAllUser(ctx echo.Context) error {
	users, err := userController.userService.GetAllUser()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}

func (userController UserController) UpdateUser(ctx echo.Context) error {
	userId := ctx.Param("id")
	id, _ := strconv.ParseUint(userId, 0, 64)

	var req UserRequest
	db := new(echo.DefaultBinder)
	err := db.Bind(&req, ctx)
	if err != nil {
		return err
	}

	updatedUser, updatedErr := userController.userService.UpdateUser(id, req)
	if updatedErr != nil {
		return ctx.JSON(http.StatusBadRequest, updatedErr.Error())
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

func (userController UserController) DeleteUser(ctx echo.Context) error {
	userId := ctx.Param("id")
	id, _ := strconv.ParseUint(userId, 0, 64)
	err := userController.userService.DeleteUser(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
