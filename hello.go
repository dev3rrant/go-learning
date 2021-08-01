package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	greetingPrefix := getGreetingPrefix(language)
	return greetingPrefix + name
}

func getGreetingPrefix(language string) (greetingPrefix string) {
	greetingsPrefixDefinition := map[string]string{
		"English":   "Hello, ",
		"Spainish":  "Hola, ",
		"French":    "Bonjour, ",
		"Esperanto": "Saluton, ",
	}
	greetingPrefix = greetingsPrefixDefinition[language]
	if greetingPrefix == "" {
		greetingPrefix = "Hello, "
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
