package server

import (
	"net/http"
)

type Server struct {
	Addr string
	mux  *http.ServeMux
}

func (s *Server) Start() error {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:    s.Addr,
		Handler: s.mux,
	}
	return server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	statusCode := 200
	switch path {
	case "/401":
		statusCode = 401
	case "/500":
		statusCode = 500
	case "/502":
		statusCode = 502
	}
	w.WriteHeader(statusCode)
	w.Write([]byte("ok"))
}
