package main

import (
	"fmt"
	"go-booking-app/custompac"
	"strings"
)

// Global var, availale  accross packages (same goes for functions)
var ConferenceName = "Go Conference"

// package level vars, availale in all functions
const conferenceTicketsCount uint = 50

var remainingTickets int64 = 50

// structs/ types

type UserData struct {
	FirstName       string
	SecondName      string
	Email           string
	NumberOfTickets uint
}

func main() {
	// does not make sense to pass this is global vars
	greetUser(ConferenceName, int(conferenceTicketsCount), int(remainingTickets))

	var userName string
	userName = "anonimus"

	var user = UserData{
		FirstName:       "bla1",
		SecondName:      "bla2",
		Email:           "a@a.com",
		NumberOfTickets: 2,
	}

	fmt.Printf("User type %v, type of struct %T", user, user)

	fmt.Println("Please enter your name:")
	fmt.Scan(&userName)
	var userTickets int = 2

	pointerToUserName := &userName
	pointerToUserTickets := &userTickets

	fmt.Printf(
		"User %v, type of string pointer %T, type of int pointer %T, pointer value %v",
		userName,
		pointerToUserName,
		pointerToUserTickets,
		pointerToUserName)

	var bookings = [conferenceTicketsCount]string{"Slava"}
	var bookings2 [50]string

	bookings[1] = "Nana"
	bookings2[0] = "Vova"

	fmt.Printf("The whole array %v\n", bookings)
	fmt.Printf("The whole array 2 %v\n", bookings2)
	fmt.Printf("Type of array 2 %T \n", bookings2)
	fmt.Printf("Lenght of array %v \n", len(bookings))

	// slices (array with dynamic size)
	var bookings3 []string

	for i := 0; i < 5; i++ {
		fmt.Println("Please enter your name:")
		fmt.Scan(&userName)
		bookings3 = append(bookings3, userName)
		remainingTickets--
		fmt.Println("Remaining tickets:", remainingTickets)
		fmt.Printf("The whole slice %v\n", bookings3)
		fmt.Printf("Type of slice %T \n", bookings3)
		fmt.Printf("Lenght of slice %v \n", len(bookings3))
		fmt.Printf("i %v \n", i)
	}

	// foreach
	for index, booking := range bookings3 {
		var names = strings.Fields(booking)
		fmt.Printf("booking # %v: %v\n", index, names[0])
	}

	//while
	for remainingTickets < 10 {
		break
	}

	// infinite loop
	for {
		break
	}

	if remainingTickets == 0 {
		fmt.Printf("All tickets are sold out\n")
	} else if remainingTickets < 25 {
		fmt.Printf("Half of tickets are remaining\n")
	} else {
		fmt.Printf("%v tickets are remaining\n", remainingTickets)
	}

	switch remainingTickets {
	case int64(conferenceTicketsCount):
		fmt.Printf("Nobody buys tickets :(\n")
	case 0:
		fmt.Printf("All tickets are sold out\n")
	case 2, 3:
		fmt.Printf("bla bla\n")
	default:
		fmt.Printf("default value\n")
	}

	resultofFuntionVar := returnSomehtingFunction("bla")
	fmt.Println(resultofFuntionVar)

	bla1, bla2 := returnMultipleVariablesFunc("rand input")
	fmt.Println(bla1, bla2)

	helperFunction1("blaaa")
	custompac.HelperFunctionFromPackage("ttata")
}

// void action
func greetUser(conferenceName string, ticketsCount int, remainingTickets int) {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", ticketsCount, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func returnSomehtingFunction(name string) []string {
	return []string{"res1", "res2"}
}

func returnMultipleVariablesFunc(inputVar string) (string, bool) {
	var stringRes string = "bla"
	boolRes := false

	return stringRes, boolRes
}
