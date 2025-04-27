package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func New(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// http.ErrServerClosed は
		// http.Server.Shutdown() が正常に終了したことを示すので異常ではない。
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to serve", "err", err)
			return err
		}
		return nil
	})

	// シグナル受信やサーバの起動失敗を待つ
	<-ctx.Done()

	slog.Info("shutting down server...")
	if err := s.srv.Shutdown(context.Background()); err != nil {
		slog.Error("failed to shutdown", "err", err)
	}

	// グレースフルシャットダウンの終了を待つ
	return eg.Wait()
}
