package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	//req, _ := http.NewRequest("GET", "/ping", nil)
	req := httptest.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPing(t *testing.T) {
	ms := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		s := []byte("pong")
		w.Write(s)
		if r.Method != "GET" {
			t.Errorf("需要GET方法，不是 '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/" {
			t.Errorf("路径不对，'%s'", r.URL.EscapedPath())
		}

	}))
	defer ms.Close()
	resp, _ := http.Get(ms.URL)

	bodybytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(bytes.NewBuffer(bodybytes).String())
}

func BenchmarkPingRoute(b *testing.B) {
	router := setupRouter()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)
	}
}
