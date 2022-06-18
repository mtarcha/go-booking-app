package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicketsCount = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTicketsCount, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}
