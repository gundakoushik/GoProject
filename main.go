package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const tickets = 50

var App = "Ticket Booking Portal"
var Rtickets = tickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	var FName string
	var LName string
	var email string
	var Utickets int
	for Rtickets > 0 {
		fmt.Printf("enter your Fisrt name ")
		fmt.Scan(&FName)
		fmt.Printf("enter your last name ")
		fmt.Scan(&LName)
		fmt.Printf("enter your mai id ")
		fmt.Scan(&email)
		FN, LN, Id := Validation(FName, LName, email)
		if FN && LN && Id {
			fmt.Printf("welcome %v %v to  %v\n", FName, LName, App)
			fmt.Printf("we have %v tickets In %v\n", Rtickets, App)
		} else {
			fmt.Printf("enter valid credentials")
			continue
		}
		fmt.Printf("No of tickets you want to book ")
		fmt.Scan(&Utickets)
		if Rtickets < Utickets {
			fmt.Printf("we only have %v  tickets left please enter valid number of tickets\n", Rtickets)
			continue
		}
		booktickets(FName, LName, email, Utickets)
		wg.Add(1)
		go sendTicket(Utickets, FName, LName, email)
	}
	wg.Wait()
}
func Validation(FName string, LName string, email string) (bool, bool, bool) {
	FN := len(FName) > 2
	LN := len(LName) > 2
	Id := strings.Contains(email, "@")
	return FN, LN, Id
}

func booktickets(FName string, LName string, email string, Utickets int) {
	Rtickets = Rtickets - Utickets
	var userData = UserData{
		firstName:       FName,
		lastName:        LName,
		email:           email,
		numberOfTickets: Utickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FName, LName, Utickets, email)
}
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}


