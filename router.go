package fwsample

import (
	"errors"
	"net/http"
)

// ルーティング処理を担当する構造体
type Router struct {
	routingTable map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routingTable: make(map[string]map[string]HandlerFunc),
	}
}

// ルーティングテーブルにハンドラを追加する関数
func (r *Router) add(method, path string, handler HandlerFunc) error {
	if r.routingTable[method] == nil {
		r.routingTable[method] = make(map[string]HandlerFunc)
	}
	if r.routingTable[method][path] != nil {
		return errors.New("handler already exists")
	}
	r.routingTable[method][path] = handler
	return nil
}

func (r *Router) Add(method, path string, handler HandlerFunc) error {
	if method != http.MethodGet && method != http.MethodPost {
		return errors.New("unsupported method")
	}
	return r.add(method, path, handler)
}

// GETメソッドのハンドラを追加する関数
func (r *Router) Get(path string, handler HandlerFunc) error {
	return r.Add(http.MethodGet, path, handler)
}

// POSTメソッドのハンドラを追加する関数
func (r *Router) Post(path string, handler HandlerFunc) error {
	return r.Add(http.MethodPost, path, handler)
}

// routerにServeHTTPメソッドを実装
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routingTable[req.Method]; ok {
		if handler, ok := handlers[req.URL.Path]; ok {
			ctx := &context{
				req: req,
				rw: w,
				params: map[string]string{},
			}
			handler(ctx)
			return
		}
	}
}
