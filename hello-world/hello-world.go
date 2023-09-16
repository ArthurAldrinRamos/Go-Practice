package main

import (
	"fmt"
)

func ipagDikit(s1, s2 string) string {
	return s1 + s2
}


type catType struct {
	name string
	color string
	traits []string
}

func (c catType) catScratch() {
	fmt.Printf("%v scratched you!!!", c.name)
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

	// type ninja struct {

	// 	name string
	// 	weapons []string
	// 	level int

	// }

	// arthur := ninja {

	// 	name: "Arthur",
	// 	weapons: []string{"Katana", "Bow and Arrows", "Ninja Star"},
	// 	level: 1,
	
	// }

	// type dojo struct {

	// 	name string
	// 	ninja ninja

	// }


	// arthurDojo := dojo {

	// 	name: "Art Dojo",
	// 	ninja: arthur,

	// }

	// fmt.Println(arthur)
	// fmt.Println(arthur.name)
	// fmt.Printf("His weapons are: %v\n", arthur.weapons)
	// fmt.Printf("He is Level: %v\n", arthur.level)
	// fmt.Printf(arthurDojo.ninja.name)


	// // fmt.Println(firstNum, secondNum)


	// //PRACTICE

	// //for-each loop

	// s := "Hello World"

	// // for i, c := range s {
	// // 	fmt.Printf("%d %c", i, c)
	// // 	fmt.Println()
	// // }

	// for _, c:= range s {
	// 	fmt.Printf("%c", c)
	// 	fmt.Println()
	// }

	// listOfNames := []string{"kekw", "meow", "lmao"}

	// for i, names := range listOfNames {
	// 	fmt.Printf("%d) %v", i, names)
	// 	fmt.Println()
	// }

	// listOfSchool := []string{"DLSAU", "PUP", "UST", "UCC"}

	// for i, school := range listOfSchool {
	// 	fmt.Printf("%d) %v", i, school)
	// 	fmt.Println()
	// }
	 
	// listOfCats := []string{"Alyanna", "Strok", "Bunso", "Tonky", "Tjoeki", "Ponkan", "Oreo", "Duryan", "Pabo"}

	// for _, cats := range listOfCats {
	// 	fmt.Printf(cats)
	// 	fmt.Println()
	// }

	// alyanna := catType {
	// 	name: "Alyanna",
	// 	color: "Black and White",
	// 	traits: []string{"Loves aircon", "Fierce", "Loyal"},
	// }

	// bunso := catType {
	// 	name: "Bunso",
	// 	color: "Orange",
	// 	traits: []string{"Noisy af", "Dont know how to fight", "Weird Tail"},
	// }

	// strok := catType {
	// 	name: "Strok",
	// 	color: "Orange",
	// 	traits: []string{"Asshole", "Loyal", "Wobbly", "Ugly ass tail"},
	// }

	// fmt.Println(alyanna.name)
	// fmt.Println(bunso.name)
	// fmt.Println(strok.name)
	// alyanna.catScratch()


	//Pointers

	// a := 21
	// b := &a
	// c := a
	// fmt.Println(*b, c)
	// *b = 5
	// fmt.Println(a)
	// a = 23
	// fmt.Println(b)

	//Maps

	// var m map[string]string
	m := map[string]string{"Strok": "Baldadong Orenj na Pussy"}

	m["Alyanna"] = "Black Pussy"
	m["Bunso"] = "Orenj Kalbong Pussy"
	m["Tonky"] = "Orenj Iyaking Pussy"

	fmt.Println(m["Alyanna"])
	fmt.Println(m)
	
	anongPussy, ok := m["Tonky"] 
	if ok {
		fmt.Println(anongPussy)
	} else {
		fmt.Println("Pussy info doesn't exist")
	}

	if anongPussy, ok := m["Alyanna"]; ok {
		fmt.Println(anongPussy)
	} else {
		fmt.Println("Pussy info doesn't exist")
	}
}