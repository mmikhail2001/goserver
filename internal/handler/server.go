package handler

import "net/http"

type Server interface {
	Run() error
}

type server struct {
	httpServer http.Server
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func newServer(addr string, mux *http.ServeMux) Server {
	return &server{
		httpServer: http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}
