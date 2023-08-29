// SIMPLE FUNCTION THAT USE POINTERS
package main

import "fmt"

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}
