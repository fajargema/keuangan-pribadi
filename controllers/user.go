package controllers

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/services"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController() UserController {
	return UserController{
		service: services.InitUserService(),
	}
}

func (uc *UserController) GetByEmail(c echo.Context) error {
	var userEmail string = c.Param("email")

	user, err := uc.service.GetByEmail(userEmail)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.User]{
		Status:  "success",
		Message: "user found",
		Data:    user,
	})
}

func (uc *UserController) Register(c echo.Context) error {
	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	user, err := uc.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.User]{
		Status:  "success",
		Message: "user created",
		Data:    user,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var userInput models.UserAuth

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	userResponse, err := uc.service.Login(userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response[string]{
			Status:  "failed",
			Message: "authentication failed",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.UserResponse]{
		Status:  "success",
		Message: "authenticated",
		Data:    userResponse,
	})
}

func (uc *UserController) UpdateMe(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	user, err := uc.service.UpdateMe(userInput, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.User]{
		Status:  "success",
		Message: "user updated",
		Data:    user,
	})
}