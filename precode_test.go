package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=someTown", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expected := `wrong city value`
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expected, responseRecorder.Body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, body)

	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
}
