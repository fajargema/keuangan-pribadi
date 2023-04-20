package controllers

import (
	"bytes"
	"encoding/json"
	"keuangan-pribadi/config"
	"keuangan-pribadi/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testCaseCategory struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var categoryController CategoryController = InitCategoryController()

func InitCategoryEcho() *echo.Echo {
	config.InitDB()

	e := echo.New()

	return e
}

func TestGetAllCategories_Success(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "success",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, categoryController.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateCategory_Success(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "success",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	var categoryInput models.CategoryInput = models.CategoryInput{
		Name:      "test",
	}

	jsonBody, err := json.Marshal(&categoryInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, categoryController.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateCategory_Failed(t *testing.T) {
	testcase := testCaseCategory {
		name:                   "failed",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	categoryInput := models.CategoryInput{}

	jsonBody, _ := json.Marshal(&categoryInput)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/categories", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, categoryController.Create(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetCategoryByID_Success(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "success",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	category, err := config.SeedCategory()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	categoryID := strconv.Itoa(int(category.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(categoryID)

	if assert.NoError(t, categoryController.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetCategoryByID_Failed(t *testing.T) {
	testcase := testCaseCategory {
		name:                   "failed",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusNotFound,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/categories", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues("-1")

	if assert.NoError(t, categoryController.GetByID(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateCategoryByID_Success(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "success",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	category, _ := config.SeedCategory()

	categoryInput := models.CategoryInput{
		Name:      "updated",
	}

	jsonBody, _ := json.Marshal(&categoryInput)
	bodyReader := bytes.NewReader(jsonBody)

	categoryID := strconv.Itoa(int(category.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/categories", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)
	c.SetParamNames("id")
	c.SetParamValues(categoryID)

	if assert.NoError(t, categoryController.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateCategoryByID_Failed(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "failed",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	category, _ := config.SeedCategory()

	categoryInput := models.CategoryInput{}

	jsonBody, _ := json.Marshal(&categoryInput)
	bodyReader := bytes.NewReader(jsonBody)

	categoryID := strconv.Itoa(int(category.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/categories", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)
	c.SetParamNames("id")
	c.SetParamValues(categoryID)

	if assert.NoError(t, categoryController.Update(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteCategoryByID_Success(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "success",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	category, err := config.SeedCategory()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	categoryID := strconv.Itoa(int(category.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(categoryID)

	if assert.NoError(t, categoryController.Delete(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteCategoryByID_Failed(t *testing.T) {
	testcase := testCaseCategory{
		name:                   "failed",
		path:                   "/api/v1/categories",
		expectedStatus:         http.StatusNotFound,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitCategoryEcho()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/categories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)
	c.SetParamNames("id")
	c.SetParamValues("-1")

	if assert.NoError(t, categoryController.Delete(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}