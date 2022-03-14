package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName = "CNCF Conference"

const totalTickets = 50

var remainingTickets uint = totalTickets

// an empty list of strings, the curly braces init it to empty list
// var bookings = []string{}
// an empty list of maps , the curly braces init it to empty list of maps
// var bookings = []map[string]string{}

// an empty list of maps , using make function to init an empty list of maps
// var bookings = make([]map[string]string, 0)

// an empty list of struct
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// var conferenceName = "CNCF Conference"
	// sugar syntacting
	// conferenceName := "CNCF Conference"
	// const totalTickets = 50
	// var remainingTickets uint = totalTickets
	// bookings := []string{}
	// var names = [3]string{"Peter", "Bob", "Alice"}
	// fmt.Printf("All names %v\n", names)

	greetUsers(totalTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		// input validation
		isValidName, isValidEmail, isValidUserTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicket {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("Here is the first names all our bookings: %v\n", firstNames)

			hasRemainingTickets := remainingTickets == 0
			if hasRemainingTickets {
				// end program
				fmt.Println("No more available tickets, thank you everyone!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid first name and last name, too short")
			}

			if !isValidEmail {
				fmt.Println("Invalid email address")
			}

			if !isValidUserTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers(totalTickets int, remainingTickets uint) {
	fmt.Println("===== ************************************************** =====")
	fmt.Printf("Hello and Welcome to %v \n", conferenceName)
	fmt.Printf("A total of %v tickets is available!\n", totalTickets)
	fmt.Println("===== ************************************************** =====")
}

func getFirstNames() []string {
	// init slice with values
	// firstNames := []string{"Peter", "Alice", "Bob"}
	firstNames := []string{}
	// underscore for unused variable
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		// firstNames = append(firstNames, names[0])

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("===== Please Enter Your Booking Information =====")
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter required ticket number: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map to store user data
	// var userData map[string]string
	// create an empty map
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// create a struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v", bookings)

	// fmt.Printf("The whole array: %v \n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Array Length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets, your tickets would be emailed to %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets are now: %v\n", remainingTickets)
}
