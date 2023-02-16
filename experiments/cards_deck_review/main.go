package main

import (
	"fmt"
	"strings"
)

func main() {

	d := NewDeck()

	for i := 0; i < 100; i++ {

		err := d.Burn(1)

		if err != nil {
			fmt.Println(err)
			break
		}

		drawn, err := d.Draw(2)

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("#%d draw 2: %s\n", i, strings.Join(drawn, ", "))
	}
}
