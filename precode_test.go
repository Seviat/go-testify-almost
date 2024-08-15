package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {

	expectedCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	resultCount := len(list)

	assert.Equal(t, resultCount, expectedCount)
}

func TestMainHandlerWhenOk(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedStatus := http.StatusOK
	resultStatus := responseRecorder.Code
	body := responseRecorder.Body.String()

	assert.NotEmpty(t, body)
	assert.Equal(t, resultStatus, expectedStatus)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=voronezh", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedStatus := http.StatusBadRequest
	resultStatus := responseRecorder.Code

	if assert.Equal(t, resultStatus, expectedStatus) {
		resultBody := responseRecorder.Body.String()
		expectedBody := "wrong city value"
		assert.Equal(t, expectedBody, resultBody)
	}
}
