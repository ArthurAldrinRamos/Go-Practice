package main

import (
	"fmt"
)

func ipagDikit(s1, s2 string) string {
	return s1 + s2
}

func main(){

	// fmt.Println("Hello World")

	// firstNum := 1
	// secondNum := 2

	// firstName := "Arthur "
	// lastName := "Ramos"

	// if (firstNum > secondNum) {
	// 	fmt.Println("First Number is greater than the Second Number")
	// } else {
	// 	fmt.Println("Mehhh")
	// }

	// fmt.Println(ipagDikit(firstName, lastName))

	type ninja struct {

		name string
		weapons []string
		level int

	}

	arthur := ninja {

		name: "Arthur",
		weapons: []string{"Katana", "Bow and Arrows", "Ninja Star"},
		level: 1,
	
	}

	type dojo struct {

		name string
		ninja ninja

	}


	arthurDojo := dojo {

		name: "Art Dojo",
		ninja: arthur,

	}

	fmt.Println(arthur)
	fmt.Println(arthur.name)
	fmt.Printf("His weapons are: %v\n", arthur.weapons)
	fmt.Printf("He is Level: %v\n", arthur.level)
	fmt.Printf(arthurDojo.ninja.name)


	// fmt.Println(firstNum, secondNum)
	
}