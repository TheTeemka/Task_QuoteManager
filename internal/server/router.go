package server

import (
	"net/http"
)

func (s *Server) Router() http.Handler {
	r := http.NewServeMux()

	return r
}
