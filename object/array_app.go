package main

import "fmt"

func main() {
	const concertTickets int = 4000
	var remainingTickets uint = 4000
	concertName := "D.R.I Tickets"
	bookings := []string{}

	fmt.Println("Ready to mosh to", concertTickets, " people. \nTotal of", remainingTickets, "still available. \nDon't stand there, get your tickets!\n")

	//collect user information via data type declaration

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter # of tickets: ")
	fmt.Scanln(&userTickets)

	//Logic
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	//Output
	fmt.Printf("Thanks %v %v for booking %v tickets. An email will be sent with confirmation at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, concertName)
	fmt.Printf("These are all of our bookings: %v\n", bookings)
}