package main

import (
	"github.com/marcosrosse/headfirstgo/greeting"
	"github.com/marcosrosse/headfirstgo/greeting/deutsch"
	"github.com/marcosrosse/headfirstgo/greeting/portuguese"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
	portuguese.Oi()
	portuguese.Ola()
}
