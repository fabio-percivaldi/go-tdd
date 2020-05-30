package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns greeting to specific name
func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
