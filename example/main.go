package main

import (
	"io"
	"log"
	"net/http"

	fwsample "github.com/k-tsurumaki/fw-sample"
	"github.com/k-tsurumaki/fw-sample/middleware"
)

func main() {
	// アプリケーションのインスタンスを初期化
	app := fwsample.New()

	stdlog := &fwsample.StdLogger{}

	logger := &middleware.LoggingMiddleware{Logger: stdlog}

	// ログを出力するミドルウェアをアプリケーション全体に適用
	app.Use(logger.Logging)

	// ハンドラを登録
	app.Router.Get("/hello", helloHandler)
	app.Router.Post("/echo", echoHandler)

	log.Println("Starting server on :8080")
	if err := app.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func helloHandler(ctx fwsample.Context) {
	ctx.Text(http.StatusOK, "Hello World!!")
}

func echoHandler(ctx fwsample.Context) {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		ctx.Text(http.StatusInternalServerError, "failed to read body")
		return
	}
	ctx.Text(http.StatusOK, string(body))
}
