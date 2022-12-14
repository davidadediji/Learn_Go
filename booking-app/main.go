package main

import (
	f "fmt"
)

func main() {

	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}        //slice
	bookingsArray := [34]string{} //array
	bookingsArray[0] = "5"
	f.Println(bookingsArray)

	//print out the types of a variable and the constants
	// fmt.Printf("conferenceTickets is %T\n", remainingTickets)
	greetUsers()

	f.Printf("We have total of %v tickets and %v, are still avaliable.\n", conferenceTickets, remainingTickets)
	f.Println("Get your tickets here to attend")
	//Arrays - arrays in Go have fixed size
	//infinite loop
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		f.Println("Enter your first name: ")
		f.Scan(&firstName)

		f.Println("Enter your last name: ")
		f.Scan(&lastName)

		f.Println("Enter your email address: ")
		f.Scan(&email)

		f.Println("Enter number of tickets: ")
		f.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// f.Printf("The whole slice: %v\n", bookings)
		// f.Printf("The first value: %v\n", bookings[0])
		// f.Printf("Slice type: %T\n", bookings)
		// f.Printf("Slice length %v\n", len(bookings))

		f.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmaton email at %v\n", firstName, lastName, userTickets, email)
		f.Printf("remains %v tickets\n", remainingTickets)
		f.Printf("All the bookings %v\n", bookings)
	}

}

func greetUsers() {
	conferenceName := "Go Conference"
	f.Printf("Welcome to %v booking application\n", conferenceName)
}
