package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keuangan-pribadi/config"
	"keuangan-pribadi/middleware"
	"keuangan-pribadi/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testCaseDetailSaving struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var detailSavingController DetailSavingController = InitDetailSavingController()

func InitDetailSavingEcho() *echo.Echo {
	config.InitDB()

	e := echo.New()

	return e
}

func TestGetAllDetailSavings_Success(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "success",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, detailSavingController.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetAllDetailSavings_Failed(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "failed",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, detailSavingController.GetAll(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateDetailSaving_Success(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "success",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	user, err := config.SeedUser()
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	saving, _ := config.SeedSaving()

	var savingInput models.DetailSavingInput = models.DetailSavingInput{
		Value: 			1,
		SavingID: 			saving.ID,
		UserID: 			user.ID,
	}

	jsonBody, err := json.Marshal(&savingInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", tokenString)

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, detailSavingController.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateDetailSaving_Failed(t *testing.T) {
	testcase := testCaseDetailSaving {
		name:                   "failed",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	detailSavingInput := models.DetailSavingInput{}

	jsonBody, _ := json.Marshal(&detailSavingInput)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/detail-savings", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, detailSavingController.Create(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetDetailSavingByID_Success(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "success",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	detailSaving, err := config.SeedDetailSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	detailSavingID := strconv.Itoa(int(detailSaving.ID))

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(detailSavingID)

	if assert.NoError(t, detailSavingController.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetDetailSavingByID_Failed(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "failed",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	detailSaving, err := config.SeedDetailSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	detailSavingID := strconv.Itoa(int(detailSaving.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(detailSavingID)

	if assert.NoError(t, detailSavingController.GetByID(ctx)) {
		assert.Equal(t, http.StatusBadRequest, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateDetailSavingByID_Failed(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "failed",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)
	saving, _ := config.SeedSaving()
	detailSaving, _ := config.SeedDetailSaving()

	detailSavingInput := models.DetailSavingInput{
		Value: 			1,
		SavingID:  		saving.ID,
		UserID:  		user.ID,
	}

	jsonBody, _ := json.Marshal(&detailSavingInput)
	bodyReader := bytes.NewReader(jsonBody)

	detailSavingID := strconv.Itoa(int(detailSaving.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/detail-savings", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenString)

	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(detailSavingID)

	if assert.NoError(t, detailSavingController.Update(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateDetailSavingByID_TokenFailed(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "failed",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	user, _ := config.SeedUser()
	saving, _ := config.SeedSaving()
	detailSaving, _ := config.SeedDetailSaving()

	detailSavingInput := models.DetailSavingInput{
		Value: 			1,
		SavingID:  		saving.ID,
		UserID:  		user.ID,
	}

	jsonBody, _ := json.Marshal(&detailSavingInput)
	bodyReader := bytes.NewReader(jsonBody)

	detailSavingID := strconv.Itoa(int(detailSaving.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/detail-savings", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(detailSavingID)

	if assert.NoError(t, detailSavingController.Update(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteDetailSavingByID_Success(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "success",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	detailSaving, err := config.SeedDetailSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	detailSavingID := strconv.Itoa(int(detailSaving.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(detailSavingID)

	if assert.NoError(t, detailSavingController.Delete(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteDetailSavingByID_Failed(t *testing.T) {
	testcase := testCaseDetailSaving{
		name:                   "success",
		path:                   "/api/v1/detail-savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitDetailSavingEcho()

	detailSaving, err := config.SeedDetailSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	detailSavingID := strconv.Itoa(int(detailSaving.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(detailSavingID)

	if assert.NoError(t, detailSavingController.Delete(ctx)) {
		assert.Equal(t, http.StatusBadRequest, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}