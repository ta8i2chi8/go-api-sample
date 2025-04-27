package logger

import (
	"context"
	"log/slog"
)

// 以下ブログを参考
// https://qiita.com/Imamotty/items/3fbe8ce6da4f1a653fae#context%E3%81%8B%E3%82%89%E5%8F%96%E3%82%8A%E5%87%BA%E3%81%97%E3%81%9F%E5%80%A4%E3%82%92%E3%83%AD%E3%82%AE%E3%83%B3%E3%82%B0

type key struct{}

var TraceIDKey = key{}

type TraceHandler struct {
	slog.Handler
}

func NewTraceHandler(h slog.Handler) *TraceHandler {
	return &TraceHandler{h}
}

func (h *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(slog.String("traceID", ctx.Value(TraceIDKey).(string)))
	return h.Handler.Handle(ctx, r)
}
