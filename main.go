package main

import "fmt"

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
}
