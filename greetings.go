package greetings

import (
	"fmt"
)

// Greet function take two variables name and language and return a greet
func Greet(name string, lang string) string {
	var result string

	if len(name) == 0 {
		result = getSimpleGreet(lang)
	} else {
		result = getGreetWithName(name, lang)
	}

	return result
}

func getSimpleGreet(lang string) string {
	var result string

	switch lang {
	case "it":
		result = "Ciao Amico!"
	case "fr":
		result = "Bonjour, mec!"
	case "de":
		result = "Hallo Kumpel!"
	case "es":
		result = "Hola, amigo!"
	default:
		result = "Hello Dude!"
	}

	return result
}

func getGreetWithName(name string, lang string) string {
	var result string

	switch lang {
	case "it":
		result = fmt.Sprintf("Ciao %v!", name)
	case "fr":
		result = fmt.Sprintf("Bonjour %v!", name)
	case "de":
		result = fmt.Sprintf("Hallo %v!", name)
	case "es":
		result = fmt.Sprintf("Hola %v!", name)
	default:
		result = fmt.Sprintf("Hello %v!", name)
	}

	return result
}
