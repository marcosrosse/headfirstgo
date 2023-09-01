package main

import (
	"fmt"
	"log"

	"github.com/marcosrosse/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("/home/mrossemo/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Avarage: %0.2f\n", sum/sampleCount)
}
