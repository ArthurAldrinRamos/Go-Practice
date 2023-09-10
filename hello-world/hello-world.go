package main

import "fmt"

func ipagDikit(s1, s2 string) string {
	return s1 + s2
}

func main(){


	firstNum := 1
	secondNum := 2

	firstName := "Arthur "
	lastName := "Ramos"

	if (firstNum > secondNum) {
		fmt.Println("First Number is greater than the Second Number")
	} else {
		fmt.Println("Mehhh")
	}

	fmt.Println(ipagDikit(firstName, lastName))



	// fmt.Println(firstNum, secondNum)
	
}