package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/camilocorreaUdeA/GoBootcampTechTest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m *MockObject) GetCustomData(c *gin.Context) (interface{}, error) {
	args := m.Called(c)
	return args.Get(0).(interface{}), args.Error(1)
}

func (m *MockObject) SayHello(c *gin.Context) (interface{}, error) {
	args := m.Called(c)
	return args.Get(0).(interface{}), args.Error(1)
}

var methodResponse models.ApiData = models.ApiData{
	Page:       2,
	PerPage:    6,
	Total:      12,
	TotalPages: 2,
	Data: []models.UserData{
		{
			ID:        7,
			Email:     "michael.lawson@reqres.in",
			FirstName: "Michael",
			LastName:  "Lawson",
			Avatar:    "https://reqres.in/img/faces/7-image.jpg",
		},
	},
	Support: models.SupportInfo{
		Url:  "https://reqres.in/#support-heading",
		Text: "To keep ReqRes free, contributions towards server costs are appreciated!",
	},
}

func TestGetCustomData(t *testing.T) {
	t.Run("GetCustomData success", func(t *testing.T) {
		mockClient := &MockObject{}
		gin.SetMode(gin.TestMode)
		mockClient.On("GetCustomData", mock.AnythingOfType("*gin.Context")).Return(methodResponse, nil)
		response := struct {
			Reponse models.ApiData `json:"response"`
			Status  string         `json:"status"`
		}{
			methodResponse,
			"OK",
		}

		expectedResponse, _ := json.Marshal(response)

		req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
		rec := httptest.NewRecorder()
		handler := gin.HandlerFunc(RequestWrapper(mockClient.GetCustomData))
		router := gin.Default()
		router.Use(handler)
		router.ServeHTTP(rec, req)

		assert.Equal(t, string(expectedResponse), rec.Body.String())
		mockClient.AssertExpectations(t)
	})
	t.Run("GetCustomData failure", func(t *testing.T) {
		mockClient := &MockObject{}
		gin.SetMode(gin.TestMode)
		err := errors.New("Error executing request")
		mockClient.On("GetCustomData", mock.AnythingOfType("*gin.Context")).Return(models.ApiData{}, err)
		response := struct {
			Error string `json:"error"`
		}{
			"Internal Error: Error executing request",
		}

		expectedResponse, _ := json.Marshal(response)

		req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
		rec := httptest.NewRecorder()
		handler := gin.HandlerFunc(RequestWrapper(mockClient.GetCustomData))
		router := gin.Default()
		router.Use(handler)
		router.ServeHTTP(rec, req)

		assert.Equal(t, string(expectedResponse), rec.Body.String())
		mockClient.AssertExpectations(t)
	})
}

func TestSayHello(t *testing.T) {
	mockClient := &MockObject{}
	gin.SetMode(gin.TestMode)
	methodResponse := "Hello World!"
	mockClient.On("SayHello", mock.AnythingOfType("*gin.Context")).Return(methodResponse, nil)

	response := struct {
		Reponse string `json:"response"`
		Status  string `json:"status"`
	}{
		"Hello World!",
		"OK",
	}

	expectedResponse, _ := json.Marshal(response)

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	handler := gin.HandlerFunc(RequestWrapper(mockClient.SayHello))
	router := gin.Default()
	router.Use(handler)
	router.ServeHTTP(rec, req)

	assert.Equal(t, string(expectedResponse), rec.Body.String())
	mockClient.AssertExpectations(t)
}
