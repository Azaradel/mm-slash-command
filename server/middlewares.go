package server

import (
	"net/http"
	"strings"
)

func (s *server) Authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Token ")

		if token != s.authToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}
