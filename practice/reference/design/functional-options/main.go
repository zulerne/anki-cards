package main

import "fmt"

type Server struct {
	Name string
	Port int
}

type Option func(*Server)

func WithName(name string) Option {
	return func(server *Server) { server.Name = name }
}

func WithPort(port int) Option {
	return func(server *Server) { server.Port = port }
}

func NewServer(options ...Option) *Server {
	server := &Server{Port: 8080}
	for _, option := range options {
		option(server)
	}
	return server
}

func main() {
	fmt.Printf("%+v\n", NewServer(WithName("api"), WithPort(9000)))
}
