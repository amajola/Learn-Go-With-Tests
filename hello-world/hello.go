package main // Packages are ways of grouping up related Go code together [Need to know more about this]

import (
	"fmt" // How go imports external packages
)

const (
	zulu               = "Zulu"
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	zuluHelloPrefix    = "Kodwa uyinja yaz wena, "
	frenchHelloPrefix  = "Hola, "
	spanishHelloPrefix = "Bonjour, "
) // Declaring const in one go is something I did not know I needed

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {

	switch language {
	case zulu:
		prefix = zuluHelloPrefix
	case spanish:
		prefix = frenchHelloPrefix
	case french:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
