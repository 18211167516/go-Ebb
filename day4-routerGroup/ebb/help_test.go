
package ebb

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
)

type mockWriter struct {
	headers http.Header
}

func newMockWriter() *mockWriter {
	return &mockWriter{
		http.Header{},
	}
}

func (m *mockWriter) Header() (h http.Header) {
	return m.headers
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockWriter) WriteHeader(int) {}

func runRequest(B *testing.B, r *Engine, method, path string) {
	// create fake request
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	w := newMockWriter()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		r.ServeHTTP(w, req)
	}
}

func PerformRequest(mothod string,url string ,body io.Reader,e *Engine) (w *httptest.ResponseRecorder){
    w = httptest.NewRecorder()
	r := httptest.NewRequest(mothod, url, body)
	r.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w,r)
	return w
}

