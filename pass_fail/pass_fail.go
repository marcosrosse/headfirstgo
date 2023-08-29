package main

import (
	"fmt"
	"log"

	"gitlab.com/mrossemo/head-first-go/keyboard"
)

func main() {
	fmt.Printf("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)

}
