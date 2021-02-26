package main

import "fmt"

const spanish = "Spanish"
const french = "Bonjour"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

//func Hello(name string, language string) string {
//	if name == "" {
//		name = "World"
//	}
//
//	prefix := englishHelloPrefix
//
//	switch language {
//	case french:
//		prefix = frenchHelloPrefix
//	case spanish:
//		prefix = spanishHelloPrefix
//	}
//
//	return prefix + name
//}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
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

func main() {
	fmt.Println(Hello("world", ""))
}
