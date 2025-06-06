package fwsample_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

type MockRouter struct {
	routingTable map[string]map[string]fwsample.HandlerFunc
}

func (m *MockRouter) Add(method, path string, handler fwsample.HandlerFunc) error {
	if m.routingTable[method] == nil {
		m.routingTable[method] = make(map[string]fwsample.HandlerFunc)
	}
	m.routingTable[method][path] = handler
	return nil
}

func (m *MockRouter) Get(path string, handler fwsample.HandlerFunc) error {
	return m.Add(http.MethodGet, path, handler)
}

func (m *MockRouter) Post(path string, handler fwsample.HandlerFunc) error {
	return m.Add(http.MethodPost, path, handler)
}

func (m *MockRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := fwsample.NewContext(w, r)
	if handlers, ok := m.routingTable[r.Method]; ok {
		if handler, ok := handlers[r.URL.Path]; ok {
			handler(ctx)
			return
		}
	}
	http.NotFound(w, r)
}

func TestAppServeHTTP(t *testing.T) {
	mockRouter := &MockRouter{
		routingTable: make(map[string]map[string]fwsample.HandlerFunc),
	}
	app := &fwsample.App{Router: mockRouter}

	// Get
	mockRouter.Get("/test", func(ctx fwsample.Context) {
		ctx.Text(http.StatusOK, "GET /test")
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	app.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
	if rr.Body.String() != "GET /test" {
		t.Errorf("Expected body 'GET /test', got '%s'", rr.Body.String())
	}
	
	// Post
	mockRouter.Post("/test", func(ctx fwsample.Context){
		ctx.Text(http.StatusOK, "POST /test")
	})
	
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
