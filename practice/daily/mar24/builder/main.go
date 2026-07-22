package main

import (
	"fmt"
	"log"
)

type Server struct {
	name string
	port int
	err  error // копим ошибки валидации
}

type ServerBuilder struct {
	server Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (b *ServerBuilder) WithName(name string) *ServerBuilder {
	if name == "" {
		b.server.err = fmt.Errorf("name cannot be empty")
	}
	b.server.name = name
	return b // возвращаем себя для chaining
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	if port < 1 || port > 65535 {
		b.server.err = fmt.Errorf("invalid port: %d", port)
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
		WithName("myserver").
		WithPort(3000).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("srv: %v\n", srv)
}
