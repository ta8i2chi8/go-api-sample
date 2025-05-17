package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// リクエストとレスポンスをログに出力するミドルウェア
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// リクエストのログを出力
		slog.Info("api request", "method", r.Method, "path", r.URL.Path, "query", r.URL.Query(), "body", r.Body)

		// レスポンスのヘッダーをキャプチャするためのレスポンスライターを作成
		rec := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rec, r)

		durationMs := time.Since(start).Milliseconds()

		// レスポンスのログを出力
		slog.Info("Response", "method", r.Method, "path", r.URL.Path, "status_code", rec.statusCode, "duration_ms", durationMs)
	})
}

// レスポンスのステータスコードをキャプチャするためのカスタムレスポンスライター
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}
