package server

import (
	"net/http"
)

func (s *Server) Router() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("POST /quotes", s.QuoteHandlerHandler.Create)
	r.HandleFunc("GET /quotes", s.QuoteHandlerHandler.Get)
	r.HandleFunc("GET /quotes/random", s.QuoteHandlerHandler.GetRandom)
	r.HandleFunc("DELETE /quotes/{id}", s.QuoteHandlerHandler.DeleteByID)

	return r
}
