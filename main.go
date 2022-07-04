package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicketsCount uint = 50
	var remainingTickets int64 = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTicketsCount, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var userName string
	userName = "anonimus"

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
}
