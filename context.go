package fwsample

import (
	"encoding/json"
	"net/http"
)

type Context interface {
	Request() *http.Request
	ResponseWriter() http.ResponseWriter

	// パスパラメータの取得
	Param(key string) string

	// レスポンスにヘッダを追加
	SetHeader(key, value string)

	// ステータスコートを設定
	WriteHeader(statusCode int)

	// JSONレスポンス出力
	JSON(statusCode int, data interface{}) error

	// テキストレスポンス出力
	Text(statusCode int, text string) error
}

type context struct {
	req       *http.Request
	rw http.ResponseWriter
	params         map[string]string // URLパスから抽出したパスパラメータを格納
}

func NewContext(w http.ResponseWriter, r *http.Request) *context {
	return &context{
		req: r,
		rw: w,
		params: map[string]string{},
	}
}

func (c *context) Request() *http.Request {
	return c.req
}

func (c *context) ResponseWriter() http.ResponseWriter {
	return c.rw
}

func (c *context) Param(key string) string {
	return c.params[key]
}

func (c *context) SetHeader(key, value string) {
	c.rw.Header().Set(key, value)
}

func (c *context) WriteHeader(statusCode int) {
	c.rw.WriteHeader(statusCode)
}

func (c *context) JSON(statusCode int, data interface{}) error {
	c.SetHeader("Content-Type", "application/json")
	c.WriteHeader(statusCode)
	return json.NewEncoder(c.rw).Encode(data)
}

func (c *context) Text(statusCode int, text string) error {
	c.SetHeader("Content-Type", "application/json")
	c.WriteHeader(statusCode)
	_, err := c.rw.Write([]byte(text))
	return err
}