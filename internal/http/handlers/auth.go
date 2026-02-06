package handlers

import (
	"net/http"
)

// Login handles authentication requests.
func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Refresh handles access token refresh requests.
func Refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
