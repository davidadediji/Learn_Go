package main

import (
	f "fmt"
	"strings"
	"time"
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
	f.Println(time.Now().UnixNano())
	f.Printf("We have total of %v tickets and %v, are still avaliable.\n", conferenceTickets, remainingTickets)
	f.Println("Get your tickets here to attend")

	//infinite loop
	for {
		//declare variables and leave unassigned....
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// f.Printf("The whole slice: %v\n", bookings)
			// f.Printf("The first value: %v\n", bookings[0])
			// f.Printf("Slice type: %T\n", bookings)
			// f.Printf("Slice length %v\n", len(bookings))
			f.Print("This is David")

			f.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmaton email at %v\n", firstName, lastName, userTickets, email)
			f.Printf("remains %v tickets\n", remainingTickets)

			firstNames := []string{}

			//for each
			for _, booking := range bookings {
				names := strings.Fields(booking)

				firstNames = append(firstNames, names[0])
			}
			f.Printf("The first names of bookings are: %v\n", firstNames)

			//conditionals (if-else)
			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				//end program
				f.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			f.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}

}

func greetUsers() {
	conferenceName := "Go Conference"
	f.Printf("Welcome to %v booking application\n", conferenceName)
}
