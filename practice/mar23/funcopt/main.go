package main

import "fmt"

type Server struct {
	port int
	name string
}
type Option func(*Server)

func NewServer(opts ...Option) *Server {
	srv := &Server{}
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}
func withPort(n int) Option {
	return func(s *Server) {
		s.port = n
	}
}

func withName(n string) Option {
	return func(s *Server) {
		s.name = n
	}
}

func main() {
	srv := NewServer(withName("myserv"), withPort(3000))
	fmt.Println(srv)
}
