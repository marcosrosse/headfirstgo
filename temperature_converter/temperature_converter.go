package main

import (
	"fmt"
	"log"

	"github.com/marcosrosse/headfirstgo/keyboard"
)

func main() {
	fmt.Printf("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
