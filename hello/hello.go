package hello

import "fmt"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

// Hello returns Hello, world
func Hello(name, lang string) string {
	greeting := getGreeting(name)
	prefix := getPrefix(lang)
	return fmt.Sprintf("%s %s", prefix, greeting)
}

func getPrefix(lang string) string {
	switch lang {
	case "Spanish":
		return spanishHelloPrefix
	case "English":
		return englishHelloPrefix
	case "French":
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func getGreeting(name string) string {
	if name == "" {
		name = "World"
	}
	return name
}
