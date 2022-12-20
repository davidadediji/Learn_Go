package main

import (
	f "fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{} //slice
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	//infinite loop
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings := bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email)
			firstNames := getFirstNames(bookings)
			f.Printf("The first names of bookings are: %v\n", firstNames)
			
			//conditionals (if-else)
			noTicketRemaining := remainingTickets == 0
			if noTicketRemaining {
				//end program
				f.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				f.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				f.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				f.Println("number of tickets you entered is invalid")
			}

		}
	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	f.Printf("Welcome to %v booking application\n", confName)
	// f.Println(time.Now().UnixNano())
	f.Printf("We have total of %v tickets and %v, are still avaliable.\n", confTickets, remainingTickets)
	f.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	//for each
	for _, booking := range bookings {
		names := strings.Fields(booking)

		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	//Go can return multiple values
	return isValidEmail, isValidName, isValidTicketNumber
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string ) ([]string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	f.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmaton email at %v\n", firstName, lastName, userTickets, email)
	f.Printf("remains %v tickets\n", remainingTickets)

	return bookings
}
