package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// NewEngine 新建一个测试引擎
func NewEngine() (engine *gin.Engine) {
	engine = gin.New()
	return
}

// KeyValue .
type KeyValue struct {
	Key   string
	Value string
}

// NewKeyValue .
func NewKeyValue(key string, value string) KeyValue {
	return KeyValue{Key: key, Value: value}
}

// NewRequest 新建请求
func NewRequest(method, path string, body io.Reader) *http.Request {
	return httptest.NewRequest(method, path, body)
}

// NewGetRequest .
func NewGetRequest(path string, queries ...KeyValue) *http.Request {
	req := NewRequest(http.MethodGet, path, nil)
	req = AddQueries(req, queries...)
	return req
}

// NewPostRequest .
func NewPostRequest(path string, contentType string, body io.Reader) *http.Request {
	req := NewRequest(http.MethodPost, path, body)
	req.Header.Add("Content-Type", contentType)
	return req
}

// NewPostFormRequest .
func NewPostFormRequest(path string, form ...KeyValue) *http.Request {
	data := make(url.Values)
	for _, kv := range form {
		data.Add(kv.Key, kv.Value)
	}
	return NewPostRequest(path, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

// AddQueries .
func AddQueries(req *http.Request, queries ...KeyValue) *http.Request {
	query := req.URL.Query()
	for _, kv := range queries {
		query.Add(kv.Key, kv.Value)
	}
	req.URL.RawQuery = query.Encode()
	return req
}

// Do 执行请求
func Do(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	return recorder
}
