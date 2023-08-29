package main

import (
	"fmt"

	"github.com/marcosrosse/headfirstgo/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("With the follow-up in", days+dates.DaysInWeek, "days")
}
