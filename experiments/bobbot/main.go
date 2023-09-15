package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bobbot := NewBobbot()
	fmt.Println(bobbot.WelcomeMessage())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Bobbot: Che vuoi?")
		fmt.Print("Tu: > ")
		question, _ := reader.ReadString('\n')
		fmt.Printf("Bobbot: %s\n\n", bobbot.Ask(question))
	}
}
