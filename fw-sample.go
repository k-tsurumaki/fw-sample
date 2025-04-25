package fwsample

import (
	"errors"
	"net/http"
)

// アプリケーション全体を管理する構造体
type App struct {
	Router *Router
}

// ルーティング処理を担当する構造体
type Router struct {
	routingTable map[string]map[string]func(http.ResponseWriter, *http.Request)
}

// ルーティングテーブルにハンドラを追加する関数
func (r *Router) add(method, path string, handler func(http.ResponseWriter, *http.Request)) error {
	if r.routingTable[method] == nil {
		r.routingTable[method] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	if r.routingTable[method][path] != nil {
		return errors.New("handler already exists")
	}
	r.routingTable[method][path] = handler
	return nil
}

func (r *Router) Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error {
	if method != http.MethodGet && method != http.MethodPost {
		return errors.New("unsupported method")
	}
	return r.add(method, path, handler)
}


// GETメソッドのハンドラを追加する関数
func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) error {
	return r.Add(http.MethodGet, path, handler)
}

// POSTメソッドのハンドラを追加する関数
func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request)) error {
	return r.Add(http.MethodPost, path, handler)
}

// ServeHTTPメソッドを実装することで、http.Handlerインターフェースを満たすようになる
// これにより、http.ListenAndServeに渡すことができるようになる
func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if handlers, ok := h.Router.routingTable[r.Method]; ok {
            if handler, ok := handlers[r.URL.Path]; ok {
                handler(w, r)
                return
            }
        }
		// ルーティングにマッチしなかった場合は404エラーを返す
		http.NotFound(w, r)
	})
	handler.ServeHTTP(w, r)
}

// App構造体の新しいインスタンスを作成し、初期化する関数
func New() *App {
	return &App{
		Router: &Router{
			routingTable: make(map[string]map[string]func(http.ResponseWriter, *http.Request)),
		},
	}
}

// アプリケーションを指定したアドレスで起動する関数
func (e *App) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
