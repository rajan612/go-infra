package main

import "fmt"

func main() {

	conferenceName := "Goooooooo Conference 🍻"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("welcome to the our %v Booking Application. \n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are the remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend  🍻🍻  ")

	var userName string
	var userTickets int

	//Asking User Details
	fmt.Printf("what is your name?\n")
	fmt.Scan(&userName)

	fmt.Printf("how many tickets you want to book?\n")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}
