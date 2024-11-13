package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/username/booking-app/helper"
)

// these are the global variable so it can be accessed throughout the app
const conferenceTickets = 50

var conferenceName = "Go Conference"

var remainingTickets uint = 50

/* this is a map syntax
var bookings = make([]map[string]string, 0) ( this variable is a slice of string *it is dynamic in size better than array)
*/

var bookings = make([]UserData, 0) // THIS A STRUCT SYNTAX

// struct is created to store the data of different types, it is similar as maps but save the data of different types whereas maps save of same types.

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

/* the wait groupp created outside the main fuction so it will wait for the concurent thread we created before exisiting the function */
var wg = sync.WaitGroup{} // var wg to save the function in the variable to use later

func main() {

	greetUsers()

	//for {  we are not using it since we don't want the app to keep going for one user

	// user input

	firstName, lastName, email, userTicket := getUserInput()

	// validating function goes here
	isValidName, isValidEmail, isValidTicketsNumber := helper.GetUserValidation(firstName, lastName, email, userTicket, remainingTickets)

	if isValidName && isValidEmail && isValidTicketsNumber {

		// calling bookTicket function

		bookTicket(userTicket, firstName, lastName, email)

		/* here we are making our app concurrent by wrinting "go" infront of my function
		           Concurrency means that i have this function which is taking time to execute so instead of waiting
				   we have created a separate thread so it won't interept the main thread.
		*/

		//wg.Add() is the function of wait group added before we add the concurrent thread
		wg.Add(1) // this is to add this thread in wait sync , we can add no. of threads we want to wait for.
		go sendTicket(userTicket, firstName, lastName, email)

		// call function print frist names, we wrote return in function below so it‹ç return in main func and prints the value here
		firstNames := getFirstName()
		fmt.Printf("These are all of our bookings: %v \n", firstNames)

		if remainingTickets == 0 {
			// end the program
			fmt.Println(" Tickets sold out ☹️, Try Luck next Year. ")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println(" your name or last name is too short")
		}
		if !isValidEmail {
			fmt.Println(" your email is missing the @")
		}
		if !isValidTicketsNumber {
			fmt.Println(" Number of tickets is invalid ")
		}

	}

	wg.Wait() // add the 2nd function of wait group add the end of the main function  to wait of func "sendTicket()"
	// }
}
func greetUsers() {

	fmt.Printf("Welcome to  our %v booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v available for booking.\n", conferenceTickets, remainingTickets)

	fmt.Printf("Get your tickets to attend %v\n", conferenceName)
}

func getFirstName() []string {

	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	// what is your name
	println("Enter your first Name: ")
	fmt.Scan(&firstName)

	println("Enter your last Name: ")
	fmt.Scan(&lastName)

	println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	fmt.Println("Enter your Email:")
	fmt.Scan(&email)

	return firstName, lastName, email, userTicket

}
func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket

	/* this is a map

	// create map for the user *the maps is the data type to store the value of multiple key value
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)      //the strvconv.formatUint is use to convert the uint to string

	*/
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of booking is %v\n", bookings)

	fmt.Printf("Thank you  %v %v ! for booking %v tickets with us, We have sent the confirmation to your email %v\n. ", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTicket, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket: %v \nto the email  %v\n ", ticket, email)
	fmt.Println("##############")
	wg.Done() // add the 3rd function of the wait group add the end of sendTicket() func so it knows it is done.
}
