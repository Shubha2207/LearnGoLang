// all code must belong to a package
package main

// import format package
import (
	"fmt"
)

// package level variables
const CONFERENCE_NAME = "GO Conference"

/**
* go vs java performance using factorial program
* https://medium.com/@radadiyasunny970/a-simple-performance-test-and-difference-go-v-s-java-e6f29ad65293
* try using multi-threading
*  https://medium.com/deno-the-complete-reference/go-gin-vs-springboot-hello-world-perfomance-comparison-e535c9d6c36
 */

// we need to provide an entry point for executing our code
func main() {
	// variables and constants
	const confTickets = 50
	var remainingTickets = confTickets

	fmt.Println("Welcome to", CONFERENCE_NAME, "booking application")
	fmt.Printf("Book your tickets for %v! (Remaining Tickets %d)\n", CONFERENCE_NAME, remainingTickets)

	// datatypes
	var userName string
	var userTickets int

	// userName = "John"
	// userTickets = 10
	// fmt.Println(userName, userTickets)
	// fmt.Printf("confTicket is of type %T", confTickets)

	// Getting user input
	// fmt.Println("Enter first name: ")
	// fmt.Scan(&userName)
	// fmt.Println("Enter the number of tickets: ")
	// fmt.Scan(&userTickets)
	// fmt.Printf("User onboarded: %v , Number of Tickets: %d\n", userName, userTickets)
	// remainingTickets = remainingTickets - userTickets
	// fmt.Println("Remaining Tickets:", remainingTickets)

	// array creation
	// var arr1 [10]string // only defined
	// arr1[0] = "abc"

	// var arr2 = [10]string{"abc", "pqr"} // arr initialization
	// fmt.Println("array : ", arr2)

	// slice creation
	var names []string
	// names = append(names, "abc")
	// fmt.Println(names)
	// names = append(names, "pqr", "xyz")
	// fmt.Println("Slice names: ", names)

	// // print data type of array and slice
	// fmt.Printf("data type %T\n", arr2)
	// fmt.Printf("data type %T\n", names)

	// // get size of a slice
	// fmt.Println(len(names))

	// loops in go
	// infinite loop
	for {
		fmt.Println("Enter first name: ")
		fmt.Scan(&userName)

		if !isValidInput(userName) {
			fmt.Println("Invalid name of the user. No of characters must be more than 1")
			continue
		}

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)
		// precheck
		if userTickets < 0 && userTickets > remainingTickets {
			fmt.Println("Cannot fulfill the request. Remaining Tickets :", remainingTickets)
			if wantToStop() {
				break
			} else {
				continue
			}

		}
		fmt.Printf("User onboarded: %v , Number of Tickets: %d\n", userName, userTickets)
		remainingTickets = remainingTickets - userTickets
		fmt.Println("Remaining Tickets:", remainingTickets)
		names = append(names, userName)
		fmt.Println("Users: ", names)
		if remainingTickets == 0 {
			fmt.Println("All tickets sold out")
			break
		}
		// check if user want to stop
		if wantToStop() {
			break
		}

	}

	// for-each loop
	nameInitials := []string{}
	for _, nameInitial := range names {
		nameInitials = append(nameInitials, nameInitial[0:1])
		// fmt.Printf("namesInitial type %T\n", nameInitial)
		// fmt.Printf("index type %T\n", index) // underscore is nothing but an index
		// we are using underscore to tell go that we are explicitly ignoring this variable
	}

	fmt.Println("names initials : ", nameInitials)

	// switch-case in golang
	// city := "abc"
	// switch city {
	// case "pqr":
	// 	fmt.Println("pqr city")
	// 	break
	// case "abc":
	// 	fmt.Println("abc city")
	// 	break
	// default:
	// 	fmt.Println("invalid city")
	// }

}

/**
* check if user wants to stop the application
* return true - user wants to stop the application
*        false - otherwise
 */
func wantToStop() bool {
	var stop string
	fmt.Println("Do you want to stop the application ?")
	fmt.Scan(&stop)
	return stop == "yes"
}

/**
* check if the input is valid
 */
func isValidInput(name string) bool {
	return len(name) > 1
}
