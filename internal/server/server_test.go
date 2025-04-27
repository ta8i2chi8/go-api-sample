package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	// Arrange
	// テスト用のリスナーを作成
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}

	// コンテキストとエラ―グループを作成
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	// テスト用のHTTPハンドラーを定義
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	// エラーグループを使ってサーバーを起動
	eg.Go(func() error {
		s := New(l, mux)
		return s.Run(ctx)
	})

	// Act
	// テスト用のリクエストを送信
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()

	// レスポンスボディを読み取る
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// Assert
	// 戻り値を検証する
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// サーバの終了動作を検証する
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
