package main

import f "fmt"

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) //list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	//infinite loop
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
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

func greetUsers() {
	f.Printf("Welcome to %v booking application\n", conferenceName)
	// f.Println(time.Now().UnixNano())
	f.Printf("We have total of %v tickets and %v, are still avaliable.\n", conferenceTickets, remainingTickets)
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

func getFirstNames() []string {
	firstNames := []string{}
	//for each
	for _, booking := range bookings {
		// names := strings.Fields(booking)

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	f.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmaton email at %v\n", firstName, lastName, userTickets, email)
	f.Printf("remains %v tickets\n", remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = f.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	f.Println("#############")

	f.Printf("sending ticket:\n %v \nto email address %v\n", ticket, email)
	f.Println("#############")
}

