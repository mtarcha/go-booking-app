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
	bookings3 = append(bookings3, "Tanja", "Vanja")
	fmt.Printf("The whole slice %v\n", bookings3)
	fmt.Printf("Type of slice %T \n", bookings3)
	fmt.Printf("Lenght of slice %v \n", len(bookings3))
}
