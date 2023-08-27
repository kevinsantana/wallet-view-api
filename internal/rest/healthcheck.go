package handlers

import (
	"net/http"
)

func Readiness(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Liveness(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
