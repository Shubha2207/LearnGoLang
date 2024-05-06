package main

import "fmt"

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
