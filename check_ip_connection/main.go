package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/marcosrosse/headfirstgo/datafile"
)

// To automate the test of connectivity in a list of IP address + port using curl telnet
// The list must be:
// ipAddress:port
// ipAddress:port

func main() {
	list := flag.String("l", "", "Specifiy a list with IP address + port number")
	output := flag.String("o", "output.txt", "Specify an output file")

	flag.Parse()

	ipAddress, err := datafile.GetTexts(*list)
	if err != nil {
		log.Fatal("You need to specify a list file with ip + port. ", err)
	}

	file, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, ip := range ipAddress {
		nc := exec.Command("curl", "-vvk", "--connect-timeout", "2", "telnet://"+ip)
		out, err := nc.Output()
		if err != nil {
			fmt.Fprintf(file, "Failed to connect to the address: %s\n", ip)
		} else {
			fmt.Fprintf(file, "Success to connect to the address: %s. Protocol: %s\n", ip, string(out))
		}

	}

}
