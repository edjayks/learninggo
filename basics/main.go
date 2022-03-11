package main

import (
	"fmt"
	"strings"
)

func main() {

	// var conferenceName = "CNCF Conference"
	// sugar syntacting
	conferenceName := "CNCF Conference"
	const totalTickets = 50
	var remainingTickets uint = totalTickets
	fmt.Printf("Hello and Welcome to %v \n", conferenceName)

	bookings := []string{}
	// var names = [3]string{"Peter", "Bob", "Alice"}
	// fmt.Printf("All names %v\n", names)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter required ticket number: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole array: %v \n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Array Length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets, your tickets would be emailed to %v\n",
			firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets are now: %v\n", remainingTickets)

		// init slice with values
		// firstNames := []string{"Peter", "Alice", "Bob"}
		firstNames := []string{}
		// underscore for unused variable
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			// var firstName = names[0]
			firstNames = append(firstNames, names[0])
		}
		// fmt.Printf("Here is all our bookings: %v\n", bookings)
		fmt.Printf("Here is the first names all our bookings: %v\n", firstNames)

		hasRemainingTickets := remainingTickets == 0
		if hasRemainingTickets {
			// end program
			fmt.Println("No more available tickets, thank you everyone!")
			break
		}
	}

}
