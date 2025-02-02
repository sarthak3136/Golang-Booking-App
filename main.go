package main

import (
	"fmt"
	"strings"
)

func main() {
	//var conferenceName string = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // Other way to write var remainingTickets = 50
	//var bookings []string
	bookings := []string{}

	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		fmt.Println(remainingTickets)
		fmt.Println(&remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

// A Go function can return multiple values
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Arrays
func Array() {
	// One way to write array
	var bookings [50]string
	bookings[0] = "Carlos Alcaraz"
	fmt.Println(bookings[0])

	// To initialize values in an array
	var persons = [50]string{"Carlos", "Alcaraz"}
	fmt.Println(persons[0])
}

// Slices or Similar to ArrayList in Java
func Slice() {
	var bookings []string
	bookings = append(bookings, "Sarthak")
	fmt.Println(bookings)

	var persons = []string{"Carlos"}
	fmt.Println(persons)

}

// Loops
func Loops() {
	// For Each Loop
	// If you don't want to use index then use blank identifier "_" instead of "index"
	var bookings = []string{"Carlos Alcaraz", "Novak Djokovich"}
	firstNames := []string{}
	for index, booking := range bookings {
		var names = strings.Fields(booking)
		index++
		firstNames = append(firstNames, names[0])
	}

	// For Loops with condition
	var remainingTickets uint = 12

	for remainingTickets > 0 && len(bookings) < 50 {

	}

	// Infinite Loop
	for {
		var firstName string
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
	}

}

// To create a module: go mod init booking-app
