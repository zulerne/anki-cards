package main

import "fmt"

type Server struct {
	name string
	port int
}
type Option func(*Server)

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
func NewServer(opts ...Option) *Server {
	server := &Server{}
	for _, o := range opts {
		o(server)
	}
	return server
}
func main() {
	srv := NewServer(WithName("my serv"), WithPort(3000))
	fmt.Printf("srv: %v\n", srv)
}
