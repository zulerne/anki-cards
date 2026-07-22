package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("not found")

func load() error {
	return errors.Join(errNotFound, errors.New("cache unavailable"))
}

func main() {
	if err := load(); err != nil {
		fmt.Println(errors.Is(err, errNotFound))
	}
}
