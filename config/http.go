package config

import (
	"net/http"
	"time"
)

func NewHTTPServer(addr string) *http.Server {
	return &http.Server{
		Addr:           addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
