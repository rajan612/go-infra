package main

import "fmt"

func main() {

	conferenceName := "Goooooooo Conference 🍻"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("welcome to the our %v Booking Application. \n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are the remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend  🍻🍻  ")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets.\n Remainig Tickets are %v\n", firstName, userTickets, remainingTickets)
	fmt.Printf("You will receice a confirmation email on %v \n", email)
}
