package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ta8i2chi8/go-api-sample/internal/config"
)

var (
	testEnvMap = map[string]string{
		"ENV":       "local",
		"PORT":      "8070",
		"API_TOKEN": "test_valid_token",
	}
)

func setupEnv(t *testing.T) {
	t.Helper()
	for k, v := range testEnvMap {
		t.Setenv(k, v)
	}
}

func TestCheckBearerToken(t *testing.T) {
	validToken := "test_valid_token"
	setupEnv(t)
	if _, err := config.Load(context.Background()); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                 string
		token                string
		expectedStatus       int
		expectedAuthenticate string
	}{
		{
			name:                 "Success",
			token:                fmt.Sprintf("Bearer %s", validToken),
			expectedStatus:       http.StatusOK,
			expectedAuthenticate: "",
		},
		{
			name:                 "missing token",
			token:                "",
			expectedStatus:       http.StatusUnauthorized,
			expectedAuthenticate: `Bearer realm="token_required"`,
		},
		{
			name:                 "invalid token format",
			token:                "Bearer test_invalid_token",
			expectedStatus:       http.StatusUnauthorized,
			expectedAuthenticate: `Bearer realm="token_invalid"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/hello", nil)
			req.Header.Set(headerAuthorization, tt.token)

			rr := httptest.NewRecorder()

			handler := CheckBearerToken(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if authenticate := rr.Header().Get(headerWWWAuthenticate); authenticate != tt.expectedAuthenticate {
				t.Errorf("handler returned wrong WWW-Authenticate header: got %v want %v", authenticate, tt.expectedAuthenticate)
			}
		})
	}
}
