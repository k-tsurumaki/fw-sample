package fwsample

import (
	"net/http"

	"github.com/k-tsurumaki/fw-sample/middleware"
	"github.com/k-tsurumaki/fw-sample/router"
)

// アプリケーション全体を管理する構造体
type App struct {
	Router     RouterInterface
	Middleware []middleware.Middleware
}

// Routerをインターフェース化
type RouterInterface interface {
	Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error
	Get(path string, handler func(http.ResponseWriter, *http.Request)) error
	Post(path string, handler func(http.ResponseWriter, *http.Request)) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (a *App) Use(m middleware.Middleware) {
	a.Middleware = append(a.Middleware, m)
}

// ServeHTTPメソッドを実装することで、http.Handlerインターフェースを満たすようになる
// これにより、http.ListenAndServeに渡すことができるようになる
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ミドルウェアを適用
	// 最終的にRouterのServeHTTPを呼び出すハンドラ
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.Router.ServeHTTP(w, r)
	})

	// ミドルウェアを逆順に適用
	for i := len(a.Middleware) - 1; i >= 0; i-- {
		handler = http.HandlerFunc(a.Middleware[i](handler).ServeHTTP)
	}

	// 最終的なハンドラを実行
	handler.ServeHTTP(w, r)
}

// App構造体の新しいインスタンスを作成し、初期化する関数
func New() *App {
	return &App{
		Router: router.New(),
	}
}

// アプリケーションを指定したアドレスで起動する関数
func (a *App) Run(addr string) error {
	return http.ListenAndServe(addr, a)
}
