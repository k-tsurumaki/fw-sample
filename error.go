package fwsample

import (
	"errors"
	"fmt"
	"net/http"
)

// Errors
var (
	ErrBadRequest                    = NewError(http.StatusBadRequest)                    // HTTP 400 Bad Request
	ErrUnauthorized                  = NewError(http.StatusUnauthorized)                  // HTTP 401 Unauthorized
	ErrPaymentRequired               = NewError(http.StatusPaymentRequired)               // HTTP 402 Payment Required
	ErrForbidden                     = NewError(http.StatusForbidden)                     // HTTP 403 Forbidden
	ErrNotFound                      = NewError(http.StatusNotFound)                      // HTTP 404 Not Found
	ErrMethodNotAllowed              = NewError(http.StatusMethodNotAllowed)              // HTTP 405 Method Not Allowed
	ErrNotAcceptable                 = NewError(http.StatusNotAcceptable)                 // HTTP 406 Not Acceptable
	ErrProxyAuthRequired             = NewError(http.StatusProxyAuthRequired)             // HTTP 407 Proxy AuthRequired
	ErrRequestTimeout                = NewError(http.StatusRequestTimeout)                // HTTP 408 Request Timeout
	ErrConflict                      = NewError(http.StatusConflict)                      // HTTP 409 Conflict
	ErrGone                          = NewError(http.StatusGone)                          // HTTP 410 Gone
	ErrLengthRequired                = NewError(http.StatusLengthRequired)                // HTTP 411 Length Required
	ErrPreconditionFailed            = NewError(http.StatusPreconditionFailed)            // HTTP 412 Precondition Failed
	ErrStatusRequestEntityTooLarge   = NewError(http.StatusRequestEntityTooLarge)         // HTTP 413 Payload Too Large
	ErrRequestURITooLong             = NewError(http.StatusRequestURITooLong)             // HTTP 414 URI Too Long
	ErrUnsupportedMediaType          = NewError(http.StatusUnsupportedMediaType)          // HTTP 415 Unsupported Media Type
	ErrRequestedRangeNotSatisfiable  = NewError(http.StatusRequestedRangeNotSatisfiable)  // HTTP 416 Range Not Satisfiable
	ErrExpectationFailed             = NewError(http.StatusExpectationFailed)             // HTTP 417 Expectation Failed
	ErrTeapot                        = NewError(http.StatusTeapot)                        // HTTP 418 I'm a teapot
	ErrMisdirectedRequest            = NewError(http.StatusMisdirectedRequest)            // HTTP 421 Misdirected Request
	ErrUnprocessableEntity           = NewError(http.StatusUnprocessableEntity)           // HTTP 422 Unprocessable Entity
	ErrLocked                        = NewError(http.StatusLocked)                        // HTTP 423 Locked
	ErrFailedDependency              = NewError(http.StatusFailedDependency)              // HTTP 424 Failed Dependency
	ErrTooEarly                      = NewError(http.StatusTooEarly)                      // HTTP 425 Too Early
	ErrUpgradeRequired               = NewError(http.StatusUpgradeRequired)               // HTTP 426 Upgrade Required
	ErrPreconditionRequired          = NewError(http.StatusPreconditionRequired)          // HTTP 428 Precondition Required
	ErrTooManyRequests               = NewError(http.StatusTooManyRequests)               // HTTP 429 Too Many Requests
	ErrRequestHeaderFieldsTooLarge   = NewError(http.StatusRequestHeaderFieldsTooLarge)   // HTTP 431 Request Header Fields Too Large
	ErrUnavailableForLegalReasons    = NewError(http.StatusUnavailableForLegalReasons)    // HTTP 451 Unavailable For Legal Reasons
	ErrInternalServerError           = NewError(http.StatusInternalServerError)           // HTTP 500 Internal Server Error
	ErrNotImplemented                = NewError(http.StatusNotImplemented)                // HTTP 501 Not Implemented
	ErrBadGateway                    = NewError(http.StatusBadGateway)                    // HTTP 502 Bad Gateway
	ErrServiceUnavailable            = NewError(http.StatusServiceUnavailable)            // HTTP 503 Service Unavailable
	ErrGatewayTimeout                = NewError(http.StatusGatewayTimeout)                // HTTP 504 Gateway Timeout
	ErrHTTPVersionNotSupported       = NewError(http.StatusHTTPVersionNotSupported)       // HTTP 505 HTTP Version Not Supported
	ErrVariantAlsoNegotiates         = NewError(http.StatusVariantAlsoNegotiates)         // HTTP 506 Variant Also Negotiates
	ErrInsufficientStorage           = NewError(http.StatusInsufficientStorage)           // HTTP 507 Insufficient Storage
	ErrLoopDetected                  = NewError(http.StatusLoopDetected)                  // HTTP 508 Loop Detected
	ErrNotExtended                   = NewError(http.StatusNotExtended)                   // HTTP 510 Not Extended
	ErrNetworkAuthenticationRequired = NewError(http.StatusNetworkAuthenticationRequired) // HTTP 511 Network Authentication Required

	ErrValidatorNotRegistered = errors.New("validator not registered")
	ErrRendererNotRegistered  = errors.New("renderer not registered")
	ErrInvalidRedirectCode    = errors.New("invalid redirect status code")
	ErrCookieNotFound         = errors.New("cookie not found")
	ErrInvalidCertOrKeyType   = errors.New("invalid cert or key type, must be string or []byte")
	ErrInvalidListenerNetwork = errors.New("invalid listener network")
)

type FwError struct {
	Code    int
	Message interface{}
	Err     error
}

func (e *FwError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("code=%d, message=%v, error=%v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("code=%d, message=%v", e.Code, e.Message)
}

func (e *FwError) Unwrap() error {
	return e.Err
}

func (e *FwError) Wrap(err error) error {
	e.Err = err
	return e
}

func NewError(code int, msg ...interface{}) *FwError {
	e := &FwError{Code: code, Message: http.StatusText(code)}
	if len(msg) > 0 {
		e.Message = msg[0]
	}
	return e
}