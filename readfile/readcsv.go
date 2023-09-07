package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Row struct {
	Col1, Col2 string
}

func main() {
	file, err := os.Open("list.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	rows := []Row{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		rows = append(rows, Row{
			Col1: row[0], Col2: row[1],
		})
	}

	file, err = os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	for _, row := range rows {
		_ = writer.Write(row)
	}
	writer.Flush()
	file.Close()

}
