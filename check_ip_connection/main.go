package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Row struct {
	Col1, Col2 string
}

// Check connection with a list the remote host in a specific port.
func main() {
	list := flag.String("l", "", "Specifiy an IP csv list with ip and port.")
	output := flag.String("o", "output.txt", "Specify an output file.")
	timeout := flag.String("t", "", "Specify a timeout to the nc command.")

	flag.Parse()

	file, err := os.Open(*list)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	ips := []Row{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		ips = append(ips, Row{
			Col1: row[0], Col2: row[1],
		})
	}

	file, err = os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, ip := range ips {
		fmt.Printf("Requesting %s on port %s\n", ip.Col1, ip.Col2)

		nc := exec.Command("nc", "-w", *timeout, ip.Col1, ip.Col2)
		_, err := nc.Output()
		if err != nil {
			fmt.Fprintf(file, "Failed to connect to the address: %s:%s\n", ip.Col1, ip.Col2)
		} else {
			fmt.Fprintf(file, "Success to connect to the address: %s:%s\n", ip.Col1, ip.Col2)
		}
	}

}
