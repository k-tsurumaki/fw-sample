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
