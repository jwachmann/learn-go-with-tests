package main

import (
	"fmt"

	"rsc.io/quote"
)

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello"
const frenchHelloPrefix = "Bonjour"
const spanishHelloPrefix = "Hola"

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return fmt.Sprintf("%v %v!", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func rscHello() string {
	return quote.Hello()
}
