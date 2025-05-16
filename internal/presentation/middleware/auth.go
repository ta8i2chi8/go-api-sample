package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ta8i2chi8/go-api-sample/internal/config"
	"github.com/ta8i2chi8/go-api-sample/internal/presentation/common"
)

const (
	headerAuthorization   = "Authorization"
	headerWWWAuthenticate = "WWW-Authenticate"
)

type bearerToken struct {
	value string
}

func (t *bearerToken) validate() error {
	cfg, err := config.Get()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	if t.value != cfg.APIToken {
		return fmt.Errorf("invalid api token")
	}

	return nil
}

func CheckBearerToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token, err := extractTokenFromRequest(req)
		if err != nil {
			handleError(req.Context(), w, err, `Bearer realm="token_required"`)
			return
		}

		if err := token.validate(); err != nil {
			handleError(req.Context(), w, err, `Bearer realm="token_invalid"`)
			return
		}

		next.ServeHTTP(w, req)
	})
}

func extractTokenFromRequest(req *http.Request) (*bearerToken, error) {
	authHeader := req.Header.Get(headerAuthorization)
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return nil, fmt.Errorf("invalid token format")
	}

	return &bearerToken{value: token}, nil
}

func handleError(ctx context.Context, w http.ResponseWriter, err error, wwwAuthenticateHeader string) {
	w.Header().Set(headerWWWAuthenticate, wwwAuthenticateHeader)
	common.WriteErrorResponse(ctx, w, http.StatusUnauthorized, err.Error())
	slog.Error("auth check error", "err", err)
}
