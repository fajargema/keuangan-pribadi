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

type testCaseFinance struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var financeController FinanceController = InitFinanceController()

func InitFinanceEcho() *echo.Echo {
	config.InitDB()

	e := echo.New()

	return e
}

func TestGetAllFinances_Success(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetAllFinances_Failed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "failed",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.GetAll(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateFinance_Success(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, err := config.SeedUser()
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	category, err := config.SeedCategory()
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var financeInput models.FinanceInput = models.FinanceInput{
		Name:      		"test",
		Type: 			1,
		Money: 			25000,
		UserID:  		user.ID,
		CategoryID:  	category.ID,
	}

	jsonBody, err := json.Marshal(&financeInput)

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

	if assert.NoError(t, financeController.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateFinance_Failed(t *testing.T) {
	testcase := testCaseFinance {
		name:                   "failed",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	financeInput := models.FinanceInput{}

	jsonBody, _ := json.Marshal(&financeInput)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/finances", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.Create(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetFinanceByID_Success(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	finance, err := config.SeedFinance()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	financeID := strconv.Itoa(int(finance.ID))

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(financeID)

	if assert.NoError(t, financeController.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetFinanceByID_Failed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "failed",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	finance, err := config.SeedFinance()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	financeID := strconv.Itoa(int(finance.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(financeID)

	if assert.NoError(t, financeController.GetByID(ctx)) {
		assert.Equal(t, http.StatusBadRequest, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestSearch_Success(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances/search",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", tokenString)
	q := req.URL.Query()
	q.Add("from", "2022-01-01")
	q.Add("to", "2022-01-31")
	req.URL.RawQuery = q.Encode()
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.Search(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestSearch_TokenFailed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "failed",
		path:                   "/api/v1/finances/search",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", "")
	q := req.URL.Query()
	q.Add("from", "2022-01-01")
	q.Add("to", "2022-01-31")
	req.URL.RawQuery = q.Encode()
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.Search(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestSearch_FromFailed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances/search",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", tokenString)
	q := req.URL.Query()
	q.Add("from", "")
	q.Add("to", "2022-01-31")
	req.URL.RawQuery = q.Encode()
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.Search(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestSearch_ToFailed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances/search",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	req.Header.Add("Authorization", tokenString)
	q := req.URL.Query()
	q.Add("from", "2022-01-01")
	q.Add("to", "")
	req.URL.RawQuery = q.Encode()
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, financeController.Search(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateFinanceByID_Failed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "failed",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)
	category, _ := config.SeedCategory()
	finance, _ := config.SeedFinance()

	financeInput := models.FinanceInput{
		Name:      		"testupdate",
		Type: 			1,
		Money: 			25000,
		UserID:  		user.ID,
		CategoryID:  	category.ID,
	}

	jsonBody, _ := json.Marshal(&financeInput)
	bodyReader := bytes.NewReader(jsonBody)

	financeID := strconv.Itoa(int(finance.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/finances", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenString)

	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(financeID)

	if assert.NoError(t, financeController.Update(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateFinanceByID_TokenFailed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "failed",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	user, _ := config.SeedUser()
	category, _ := config.SeedCategory()
	finance, _ := config.SeedFinance()

	financeInput := models.FinanceInput{
		Name:      		"testupdate",
		Type: 			1,
		Money: 			25000,
		UserID:  		user.ID,
		CategoryID:  	category.ID,
	}

	jsonBody, _ := json.Marshal(&financeInput)
	bodyReader := bytes.NewReader(jsonBody)

	financeID := strconv.Itoa(int(finance.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/finances", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "")

	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues(financeID)

	if assert.NoError(t, financeController.Update(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteFinanceByID_Success(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	finance, err := config.SeedFinance()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	user, _ := config.SeedUser()
	token, _ := middleware.CreateToken(user.ID, user.Name)
	tokenString := fmt.Sprintf("Bearer %s", token)

	financeID := strconv.Itoa(int(finance.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", tokenString)
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(financeID)

	if assert.NoError(t, financeController.Delete(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteFinanceByID_Failed(t *testing.T) {
	testcase := testCaseFinance{
		name:                   "success",
		path:                   "/api/v1/finances",
		expectedStatus:         http.StatusBadRequest,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitFinanceEcho()

	finance, err := config.SeedFinance()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	financeID := strconv.Itoa(int(finance.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)
	request.Header.Add("Authorization", "")
	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(financeID)

	if assert.NoError(t, financeController.Delete(ctx)) {
		assert.Equal(t, http.StatusBadRequest, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}