package fwsample

import (
	"net/http"
)

// Headers
const (
	HeaderAccept              = "Accept"
	HeaderAcceptEncoding      = "Accept-Encoding"
	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderRetryAfter          = "Retry-After"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-Ip"
	HeaderXRequestID          = "X-Request-Id"
	HeaderXCorrelationID      = "X-Correlation-Id"
	HeaderXRequestedWith      = "X-Requested-With"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"
	HeaderCacheControl        = "Cache-Control"
	HeaderConnection          = "Connection"

	// Access control
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"

	// Security
	HeaderStrictTransportSecurity         = "Strict-Transport-Security"
	HeaderXContentTypeOptions             = "X-Content-Type-Options"
	HeaderXXSSProtection                  = "X-XSS-Protection"
	HeaderXFrameOptions                   = "X-Frame-Options"
	HeaderContentSecurityPolicy           = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	HeaderXCSRFToken                      = "X-CSRF-Token"
	HeaderReferrerPolicy                  = "Referrer-Policy"
)

// アプリケーション全体を管理する構造体
type App struct {
	Router     RouterInterface
	Middleware []Middleware
}

type HandlerFunc http.HandlerFunc

type Middleware func(HandlerFunc) HandlerFunc

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	StatusCode int
}

func (w *ResponseWriterWithStatus) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

type RouterInterface interface {
	Add(method, path string, handler func(http.ResponseWriter, *http.Request)) error
	Get(path string, handler func(http.ResponseWriter, *http.Request)) error
	Post(path string, handler func(http.ResponseWriter, *http.Request)) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (a *App) Use(m Middleware) {
	a.Middleware = append(a.Middleware, m)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		a.Router.ServeHTTP(w, r)
	}

	// apply middleware in reverse order
	for i := len(a.Middleware) - 1; i >= 0; i-- {
		handler = a.Middleware[i](handler)
	}

	handler(w, r)
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
