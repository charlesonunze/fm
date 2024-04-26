package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

// makeGetRequest sends a GET request and returns the status code, response headers, and the body.
func makeGetRequest(target string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	return e.NewContext(req, res), res
}

// makePostRequest sends a POST request and returns the status code, response header, and the body.
func makePostRequest(target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	return e.NewContext(req, res), res
}
