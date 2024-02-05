package web

import (
	"fmt"
	"net/http"
)

func (s *Server) HandleTestText() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test handler Text (from gotools module)")
	}
}

func (s *Server) HandleTestJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Response string `json:"response"`
		}{Response: "Test handler JSON (from gotools module)"}
		s.RespondJSON(w, r, res, http.StatusOK)
	}
}

func HandleTestText2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test handler Text (from gotools module)")
	}
}
