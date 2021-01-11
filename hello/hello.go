package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(hello("World"))
	fmt.Println(rscHello())
}

func hello(name string) string {
	return fmt.Sprintf("Hello %v!", name)
}

func rscHello() string {
	return quote.Hello()
}
