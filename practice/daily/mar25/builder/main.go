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

func (b *ServerBuilder) WithName(name string) *ServerBuilder {
	if name == "" {
		b.server.err = fmt.Errorf("invalid name")
	}
	b.server.name = name
	return b
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	if port <= 0 || port > 9000 {
		b.server.err = fmt.Errorf("invalid port %d", port)
	}
	b.server.port = port
	return b
}

func (b *ServerBuilder) Build() (*Server, error) {
	if b.server.err != nil {
		return nil, b.server.err
	}
	return &b.server, nil
}

func main() {
	srv, err := NewServerBuilder().
		WithName("my").
		WithPort(409).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("srv: %v\n", srv)
}
