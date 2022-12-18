package main

import "strings"

type Predicate[T any] func(t T) bool

type Formatted[T any] func(t T) string

func publish(text string, filter Predicate[string], formatted Formatted[string]) {
	if filter(text) {
		println(formatted(text))
	}
}

func main() {
	errorFilter := func(text string) bool {
		return strings.HasPrefix(text, "ERROR")
	}

	errorFormatter := func(text string) string {
		return strings.ToUpper(text)
	}

	publish("ERROR - something bad happened", errorFilter, errorFormatter)
	publish("DEBUG - I'm here", errorFilter, errorFormatter)
}
