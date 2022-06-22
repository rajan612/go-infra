package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Goooooooo Conference 🍻"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("welcome to the our %v Booking Application. \nWe have total of %v tickets and %v are the remaining tickets.\nGet your tickets here to attend  🍻🍻\n ", conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//Asking User Details
		fmt.Println("what is your first name?")
		fmt.Scan(&firstName)

		fmt.Println("what is your flast name?")
		fmt.Scan(&lastName)

		fmt.Println("what is your email?")
		fmt.Scan(&email)

		fmt.Println("how many tickets you want to book?")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "a")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v booked %v tickets.\nRemainig Tickets are %v\n", firstName, userTickets, remainingTickets)
			fmt.Printf("You will receice a confirmation email on %v \n", email)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Please visit next year\n")
				break
			}
		} else {
			fmt.Println("Your input data is invalid, try again\n")
		}
	}

}
