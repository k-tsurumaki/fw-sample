package fwsample

import (
	"net/http"
)

// アプリケーション全体を管理する構造体
type App struct {
	Router     RouterInterface
	Middleware []MiddlewareFunc
}

type RouterInterface interface {
	Add(method, path string, handler HandlerFunc) error
	Get(path string, handler HandlerFunc) error
	Post(path string, handler HandlerFunc) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerFunc func(Context)

type MiddlewareFunc func(HandlerFunc) HandlerFunc

func (a *App) Use(m MiddlewareFunc) {
	a.Middleware = append(a.Middleware, m)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &context{
		req:    r,
		rw:     w,
		params: map[string]string{},
	}

	handler := func(c Context) {
		a.Router.ServeHTTP(c.ResponseWriter(), c.Request())
	}

	// ミドルウェアを逆順に適用
	for i := len(a.Middleware) - 1; i >= 0; i-- {
		handler = a.Middleware[i](handler)
	}

	handler(ctx)
}

func New() *App {
	return &App{
		Router: NewRouter(),
	}
}

func (a *App) Run() error {
	return a.RunWithConfig(DefaultConfig)
}

func (a *App) RunWithConfig(cfg Config) error {
	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      a,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	return server.ListenAndServe()
}
