package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("streamed data")
	var destination bytes.Buffer

	if _, err := io.Copy(&destination, source); err != nil {
		panic(err)
	}
	fmt.Println(destination.String())
}
