package main

import (
	"fmt"
	"strings"
)

func main() {
	// introduction and welcoming message
	var conferenceName string = "MODE3"
	fmt.Println(conferenceName)

	const conferenceTickets int = 20

	var remainingTickets uint = 20

	var bookings = [20]string{}

    Welcomeusers(conferenceName,conferenceTickets,remainingTickets)

	for { 
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName,email,userTickets, remainingTickets)

		

		if isValidEmail && isValidName && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets //deduct user booking from remainTickets

			bookings[0] = firstName + " " + lastName // slice

			fmt.Printf("Whole Slice: %v\n", bookings)
			fmt.Printf("The First Value: %v\n", bookings[0])
			fmt.Printf("Slice Type: %T\n", bookings)
			fmt.Printf("Slice lenght: %v\n", len(bookings))

			fmt.Printf("THANK YOU %v %v FOR BOOKING %v TICKET. YOU WILL RECIEVE A CONFIRMATION EMAIL AT %v THANK YOU\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Tickets are remaining for the conference\n", remainingTickets)

			fmt.Printf("THESE ARE ALL OUR BOOKINGS %v\n", bookings)

			if remainingTickets == 0 { // end booking
				fmt.Printf("OUR TICKECTS IS BOOKED OUT COME BACK LATER")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("First or last name is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("Must contain @\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets is invalid\n")
			}
			fmt.Printf("Invalid data input, please try again\n")
		}
	}
}


func Welcomeusers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("WELCOME TO %v BOOKING APP\n", conferenceName)
	fmt.Printf("WE HAVE A TOTAL OF %v TICKETS AND %v ARE STILL AVAILABLE FOR BOOKING\n", conferenceTickets, remainingTickets)
	fmt.Printf("Please Kindly Book your tickets now to Attend\n")

}



func validateUserInput(firstName string,lastName string, email string , userTickets uint, remainingTickets uint)(bool, bool, bool){ 
        var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
    	return isValidName, isValidEmail, isValidTicketNumber
}



func getUserInput() ( string , string,  string, uint,) {
   var firstName string
   var lastName string 
   var email string
   var userTickets uint
   
   
   fmt.Println("ENTER FIRST NAME:")
   fmt.Scan(&firstName)

   fmt.Println("ENTER LAST NAME:")
   fmt.Scan(&lastName)


   fmt.Println("ENTER YOUR EMAIL:")
   fmt.Scan(&email)

   fmt.Println("ENTER NUMBER OF TICKETS:")
   fmt.Scan(&userTickets)

   return firstName, lastName, email, userTickets
}