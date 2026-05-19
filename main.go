package main

import (
	"fmt"

	"github.com/AlinaMavlekaeva/go-testing/hello"
)

func main() {
	greeting := hello.Hello("Alina")
	fmt.Println(greeting)
}
