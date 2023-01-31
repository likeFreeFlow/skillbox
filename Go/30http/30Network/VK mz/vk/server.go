package vk

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	fmt.Println("Server is running")
	return s.httpServer.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)

}
