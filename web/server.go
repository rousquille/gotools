package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Config is a configuration container to set up Server
type Config struct {
	Url    string
	Cors   bool
	Routes []Route
}

type Server struct {
	router *mux.Router
	config Config
}

type Route struct {
	Path       string
	Handler    http.Handler
	Methods    string
	PathPrefix bool
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	return s
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	LogRequestMiddleware(s.router.ServeHTTP).ServeHTTP(w, r)
}

func (s *Server) RunWebServer(config Config) error {
	var err error
	s.config = config

	s.Routing()

	if s.config.Cors {
		s.router.Use(CorsMiddleware)
	}

	http.HandleFunc("/", s.serveHTTP)
	log.Println("Serving HTTP on", s.config.Url)

	err = http.ListenAndServe(s.config.Url, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DecodeJSON(_ http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) RespondJSON(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func (s *Server) Routing() {
	for _, route := range s.config.Routes {
		if !route.PathPrefix {
			s.router.Handle(route.Path, route.Handler).Methods(route.Methods)
		} else {
			s.router.PathPrefix(route.Path).Handler(route.Handler)
		}
	}
}
