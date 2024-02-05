package web

import (
	"log"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()

	routes := []Route{
		{Path: "/testtext", Handler: s.HandleTestText(), Methods: "GET"},
		{Path: "/testjson", Handler: s.HandleTestJSON(), Methods: "GET"},
	}

	config := Config{
		Url:    ":9999",
		Cors:   true,
		Routes: routes,
	}

	err := s.RunWebServer(config)
	if err != nil {
		log.Fatal(err)
	}
}
