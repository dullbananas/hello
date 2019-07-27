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
func greeting() string {
	return "Hello flat world, the time is: " + time.Now().String()
}

func main() {
	hotelName := "Trivag"
	hotelName += "o"
	fmt.Println(greeting())
	fmt.Println("Hotel: " + hotelName)
}
