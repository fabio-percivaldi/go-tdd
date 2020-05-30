package main

import "fmt"

// Hello returns greeting to specific name
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
