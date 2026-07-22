package main

import (
	"errors"
	"fmt"
)

type Server struct {
	Name string
	Port int
}

type ServerBuilder struct {
	server Server
	err    error
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (b *ServerBuilder) WithName(name string) *ServerBuilder {
	if b.err != nil {
		return b
	}
	if name == "" {
		b.err = errors.New("name must not be empty")
		return b
	}
	b.server.Name = name
	return b
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	if b.err != nil {
		return b
	}
	if port < 1 || port > 65535 {
		b.err = fmt.Errorf("invalid port %d", port)
		return b
	}
	b.server.Port = port
	return b
}

func (b *ServerBuilder) Build() (Server, error) {
	if b.err != nil {
		return Server{}, b.err
	}
	return b.server, nil
}

func main() {
	server, err := NewServerBuilder().
		WithName("api").
		WithPort(8080).
		Build()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", server)
}
