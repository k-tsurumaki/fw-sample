package fwsample

import (
	"net/http"

	"github.com/k-tsurumaki/fw-sample/config"
	"github.com/k-tsurumaki/fw-sample/middleware"
	"github.com/k-tsurumaki/fw-sample/router"
)

// アプリケーション全体を管理する構造体
type App struct {
	Router     RouterInterface
	Middleware []middleware.Middleware
}

type RouterInterface interface {
	Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error
	Get(path string, handler func(http.ResponseWriter, *http.Request)) error
	Post(path string, handler func(http.ResponseWriter, *http.Request)) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (a *App) Use(m middleware.Middleware) {
	a.Middleware = append(a.Middleware, m)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.Router.ServeHTTP(w, r)
	})

	// ミドルウェアを逆順に適用
	for i := len(a.Middleware) - 1; i >= 0; i-- {
		handler = http.HandlerFunc(a.Middleware[i](handler).ServeHTTP)
	}

	handler.ServeHTTP(w, r)
}

func New() *App {
	return &App{
		Router: router.New(),
	}
}

func (a *App) Run() error {
	return a.RunWithConfig(config.DefaultConfig)
}

func (a *App) RunWithConfig(cfg config.Config) error {
	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      a,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	return server.ListenAndServe()
}
