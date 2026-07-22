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
	server Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (sb *ServerBuilder) WithName(name string) *ServerBuilder {
	if name == "" {
		sb.server.err = fmt.Errorf("invalid name")
	}
	sb.server.name = name
	return sb
}

func (sb *ServerBuilder) WithPort(port int) *ServerBuilder {
	if port <= 0 || port >= 65536 {
		sb.server.err = fmt.Errorf("invalid port")
	}
	sb.server.port = port
	return sb
}

func (sb *ServerBuilder) Build() (*Server, error) {
	if sb.server.err != nil {
		return nil, sb.server.err
	}
	return &sb.server, nil
}

func main() {
	server, err := NewServerBuilder().
		WithName("my serv").
		WithPort(3000).
		Build()

	if err != nil {
		log.Panic(err)
	}

	log.Println(server)
}
