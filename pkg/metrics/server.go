package metrics

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
)

type Server struct {
	Addr string
	mux  *http.ServeMux
}

var (
	HttpRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request",
			Help: "http request infromation",
		},
		[]string{"path","status_code"},
	)
)

func (s *Server) Start() error {
	s.mux = http.NewServeMux()

	prometheus.MustRegister(HttpRequest)
	s.mux.Handle("/metrics", promhttp.Handler())
	server := http.Server{
		Addr:    s.Addr,
		Handler: s.mux,
	}
	return server.ListenAndServe()
}
