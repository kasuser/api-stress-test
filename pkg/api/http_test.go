package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"stresstest/pkg/api"
	"stresstest/pkg/model"
	"testing"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(api.HandleRequest()))
}
func adminRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(api.HandleAdminRequest()))
}

func TestHandleRequest(t *testing.T) {
	model.InitializeOrders()

	req, err := http.NewRequest("GET", "/request", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(rr.Body.String()) != 2 {
		t.Errorf("handler returned unexpected body length: got %v want %v",
			len(rr.Body.String()), 2)
	}
}

func TestHandleAdminRequest(t *testing.T) {
	model.InitializeOrders()

	req, err := http.NewRequest("GET", "/admin/request", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(adminRequestHandler)

	handler.ServeHTTP(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	re := regexp.MustCompile("([a-z]{2}-\\d\n){50}")
	b := rr.Body.String()

	if re.FindString(b) != b {
		t.Errorf("handler returned unexpected body structure: got %v want match regular exptession %v",
			rr.Body.String(), "([a-z]{2}-\\d\n){50}")
	}
}