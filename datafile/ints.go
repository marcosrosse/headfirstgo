package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetInts(fileName string) ([3]int64, error) {
	var numbers [3]int64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseInt(scanner.Text(), 36, 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
