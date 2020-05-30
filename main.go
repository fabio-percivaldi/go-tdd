package main

import "fmt"

const italian = "Italian"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const italianHelloPrefix = "Ciao, "
const spanishHelloPrefix = "Hola, "

// Hello returns greeting to specific name
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case italian:
		prefix = italianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Chris", ""))
}
