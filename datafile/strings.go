package datafile

import (
	"bufio"
	"os"
)

func GetTexts(fileName string) ([]string, error) {
	var texts []string
	file, err := os.Open(fileName)
	if err != nil {
		return texts, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if err != nil {
			return texts, err
		}
		texts = append(texts, text)
	}
	err = file.Close()
	if err != nil {
		return texts, err
	}
	if scanner.Err() != nil {
		return texts, scanner.Err()
	}
	return texts, nil
}
