/*
hello is a terrible program that does the following things:
  * says that the earth is flat
  * shows the current time
  * advertises Trivago
*/
package main

import (
	"fmt"
	"time"
)

// greeting returns a string with some info.
func greeting(adjective string) string {
	return "Hello " + adjective + " world, the time is: " + time.Now().String()
}

func main() {
	hotelName := "Trivag"
	hotelName += "o"
	fmt.Println(greeting("flat"))
	fmt.Println("Hotel: " + hotelName)
}
