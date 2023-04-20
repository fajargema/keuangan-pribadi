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

type testCaseSaving struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var savingController SavingController = InitSavingController()

func InitSavingEcho() *echo.Echo {
	config.InitDB()

	e := echo.New()

	return e
}

func TestGetAllSavings_Success(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "success",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, savingController.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetAllSavings_Failed(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "failed",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, savingController.GetAll(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateSaving_Success(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "success",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	user, err := config.SeedUser()
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	var savingInput models.SavingInput = models.SavingInput{
		Name:      		"test",
		Value: 			1,
		Goal: 			25000,
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

	if assert.NoError(t, savingController.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateSaving_Failed(t *testing.T) {
	testcase := testCaseSaving {
		name:                   "failed",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	savingInput := models.SavingInput{}

	jsonBody, _ := json.Marshal(&savingInput)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/savings", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, savingController.Create(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetSavingByID_Success(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "success",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	saving, err := config.SeedSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	savingID := strconv.Itoa(int(saving.ID))

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(savingID)

	if assert.NoError(t, savingController.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetSavingByID_Failed(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "failed",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	saving, err := config.SeedSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	savingID := strconv.Itoa(int(saving.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(savingID)

	if assert.NoError(t, savingController.GetByID(ctx)) {
		assert.Equal(t, http.StatusBadRequest, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateSavingByID_Failed(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "failed",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)
	saving, _ := config.SeedSaving()

	savingInput := models.SavingInput{
		Name:      		"testupdate",
		Value: 			1,
		Goal: 			25000,
		UserID:  		user.ID,
	}

	jsonBody, _ := json.Marshal(&savingInput)
	bodyReader := bytes.NewReader(jsonBody)

	savingID := strconv.Itoa(int(saving.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/savings", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenString)

	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(savingID)

	if assert.NoError(t, savingController.Update(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateSavingByID_TokenFailed(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "failed",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	user, _ := config.SeedUser()
	saving, _ := config.SeedSaving()

	savingInput := models.SavingUpdate{
		Name:      		"testupdate",
		Goal: 			1,
		UserID:  		user.ID,
	}

	jsonBody, _ := json.Marshal(&savingInput)
	bodyReader := bytes.NewReader(jsonBody)

	savingID := strconv.Itoa(int(saving.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/savings", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(savingID)

	if assert.NoError(t, savingController.Update(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteSavingByID_Success(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "success",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	saving, err := config.SeedSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	savingID := strconv.Itoa(int(saving.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(savingID)

	if assert.NoError(t, savingController.Delete(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteSavingByID_Failed(t *testing.T) {
	testcase := testCaseSaving{
		name:                   "success",
		path:                   "/api/v1/savings",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitSavingEcho()

	saving, err := config.SeedSaving()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	savingID := strconv.Itoa(int(saving.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(savingID)

	if assert.NoError(t, savingController.Delete(ctx)) {
		assert.Equal(t, http.StatusBadRequest, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}