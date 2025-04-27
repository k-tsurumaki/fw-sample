package fwsample

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRouter struct {
	routingTable map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func (m *MockRouter) Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error {
	if m.routingTable[method] == nil {
		m.routingTable[method] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	m.routingTable[method][path] = handler
	return nil
}

func (m *MockRouter) Get(path string, handler func(http.ResponseWriter, *http.Request)) error {
	return m.Add(http.MethodGet, path, handler)
}

func (m *MockRouter) Post(path string, handler func(http.ResponseWriter, *http.Request)) error {
	return m.Add(http.MethodPost, path, handler)
}

func (m *MockRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handlers, ok := m.routingTable[r.Method]; ok {
		if handler, ok := handlers[r.URL.Path]; ok {
			handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func TestAppServeHTTP(t *testing.T) {
	mockRouter := &MockRouter{
		routingTable: make(map[string]map[string]func(http.ResponseWriter, *http.Request)),
	}
	app := &App{Router: mockRouter}

	// Getリクエストが来たらstatusOKを返すハンドラを追加
	mockRouter.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET /test"))
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	// Getリクエストを投げるとstatusOKが返るかテスト
	app.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
	if rr.Body.String() != "GET /test" {
		t.Errorf("Expected body 'GET /test', got '%s'", rr.Body.String())
	}
	
	// Postリクエストが来たらstatusOKを返すハンドラを追加
	mockRouter.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("POST /test"))
	})
	
	// Postリクエストを投げるとstatusOKが返るかテスト
	req, _ = http.NewRequest("POST", "/test", nil)
	rr = httptest.NewRecorder()
	app.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
	if rr.Body.String() != "POST /test" {
		t.Errorf("Expected body 'POST /test', got '%s'", rr.Body.String())
	}
}
