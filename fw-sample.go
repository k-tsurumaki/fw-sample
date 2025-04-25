package fwsample

import (
    "errors"
    "net/http"
)

// Routerインターフェース
type Router interface {
    Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error
    Get(path string, handler func(http.ResponseWriter, *http.Request)) error
    Post(path string, handler func(http.ResponseWriter, *http.Request)) error
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// App構造体: アプリケーション全体を管理
type App struct {
    Router Router
}

// New関数: App構造体の新しいインスタンスを作成し、初期化
func New() *App {
    return &App{
        Router: &router{
            routingTable: make(map[string]map[string]func(http.ResponseWriter, *http.Request)),
        },
    }
}

// ServeHTTP: App構造体がhttp.Handlerインターフェースを満たす
func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.Router.ServeHTTP(w, r)
}

// Run: アプリケーションを指定したアドレスで起動
func (h *App) Run(addr string) error {
    return http.ListenAndServe(addr, h)
}

// router構造体: ルーティング処理を担当
type router struct {
    routingTable map[string]map[string]func(http.ResponseWriter, *http.Request)
}

// Add: ルーティングテーブルにハンドラを追加
func (r *router) Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error {
    if method != http.MethodGet && method != http.MethodPost {
        return errors.New("unsupported method")
    }
    return r.add(method, path, handler)
}

// Get: GETメソッドのハンドラを追加
func (r *router) Get(path string, handler func(http.ResponseWriter, *http.Request)) error {
    return r.Add(http.MethodGet, path, handler)
}

// Post: POSTメソッドのハンドラを追加
func (r *router) Post(path string, handler func(http.ResponseWriter, *http.Request)) error {
    return r.Add(http.MethodPost, path, handler)
}

// ServeHTTP: http.Handlerインターフェースを満たす
func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if handlers, ok := rt.routingTable[r.Method]; ok {
        if handler, ok := handlers[r.URL.Path]; ok {
            handler(w, r)
            return
        }
    }
    http.NotFound(w, r)
}

// add: ルーティングテーブルにハンドラを追加する内部関数
func (r *router) add(method, path string, handler func(http.ResponseWriter, *http.Request)) error {
    if r.routingTable[method] == nil {
        r.routingTable[method] = make(map[string]func(http.ResponseWriter, *http.Request))
    }
    if r.routingTable[method][path] != nil {
        return errors.New("handler already exists")
    }
    r.routingTable[method][path] = handler
    return nil
}