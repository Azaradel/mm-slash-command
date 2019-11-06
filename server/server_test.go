package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	token := "testingauthorization"
	server := NewServer(token)

	t.Run("should return status 200 if token is correct", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Add("Authorization", "Token "+token)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, http.StatusOK, res.Code)
	})

	t.Run("should return status 401 if token is incorrect", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Add("Authorization", "Token")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, http.StatusUnauthorized, res.Code)
	})

	t.Run("should set Content-Type to application/json", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Add("Authorization", "Token "+token)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertHeader(t, "application/json", res.Header().Get("Content-Type"))
	})
}

func assertStatusCode(t *testing.T, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("status code want: %v, got: %v", want, got)
	}
}

func assertHeader(t *testing.T, want, got string) {
	t.Helper()

	if got != want {
		t.Errorf("want value of header: %s, got: %s", want, got)
	}
}
