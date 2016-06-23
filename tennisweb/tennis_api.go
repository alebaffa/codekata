package tennisweb

import "net/http"

func getScore(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
