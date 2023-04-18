package controllers

import (
	"keuangan-pribadi/models"
	"keuangan-pribadi/services"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type FinanceController struct {
	service services.FinanceService
}

func InitFinanceController() FinanceController {
	return FinanceController{
		service: services.InitFinanceService(),
	}
}

func (fc *FinanceController) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	finances, err := fc.service.GetAll(token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch finances data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Finance]{
		Status:  "success",
		Message: "all finances",
		Data:    finances,
	})
}

func (fc *FinanceController) GetByID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var financeID string = c.Param("id")

	finance, err := fc.service.GetByID(financeID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "finance not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Finance]{
		Status:  "success",
		Message: "finance found",
		Data:    finance,
	})
}

func (fc *FinanceController) Search(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	from, err := time.Parse(time.DateOnly, c.QueryParam("from"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid from date",
		})
	}
	
	to, err := time.Parse(time.DateOnly, c.QueryParam("to"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid to date",
		})
	}

	finances, err := fc.service.Search(from, to, token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch finances data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Finance]{
		Status:  "success",
		Message: "all finances",
		Data:    finances,
	})
}

func (fc *FinanceController) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var financeInput models.FinanceInput

	if err := c.Bind(&financeInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(financeInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	finance, err := fc.service.Create(financeInput, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Finance]{
		Status:  "success",
		Message: "finance created",
		Data:    finance,
	})
}

func (fc *FinanceController) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var financeID string = c.Param("id")

	var financeInput models.FinanceInput

	if err := c.Bind(&financeInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	validate := validator.New()
    if err := validate.Struct(financeInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
    }

	finance, err := fc.service.Update(financeInput, financeID, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Finance]{
		Status:  "success",
		Message: "finance updated",
		Data:    finance,
	})
}

func (fc *FinanceController) Delete(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
    if token == "" {
        return c.JSON(http.StatusBadRequest, models.Response[string]{
            Status:  "failed",
            Message: "Missing token in request header",
        })
    }
	token = strings.ReplaceAll(token, "Bearer ", "")

	var financeID string = c.Param("id")

	err := fc.service.Delete(financeID, token)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "Not Found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "finance deleted",
	})
}
