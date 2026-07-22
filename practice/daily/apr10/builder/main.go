package main

import (
	"fmt"
	"log"
)

type Server struct {
	name string
	port int
	err  error
}

type ServerBuilder struct {
	s Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (sb *ServerBuilder) WithName(name string) *ServerBuilder {
	if name == "" {
		sb.s.err = fmt.Errorf("empty name")
	}
	sb.s.name = name
	return sb
}

func (sb *ServerBuilder) WithPort(port int) *ServerBuilder {
	if port <= 0 || port > 5000 {
		sb.s.err = fmt.Errorf("invalid port")
	}
	sb.s.port = port
	return sb
}

func (sb *ServerBuilder) Build() (*Server, error) {
	if sb.s.err != nil {
		return nil, sb.s.err
	}
	return &sb.s, nil
}

func main() {
	server, err := NewServerBuilder().
		WithName("build serv").
		WithPort(21222).
		Build()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server: %v\n", server)
}
