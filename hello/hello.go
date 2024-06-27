package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	portugueseBR = "Portuguese-BR"

	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	portugueseBRHellowPrefix = "Olá, "
	englishHelloPrefix = "Hello, "
)

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		case portugueseBR: 
			prefix = portugueseBRHellowPrefix
		default:
			prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(HelloWorld("World", ""))
	fmt.Println(HelloWorld("", ""))
}