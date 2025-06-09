package main

import (
	"log"
	"net/http"

	fwsample "github.com/k-tsurumaki/fw-sample"
	"github.com/k-tsurumaki/fw-sample/middleware"
)

func main() {
	// アプリケーションのインスタンスを初期化
	app := fwsample.New()

	stdlog := &fwsample.StdLoggerWithRequestID{}
	logger := &middleware.LoggingMiddleware{Logger: stdlog}

	app.Use(middleware.RequestID)
	app.Use(logger.Logging)

	// ハンドラを登録
	app.Router.Get("/hello", helloHandler)
	app.Router.Post("/echo", echoHandler)

	log.Println("Starting server on :8080")
	if err := app.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	w.Write(body)
}
