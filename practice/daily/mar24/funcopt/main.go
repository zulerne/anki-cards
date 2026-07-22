package main

import "fmt"

type Server struct {
	name string
	port int
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	server := &Server{}
	for _, opt := range opts {
		opt(server)
	}
	return server
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
	srv := NewServer(WithName("sever"), WithPort(3000))
	fmt.Printf("srv: %v\n", srv)
}
