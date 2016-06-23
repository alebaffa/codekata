package tennisweb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCallAPI(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/score", getScore).Methods("GET")
	req, err := http.NewRequest("GET", "/score", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP status expected: 200 got: %d", w.Code)
	}
}
