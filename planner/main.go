package main

import (
	"fmt"

	"gitlab.com/mrossemo/head-first-go/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("With the follow-up in", days+dates.DaysInWeek, "days")
}
