package main

import (
	"fmt"
)

func main() {
	// #1
	type person struct {
		name           string
		age            int
		favoriteColors []string
	}

	me := person{
		name:           "John",
		age:            20,
		favoriteColors: []string{"Red", "Blue"},
	}

	you := person{
		name:           "Jane",
		age:            20,
		favoriteColors: []string{"Green", "Lilac"},
	}

	fmt.Printf("%+v\n", me)  // {name:John age:20 favoriteColors:[Red Blue]}
	fmt.Printf("%+v\n", you) // {name:Jane age:20 favoriteColors:[Green Lilac]}

	// #2
	me.name = "New Name"
	yourFavoriteColors := you.favoriteColors
	you.favoriteColors = append(you.favoriteColors, "Light Blue")
	fmt.Println(yourFavoriteColors) // [Green Lilac]
	fmt.Println(you)                // {Jane 20 [Green Lilac Light Blue]}

	// #3
	for index, color := range me.favoriteColors {
		fmt.Printf("%d => %v\n", index, color)
	}
	// 0 => Red
	// 1 => Blue

	// #4
	type grade struct {
		course string
		grade  int
	}

	type student struct {
		name           string
		age            int
		favoriteColors []string
		grades         []grade
	}

	studentA := student{
		name:           "Student A",
		age:            20,
		favoriteColors: []string{"Green", "Black"},
		grades: []grade{
			grade{
				course: "Golang",
				grade:  98,
			},
			grade{
				course: "JavaScript",
				grade:  98,
			},
		},
	}

	// {name:Student A age:20 favoriteColors:[Green Black] grades:[{course:Golang grade:98} {course:JavaScript grade:98}]}
	fmt.Printf("%+v\n", studentA)
}
