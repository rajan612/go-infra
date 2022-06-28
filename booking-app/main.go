package main

import (
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Goooooooo Conference 🍻"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		//call validateuserinput function
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The First name of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Please visit next year\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is to short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered does not contains @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid")
			}
			fmt.Println("Your input data is invalid, try again\n")
		}
	}
}
func greetUsers() {
	fmt.Printf("welcome to the our %v Booking Application. \nWe have total of %v tickets and %v are the remaining tickets.\nGet your tickets here to attend  🍻🍻\n", conferenceName, conferenceTickets, remainingTickets)

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("The first names of the bookings are: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Asking User Details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter numb er of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (string, uint, uint, string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("User %v booked %v tickets.\nRemainig Tickets are %v\n", firstName, userTickets, remainingTickets)
	fmt.Printf("You will receice a confirmation email on %v \n", email)
	return firstName, userTickets, remainingTickets, email
}
