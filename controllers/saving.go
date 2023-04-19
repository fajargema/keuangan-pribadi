package controllers

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/services"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type SavingController struct {
	service services.SavingService
}

func InitSavingController() SavingController {
	return SavingController{
		service: services.InitSavingService(),
	}
}

func (fc *SavingController) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	savings, err := fc.service.GetAll(token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch savings data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Saving]{
		Status:  "success",
		Message: "all savings",
		Data:    savings,
	})
}

func (fc *SavingController) GetByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var savingID string = c.Param("id")

	saving, err := fc.service.GetByID(savingID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "saving not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Saving]{
		Status:  "success",
		Message: "saving found",
		Data:    saving,
	})
}

func (fc *SavingController) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var savingInput models.SavingInput

	if err := c.Bind(&savingInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(savingInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	saving, err := fc.service.Create(savingInput, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Saving]{
		Status:  "success",
		Message: "saving created",
		Data:    saving,
	})
}

func (fc *SavingController) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var savingID string = c.Param("id")

	var savingUpdate models.SavingUpdate

	if err := c.Bind(&savingUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(savingUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	saving, err := fc.service.Update(savingUpdate, savingID, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Saving]{
		Status:  "success",
		Message: "saving updated",
		Data:    saving,
	})
}

func (fc *SavingController) Delete(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var savingID string = c.Param("id")

	err := fc.service.Delete(savingID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "Not Found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "saving deleted",
	})
}
