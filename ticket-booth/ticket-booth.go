package main

import (
	"fmt"
	"strings"	
)

func main() {

	bookings := []string{}

	conferenceName := "Go Conference"
	remainingTickets := 50
	const conferenceTickets uint = 50

	fmt.Println("======================================")
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now to attend!")
	fmt.Println("======================================")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		
		fmt.Print("\n")

		fmt.Print("Enter your First Name: ")
		fmt.Scan(&firstName)
		
		fmt.Print("Enter your Last Name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your Email: ")
		fmt.Scan(&email)

		fmt.Print("Tickets to book: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Println("======================================")
			fmt.Printf("Thank you, %v %v! For buying %v tickets. An email will be sent to %v for confirmation.\n", firstName, lastName, userTickets, email)
			fmt.Println("======================================")

			fmt.Printf("%v Remaining Tickets\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("People who booked: %v", firstNames)
			fmt.Println("\n======================================")

			if remainingTickets == 0 {
				fmt.Println("======================================")
				fmt.Println("Oops!! Tickets are sold out!😟 Come back next year.")
				fmt.Println("======================================")
				break
			}
		} else {
			fmt.Println("======================================")
			fmt.Printf("We only have %v remaining tickets!! You can't book %v tickets\n", remainingTickets, userTickets)
			fmt.Println("======================================")
		}
	}
}