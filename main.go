package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const japaneseHelloPrefix = "こんにちは, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "English":
		return englishHelloPrefix + name
	case "Portuguese":
		return portugueseHelloPrefix + name
	case "Japanese":
		return japaneseHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("TDD", ""))
}
