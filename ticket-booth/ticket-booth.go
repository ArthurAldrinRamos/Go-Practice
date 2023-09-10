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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketCount {
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
				fmt.Println("ERROR: INPUT DATA IS WRONG. PLEASE TRY AGAIN!")

			if !isValidName {
				fmt.Println("-Please Check your Name: It should be at least 2 characters long.")
			}

			if !isValidEmail {
				fmt.Println("-Please Check your Email: It should contanin '@'")
			}

			if !isValidTicketCount {
				fmt.Printf("-We only have %v tickets remaining! We can't process %v tickets.\n", remainingTickets, userTickets)
			}

				fmt.Println("======================================")
		}
	}
}