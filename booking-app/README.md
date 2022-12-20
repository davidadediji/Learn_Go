//array
//arrays in go have fixed size...
// var variable_name = [size]data_type{"some", "elements"}
//arrays does not allow mixed data types
//To declare an array in Go



// bookingsArray := [34]string{} //array
	// bookingsArray[0] = "5"
	// f.Println(bookingsArray)

	//print out the types of a variable and the constants
	// fmt.Printf("conferenceTickets is %T\n", remainingTickets)

var paidBooks [50]string
f.Printf("The lenght of the array is:%v", len(paidBooks))

books := []string{} // -> slice of strings a list (array ) that is dynamic in size. A slice is an array abstration with dynamic size.
books = append(books, "david", "favour")
f.Println(books)


    // f.Printf("The whole slice: %v\n", bookings)
    // f.Printf("The first value: %v\n", bookings[0])
    // f.Printf("Slice type: %T\n", bookings)
    // f.Printf("Slice length %v\n", len(bookings))
    // f.Print("This is David")

length of array in Go

len(array_variable);


Loops in Go

A loop is used to execute code multiple times
There is only one loop in Go : for-loop


For-Each loop


Package Level variables 

variables that are defined outside all functions 
Therefore this variables is assessible by any of the functions including the main function and in all files with the same package