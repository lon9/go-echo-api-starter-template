package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	controller := NewHealthController()

	if err := controller.Index(c); err != nil {
		t.Error(err)
		if rec.Code != http.StatusOK {
			t.Errorf("Status code should be 200:%d", rec.Code)
		}
		if rec.Body.String() != `{"status":200,"message":"OK","result":"OK"}` {
			t.Errorf("Body is wrong")
		}
	}
}
