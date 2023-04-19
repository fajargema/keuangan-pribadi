package controllers

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/services"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DetailSavingController struct {
	service services.DetailSavingService
}

func InitDetailSavingController() DetailSavingController {
	return DetailSavingController{
		service: services.InitDetailSavingService(),
	}
}

func (fc *DetailSavingController) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	detailSavings, err := fc.service.GetAll(token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch detailSavings data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.DetailSaving]{
		Status:  "success",
		Message: "all detail savings",
		Data:    detailSavings,
	})
}

func (fc *DetailSavingController) GetByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var detailSavingID string = c.Param("id")

	detailSaving, err := fc.service.GetByID(detailSavingID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "detail saving not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.DetailSaving]{
		Status:  "success",
		Message: "detail saving found",
		Data:    detailSaving,
	})
}

func (fc *DetailSavingController) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var detailSavingInput models.DetailSavingInput

	if err := c.Bind(&detailSavingInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(detailSavingInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	detailSaving, err := fc.service.Create(detailSavingInput, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.DetailSaving]{
		Status:  "success",
		Message: "detail saving created",
		Data:    detailSaving,
	})
}

func (fc *DetailSavingController) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var detailSavingID string = c.Param("id")

	var detailSavingInput models.DetailSavingInput

	if err := c.Bind(&detailSavingInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(detailSavingInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	detailSaving, err := fc.service.Update(detailSavingInput, detailSavingID, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.DetailSaving]{
		Status:  "success",
		Message: "detail saving updated",
		Data:    detailSaving,
	})
}

func (fc *DetailSavingController) Delete(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var detailSavingID string = c.Param("id")

	err := fc.service.Delete(detailSavingID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "Not Found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "detail saving deleted",
	})
}
