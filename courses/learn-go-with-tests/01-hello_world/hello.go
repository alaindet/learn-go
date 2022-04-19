package main

import "fmt"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	switch lang {
	case "en":
		return fmt.Sprintf("%s, %s", "Hello", name)
	case "it":
		return fmt.Sprintf("%s, %s", "Ciao", name)
	case "fr":
		return fmt.Sprintf("%s, %s", "Salut", name)
	case "de":
		return fmt.Sprintf("%s, %s", "Hallo", name)
	case "es":
		return fmt.Sprintf("%s, %s", "Hola", name)
	default:
		return fmt.Sprintf("%s, %s", "Hello", name)
	}
}

func main() {
	message := Hello("world", "it")
	fmt.Println(message)
}
