package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("Could not open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() { //It'll iterate as long scanner.Scan() returns true.It returns false when the scan stops, either by reaching the end of the input or an error.
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() //After Scan returns false, the Err method will return any error that occurred during scanning, except that if it was io.EOF, Err will return nil.

	if err != nil {
		file.Close()
		return nil, errors.New("Reading the file content failed")
	}

	file.Close()
	return lines, nil
}
