package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPage_ValidationInput(t *testing.T) {
	body := `{"name": "test page"}`
	req := httptest.NewRequest(http.MethodPost, "/page", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	Get(rec, req)
	fmt.Println(rec.Body.String())
}
