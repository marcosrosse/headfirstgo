package main

import (
	"fmt"
	"log"
)

func main() {
	var amount, total float64
	amount, err := paintNeeded(-4.2, 3.0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
}

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid! The value must be higer than 0.", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid! The value must be higer than 0.", height)
	}
	area := width * height
	return area / 10.0, nil
}
