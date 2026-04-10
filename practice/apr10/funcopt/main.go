package main

import "fmt"

type Server struct {
	name string
	port int
}

type Option func(s *Server)

func NewServer(opts ...Option) *Server {
	srv := &Server{}
	for _, o := range opts {
		o(srv)
	}
	return srv
}

func WithName(name string) Option {
	return func(s *Server) {
		s.name = name
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}
func main() {
	srv := NewServer(WithName("my server"), WithPort(1222))
	fmt.Printf("srv: %v\n", srv)
}
