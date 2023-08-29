package main

import (
	"gitlab.com/mrossemo/head-first-go/greeting"
	"gitlab.com/mrossemo/head-first-go/greeting/deutsch"
	"gitlab.com/mrossemo/head-first-go/greeting/portuguese"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
	portuguese.Oi()
	portuguese.Ola()
}
