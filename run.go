package main

import "fmt"


func main(){
	//create an array
	var firstName string
	var lastName string
	const age uint = 24

	fmt.Scan(&firstName)
	fmt.Scan(&lastName)

	names := []string{}

	names = append(names, firstName + " " + lastName)
	fmt.Printf("The whole slice %v\n",names)
	fmt.Printf("Hello there\n")
}