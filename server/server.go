package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	GetAddress() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type SimpleServer struct {
	Address string
	Proxy   *httputil.ReverseProxy
}

func NewSimpleServer(address string) *SimpleServer {
	serverUrl, err := url.Parse(address)

	if err != nil {
		fmt.Sprintf("error: %v\n", err)
		os.Exit(1)
	}
	return &SimpleServer{
		Address: address,
		Proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func (s *SimpleServer) GetAddress() string {
	return s.Address
}

func (s *SimpleServer) IsAlive() bool {
	return true
}

func (s *SimpleServer) Serve(rw http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(rw, r)
}
